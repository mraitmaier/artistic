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
    "bitbucket.org/miranr/artistic/utils"
    "bitbucket.org/miranr/artistic/db"
    "labix.org/v2/mgo"
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
/*
type ArtisticUser struct {
    // 
    Username string

    // Is the user 
    Authenticated bool
}
*/

type ArtisticCtrl struct {

    // a list of current users
//    Users []ArtisticUser

    // working folder
    workDir string

    // a log filename
    logFname string

    // a syslog IP address
    syslogIP string

    // a logger
    log *utils.Log

    // config file path
    configFile string

    // MongoDB session 
    dbsess *mgo.Session

    // folder for session files
    sessDir string

    // a debug flag (only for testing purposes)
    debug bool
}

// create new global struct instance
var ac = new(ArtisticCtrl)

/*
 * parseArgs - parse the CLI arguments
 */
func parseArgs(ac *ArtisticCtrl) {

    if ac == nil {
        fmt.Println("FATAL: The main control structure is NOT defined...")
        os.Exit(1)
    }
    flag.StringVar(&ac.logFname, "l", defineDefLogFname(),
            "define the custom log file path (absolute, please!)")
    flag.StringVar(&ac.syslogIP, "s", "", "IP address of the Syslog server")
    flag.StringVar(&ac.configFile, "c", DefConfigFile,
            "define custom path for config file")
    flag.BoolVar(&ac.debug, "d", false, "enable debug mode (only for testing!)")
    flag.Parse()
}

// Define working folder and create it, if it doesn't exist
func setWorkDir() bool {

    // if working folder is already set in global struct, use it
    wdir := ac.workDir

    // otherwise, the default working folder is the "artistic" folder in $HOME
    if wdir == "" {
        switch runtime.GOOS {
            case "windows": wdir = os.Getenv("USERPROFILE")
            default: wdir = os.Getenv("HOME")
        }
        ac.workDir = filepath.Join(wdir, "artistic")
    }

    // create the working folder, if it doesn't exist
    if err := os.MkdirAll(ac.workDir, 0755); err != nil {
        return false
    }

    return true
}

// Cleanup oprocedure when app is terminated.
func cleanup() {

    // close the DB connection
    if ac.dbsess != nil {
        db.Close(ac.dbsess)
        ac.log.Notice("Connection to MongoDb closed.\n")
    }

    // clean the sessions directory
    cleanSessDir()
    ac.log.Info("Sessions folder deleted.\n")

    // close the log 
    ac.log.Info("Closing log.\n")
    ac.log.Close()
}

func main () {
    // parse the CLI arguments
    parseArgs(ac)

    // handle config file
    handleConfigFile(ac)

    // set working directory
    if !setWorkDir() {
        fmt.Println("FATAL: cannot create working folder, cannot continue...")
        return
    }

    // create the logger
    createLog(ac)

    // deferring the the cleanup procedure when app is terminated normally
    defer cleanup()

    var err error
    // connect to MongoDB (FIXME: currently hardcoded, should be read from 
    // config file in the final version)
    url := db.CreateUrl("localhost", 27017, "artistic", "artistic", "artistic")
    if ac.dbsess, err = db.Connect(url, DatabaseTimeout); err != nil {
        ac.log.Critical("Connection to MongoDB cannot be established.\n")
        fmt.Println("Connection to MongoDB cannot be established.")
        fmt.Println("Exiting...")
        return
    }
    ac.log.Notice("Connection to MongoDB established.")

    // handle CTRL-C signal and perform cleanup before app is terminated
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        <-c
        ac.log.Info("Received a CTRL-C signal to terminate.\n")
        cleanup()
        os.Exit(0) // CTRL-C is clean exit for this app...
    }()

    // start web interface
    fmt.Println("Serving application on 'localhost:8088'...")
    ac.log.Info("Serving application on 'localhost:8088'...\n")
    webStart(DefWebRoot)
}
