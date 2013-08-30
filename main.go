/*
    main.go -
 */
package main

import (
    "fmt"
    "flag"
    "net/http"
    "bitbucket.org/miranr/artistic/core"
    "bitbucket.org/miranr/artistic/utils"
    "bitbucket.org/miranr/artistic/db"
    "labix.org/v2/mgo"
)

const (

    // default path to config file
    DefConfigFile string = "./artistic.cfg"
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

func testHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Artistic Test Web Page")
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
    flag.StringVar(&ac.syslogIP, "s", "127.0.0.1", 
            "IP address of the Syslog server")
    flag.BoolVar(&ac.debug, "d", false, "enable debug mode (only for testing!)")
    flag.StringVar(&ac.configFile, "c", DefConfigFile, 
            "Define custom path for config file")

    flag.Parse()
}

func main () {
    //
    ac := new(ArtisticCtrl)

    // handle config file
    //handleConfigFile(ac)

    // parse the CLI arguments
    parseArgs(ac)

    // create the logger
    createLog(ac)

    // connect to MongoDB
    uri := db.CreateUri("artistic", "artistic", "127.0.0.1")
    s, err := db.Connect(uri)
    if err != nil {
        panic("Connection to MongoDB cannot be established.")
    }
    ac.db = s
    ac.log.Notice("Connection to MongoDB established.")

    //testing import for local code
    p := core.CreatePainter()
    p.Name = core.CreateName("Vincent", "", "Van Gogh")
    fmt.Println(p.String())

    fmt.Println("Serving application on 'localhost:8080'...")

    http.HandleFunc("/", testHandler)
    http.ListenAndServe(":8080", nil)
}
