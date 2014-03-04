/*
    main.go -
 */
package main

import (
    "fmt"
    "flag"
    "time"
//    "net/http"
    "bitbucket.org/miranr/artistic/core"
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

type ArtisticUser struct {
    // 
    Username string

    // Is the user 
    Authenticated bool
}

type ArtisticCtrl struct {

    // a list of current users
    Users []ArtisticUser

    // a log filename
    logFname string

    // a syslog IP address
    syslogIP string

    // a logger
    log *utils.Log

    // config file path
    configFile string

    // MongoDB session 
    db *mgo.Session

    // a debug flag (only for testing purposes)
    debug bool
}

/*
 * parseArgs - parse the CLI arguments
 */
func parseArgs(ac *ArtisticCtrl) {

    if ac == nil {
        panic("FATAL: The main control structure is NOT defined...")
    }
    flag.StringVar(&ac.logFname, "l", defineDefLogFname(),
            "define the custom log file path (absolute, please!)")
    flag.StringVar(&ac.syslogIP, "s", "", "IP address of the Syslog server")
    flag.StringVar(&ac.configFile, "c", DefConfigFile,
            "define custom path for config file")
    flag.BoolVar(&ac.debug, "d", false, "enable debug mode (only for testing!)")
    flag.Parse()
}

func main () {
    //
    ac := new(ArtisticCtrl)

    // handle config file
    handleConfigFile(ac)

    // parse the CLI arguments
    parseArgs(ac)

    // create the logger
    createLog(ac)

    var err error

    // connect to MongoDB (NOTE: currently hardcoded, should be read from 
    // config file in the final version)
    url := db.CreateUrl("localhost", 27017, "artistic", "artistic", "artistic")
    if ac.db, err = db.Connect(url, DatabaseTimeout); err != nil {
        panic("Connection to MongoDB cannot be established.")
    }
    ac.log.Notice("Connection to MongoDB established.")
    defer db.Close(ac.db)

    //testing import for local code
    p := core.CreatePainter()
    p.Name = core.CreateName("Vincent", "", "Van Gogh")
    fmt.Println(p.String())

    fmt.Println("Serving application on 'localhost:8080'...")

    webStart(ac, DefWebRoot)
}
