/*
    main.go -
 */
package main

import (
    "fmt"
    "flag"
    "time"
    "os"
//    "runtime"
    "os/signal"
//    "path/filepath"
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

func main () {

    configfile := ""
    // parse the CLI arguments
    parseArgs(aa, &configfile)

    // handle config file
    if err := aa.HandleConfigFile(configfile); err != nil {
        fmt.Println("FATAL: cannot read config file. Cannot continue...")
        os.Exit(1)
    }

    // set working directory
    if !aa.SetWorkDir() {
        fmt.Println("FATAL: cannot create working folder, cannot continue...")
        return
    }

    // create the logger
    aa.createLogs()

    // deferring the the cleanup procedure when app is terminated normally
    defer aa.Cleanup()

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
    aa.Log.Info("Connection to MongoDB established.")

    // catch CTRL-C signal and perform cleanup before app is terminated
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        <-c
        aa.Log.Info("Received a CTRL-C signal to terminate.")
        aa.Cleanup()
        os.Exit(0) // CTRL-C is clean exit for this app...
    }()

    // start web interface
    fmt.Println("Serving application on 'localhost:8088'...")
    aa.Log.Info("Serving application on 'localhost:8088'...")
    if err = aa.startWeb(DefWebRoot); err != nil {
        aa.Log.Error(err.Error())
    }
}
