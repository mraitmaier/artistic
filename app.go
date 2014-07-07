/*
   app.go -
*/
package main

import (
	"fmt"
	"os"
	"runtime"
	//    "time"
	//    "errors"
	"path/filepath"
	//    "bitbucket.org/miranr/artistic/core"
	"bitbucket.org/miranr/artistic/db"
	"bitbucket.org/miranr/artistic/utils"
)

//var cleanupTime = time.Second * 1

type ArtisticApp struct {

	// working folder
	WorkDir string

	// a log filename
	LogFname string

	// a syslog IP address
	SyslogIP string

	// a logger
	Log *utils.Log

	// DB session
    DbSess db.DbConnector

    // DB data provider
    DataProv db.DataProvider

	// folder for session files
	SessDir string

	// a web stuff structure instance
	*WebInfo

	// a debug flag (only for testing purposes)
	Debug bool
}

func (a *ArtisticApp) createLogs() { createLog(a) }

func (a *ArtisticApp) startWeb(path string) error { return webStart(a, path) }

func (a *ArtisticApp) HandleConfigFile(cfgfile string) error {

	fmt.Printf("DEBUG config file: %q\n", cfgfile) // DEBUG
    // TODO

	return nil
}

// Define working folder and create it, if it doesn't exist
func (a *ArtisticApp) SetWorkDir() bool {

	// if working folder is already set in global struct, use it
	wdir := a.WorkDir

	// otherwise, the default working folder is the "artistic" folder in $HOME
	if wdir == "" {
		switch runtime.GOOS {
		case "windows":
			wdir = os.Getenv("USERPROFILE")
		default:
			wdir = os.Getenv("HOME")
		}
		a.WorkDir = filepath.Join(wdir, "artistic")
	}

	// create the working folder, if it doesn't exist
	if err := os.MkdirAll(a.WorkDir, 0755); err != nil {
		return false
	}

	return true
}

// Cleanup procedure when app is terminated.
func (a *ArtisticApp) Cleanup() {

	// close the DB connection
	if a.DbSess != nil {
		a.DbSess.Close()
		a.Log.Notice("Connection to MongoDb closed.")
	}

	// clean the sessions directory
    if a.WebInfo != nil {
	    cleanSessDir(a)
	    a.Log.Info("Sessions folder deleted.")
    }

	// close the log
	a.Log.Info("Closing log.")
	a.Log.Close()
}

/*
func (a *ArtisticApp) Init() error {

    configfile := ""

    // handle config file
    handleConfigFile(a, configfile)

    // set working directory
    if !a.SetWorkDir() {
        return errors.New("Cannot create working folder, cannot continue...")
    }

    // create the logger
    //createLog(a)

    // deferring the the cleanup procedure when app is terminated normally
    //defer a.cleanup()

    var err error
    // connect to MongoDB (FIXME: currently hardcoded, should be read from
    // config file in the final version)
    url := db.CreateUrl("localhost", 27017, "artistic", "artistic", "artistic")
    if a.DbSess, err = db.Connect(url, DatabaseTimeout); err != nil {
        a.Log.Critical("Connection to MongoDB cannot be established.")
        fmt.Println("Connection to MongoDB cannot be established.")
        fmt.Println("Exiting...")
        return err
    }
    a.Log.Notice("Connection to MongoDB established.")

    // handle CTRL-C signal and perform cleanup before app is terminated
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, os.Kill)
    go func() {
        <-c
        ac.log.Info("Received a CTRL-C signal to terminate.")
        cleanup()
        os.Exit(0) // CTRL-C is clean exit for this app...
    }()

    // start web interface
    fmt.Println("Serving application on 'localhost:8088'...")
    a.Log.Info("Serving application on 'localhost:8088'...")
    webStart(DefWebRoot)
    return nil
}
*/
