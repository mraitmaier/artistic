/*
    main.go -
 */
package main

import (
    "fmt"
    "flag"
    "time"
    "os"
    "runtime"
    "os/signal"
    "path/filepath"
//    "net/http"
//    "bitbucket.org/miranr/artistic/core"
//    "bitbucket.org/miranr/artistic/utils"
    "bitbucket.org/miranr/artistic/db"
//    "labix.org/v2/mgo"
)

const (

    // default path to config file
    DefConfigFile string = "./artistic.cfg"

    // default timeout for DB connect 
    DatabaseTimeout time.Duration = 5 * time.Second

    // default web server root
    //DefWebRoot = "D:/Miran/koda/hg/go/src/bitbucket.org/miranr/artistic/web"
    DefWebRoot = "./web/"
)

// create a new Artistic application instance
var aa = new(ArtisticApp)

/*
 * parseArgs - parse the CLI arguments
 */
func parseArgs(ac *ArtisticApp, cfgfile *string) {

    if ac == nil {
        fmt.Println("FATAL: The main control structure is NOT defined...")
        os.Exit(1)
    }
    //flag.StringVar(&ac.logFname, "l", defineDefLogFname(),
    flag.StringVar(&ac.LogFname, "l", "",
            "define the custom log file path (absolute, please!)")
    flag.StringVar(&ac.SyslogIP, "s", "", "IP address of the Syslog server")
    flag.StringVar(cfgfile, "c", DefConfigFile,
            "define custom path for config file")
    flag.BoolVar(&ac.Debug, "d", false, "enable debug mode (only for testing!)")
    flag.Parse()
}

// Define working folder and create it, if it doesn't exist
func setWorkDir() bool {

    // if working folder is already set in global struct, use it
    wdir := aa.WorkDir

    // otherwise, the default working folder is the "artistic" folder in $HOME
    if wdir == "" {
        switch runtime.GOOS {
            case "windows": wdir = os.Getenv("USERPROFILE")
            default: wdir = os.Getenv("HOME")
        }
        aa.WorkDir = filepath.Join(wdir, "artistic")
    }

    // create the working folder, if it doesn't exist
    if err := os.MkdirAll(aa.WorkDir, 0755); err != nil {
        return false
    }

    return true
}

// Cleanup oprocedure when app is terminated.
func cleanup() {

    // close the DB connection
    if aa.DbSess != nil {
        db.Close(aa.DbSess)
        aa.Log.Notice("Connection to MongoDb closed.")
    }

    // clean the sessions directory
    cleanSessDir()
    aa.Log.Info("Sessions folder deleted.")

    // close the log 
    aa.Log.Info("Closing log.")
    aa.Log.Close()
}

func main () {
    configfile := ""
    // parse the CLI arguments
    parseArgs(aa, &configfile)

    // handle config file
    handleConfigFile(aa, configfile)

    // set working directory
    if !aa.SetWorkDir() {
        fmt.Println("FATAL: cannot create working folder, cannot continue...")
        return
    }

    // create the logger
    createLog(aa)

    // deferring the the cleanup procedure when app is terminated normally
    defer cleanup()

    var err error
    // connect to MongoDB (FIXME: currently hardcoded, should be read from 
    // config file in the final version)
    url := db.CreateUrl("localhost", 27017, "artistic", "artistic", "artistic")
    if aa.DbSess, err = db.Connect(url, DatabaseTimeout); err != nil {
        aa.Log.Critical("Connection to MongoDB cannot be established.")
        fmt.Println("Connection to MongoDB cannot be established.")
        fmt.Println("Exiting...")
        return
    }
    aa.Log.Notice("Connection to MongoDB established.")

    // handle CTRL-C signal and perform cleanup before app is terminated
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        <-c
        aa.Log.Info("Received a CTRL-C signal to terminate.")
        cleanup()
        os.Exit(0) // CTRL-C is clean exit for this app...
    }()

    // start web interface
    fmt.Println("Serving application on 'localhost:8088'...")
    aa.Log.Info("Serving application on 'localhost:8088'...")
    webStart(DefWebRoot)
}
