/*
    log.go -
 */
package main

import (
    "fmt"
//    "os"
    "path"
    "path/filepath"
    "runtime"
    "bitbucket.org/miranr/artistic/utils"
)

// Let's define the default log levels for different log handlers:
//   
const (
    numOfLogHandlers int = 2
    defSyslogLevel utils.LogLevel = utils.NoticeLogLevel
    defFileLevel   utils.LogLevel = utils.InfoLogLevel
)

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
    f, err := utils.NewFileHandler(ac.logFname,
                                   fmt.Sprintf("%s\n", format), fLevel)
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
//        defDir = path.Join(os.Getenv("USERPROFILE"), "artistic.log")
        defDir = path.Join(ac.workDir, "artistic.log")
    }
    return filepath.Clean(defDir)
}

