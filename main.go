
package main

import (
    "fmt"
    "flag"
    "net/http"
    "bitbucket.org/miranr/artistic/artistic"
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
    log *artistic.Log

    // a debug flag (only for testing purposes)
    debug bool

}

// Let's define the default log levels for different log handlers:
//   
const (
    numOfLogHandlers int = 2
    defSyslogLevel artistic.LogLevel = artistic.NoticeLogLevel
    defFileLevel   artistic.LogLevel = artistic.InfoLogLevel
)

func testHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Artistic Test Web Page")
}


/*
 * createLog - 
 */
func createLog(ac *ArtisticCtrl) (err error) {

    if ac == nil {
        panic("FATAL: The main control structure is NOT defined...")
    }

    ac.log = artistic.NewLog(numOfLogHandlers)

    if ac.logFname != "" {
        format := "%s %s %s"
        err := createLoggers(ac, format, ac.debug)
        if err != nil {
            return err
        }
        ac.log.Info("Log successfully created\n")
    }
    return nil
}

func createLoggers(ac * ArtisticCtrl, format string, debug bool) error {

    if ac == nil {
        panic("FATAL: The main control structure is NOT defined...")
    }
   var err error = nil
    // define default log levels
    fLevel := defFileLevel
    sLevel := defSyslogLevel
    if ac.debug {
        fLevel = artistic.DebugLogLevel
        sLevel = artistic.DebugLogLevel
    }
    // add file log handler
    f, err := artistic.NewFileHandler(ac.logFname, format, fLevel)
    if f != nil {
        ac.log.Handlers = ac.log.AddHandler(f)
    }
    // add syslog log handler
    if ac.syslogIP != "" {
        s := artistic.NewSyslogHandler(ac.syslogIP, format, sLevel)
        if s != nil {
            ac.log.Handlers = ac.log.AddHandler(s)
        }
    }
    return err
}

/*
 *
 */
func defineDefLogFname() string {
    return "path to log file (default is /var/log/artistic.log)" // FIXME
}

/*
 * parseArgs - parse the CLI arguments
 */
func parseArgs(ac *ArtisticCtrl) {

    if ac == nil {
        panic("FATAL: The main control structure is NOT defined...")
    }

    flag.StringVar(&ac.logFname, "l", "/var/log/artistic.log",
                                 defineDefLogFname())
    flag.StringVar(&ac.syslogIP, "s", "", "IP address of the Syslog server")
    flag.BoolVar(&ac.debug, "d", false, "enable debug mode (only for testing!)")

    flag.Parse()
}

func main () {
    //
    ac := new(ArtisticCtrl)

    // parse the CLI arguments
    parseArgs(ac)

    // create the logger
    createLog(ac)

    //testing import for local code
    p := artistic.CreatePainter()
    p.Name = artistic.CreateName("Vincent", "", "Van Gogh")
    fmt.Println(p.String())

    fmt.Println("Serving application on 'localhost:8080'...")

    http.HandleFunc("/", testHandler)
    http.ListenAndServe(":8080", nil)
}
