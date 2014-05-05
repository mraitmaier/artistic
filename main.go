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
    dbsess *mgo.Session

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

    // parse the CLI arguments
    parseArgs(ac)

    // handle config file
    handleConfigFile(ac)

    // create the logger
    createLog(ac)

    var err error

    // connect to MongoDB (NOTE: currently hardcoded, should be read from 
    // config file in the final version)
    url := db.CreateUrl("localhost", 27017, "artistic", "artistic", "artistic")
    if ac.dbsess, err = db.Connect(url, DatabaseTimeout); err != nil {
        //panic("Connection to MongoDB cannot be established.")
        fmt.Println("Connection to MongoDB cannot be established.")
        fmt.Println("Exiting...")
        return
    }
    ac.log.Notice("Connection to MongoDB established.")
    // deferring the DB connection close; must use a closure
    defer func() {
                    db.Close(ac.dbsess)
                    ac.log.Notice("Connection to MongoDb closed.")
          }()

    //testing import for local code
    p := core.CreatePainter()
    p.Name = core.CreateName("Vincent", "", "Van Gogh")
    fmt.Println(p.String())

    fmt.Println("Serving application on 'localhost:8088'...")

    webStart(DefWebRoot)
}
