
package main

import (
    "fmt"
    "flag"
    "os"
    "path"
    "path/filepath"
    "runtime"
    "net/http"
    "bitbucket.org/miranr/artistic/core"
    "bitbucket.org/miranr/artistic/utils"
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

    // a debug flag (only for testing purposes)
    debug bool
}

// Let's define the default log levels for different log handlers:
//   
const (
    numOfLogHandlers int = 2
    defSyslogLevel utils.LogLevel = utils.NoticeLogLevel
    defFileLevel   utils.LogLevel = utils.InfoLogLevel
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

    ac.log = utils.NewLog(numOfLogHandlers)

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
        fLevel = utils.DebugLogLevel
        sLevel = utils.DebugLogLevel
    }
    // add file log handler
    f, err := utils.NewFileHandler(ac.logFname, format, fLevel)
    if f != nil {
        ac.log.Handlers = ac.log.AddHandler(f)
    }
    // add syslog log handler
    if ac.syslogIP != "" {
        s := utils.NewSyslogHandler(ac.syslogIP, format, sLevel)
        if s != nil {
            ac.log.Handlers = ac.log.AddHandler(s)
        }
    }
    return err
}

/*
 * defineDefLogName - define a default log file location
 *
 * This is private function that defines the default path for log file.
 * If app is run on Unix/Linux environment, the default path is standard
 * '/var/log/artistc.log'. In the case of WinXY environment, the default is
 * taken from '%USERPROFILE%' env variable (this is usually
 * 'c:\Users\<Username>'). 
 */
func defineDefLogFname() string {

    defDir := "/var/log/artistic.log"

    if runtime.GOOS == "windows" {
        defDir = path.Join(os.Getenv("USERPROFILE"), "artistic.log")
    }

    return filepath.Clean(defDir)
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
    p := core.CreatePainter()
    p.Name = core.CreateName("Vincent", "", "Van Gogh")
    fmt.Println(p.String())

    fmt.Println("Serving application on 'localhost:8080'...")

    http.HandleFunc("/", testHandler)
    http.ListenAndServe(":8080", nil)
}
