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

func (a *ArtisticApp) HandleConfigFile(cfgfile string) error {

	fmt.Printf("DEBUG config file: %q\n", cfgfile) // DEBUG
    // TODO

	return nil
}

// Define working folder and create it, if it doesn't exist
func (a *ArtisticApp) SetWorkDir() error {

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
		return err
	}

	return nil
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

    // close the websockets connection
    //if a.WebInfo.wsConn != nil { a.WebInfo.wsConn.Close() }

	// close the log
	a.Log.Info("Closing log.")
	a.Log.Close()
}


