package main

/*
   log.go -
*/
import (
	"bitbucket.org/miranr/artistic/utils"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// Let's define the default log levels for different log handlers
const (
	defSyslogLevel utils.Severity = utils.Notice
	defFileLevel   utils.Severity = utils.Informational
)

/*
 * createLog -
 */
func createLog(ac *ArtisticApp) (err error) {

	if ac == nil {
		panic("FATAL: The main control structure is NOT defined...")
	}

	ac.Log = utils.NewLog()

	// define the name of the log file
	if ac.LogFname == "" {
		ac.LogFname = defineDefLogFname(ac.WorkDir)
	} else {
		ac.LogFname = filepath.Join(ac.WorkDir, "artistic.log")
	}

	format := "%s %s %s"
	err = createLoggers(ac, format, ac.Debug)
	if err != nil {
		return err
	}

	// start the logger and display the first message...
	ac.Log.Start()
	ac.Log.Info("Log successfully created.")

	return nil
}

func createLoggers(ac *ArtisticApp, format string, debug bool) error {

	if ac == nil {
		panic("FATAL: The main control structure is NOT defined...")
	}
	var err error = nil

	// define default log levels
	fLevel := defFileLevel
	sLevel := defSyslogLevel
	if ac.Debug {
		fLevel = utils.Debug
		sLevel = utils.Debug
	}

	// add file log handler

	f, err := utils.NewFileHandler(ac.LogFname, fmt.Sprintf("%s\n", format), fLevel)
	if f != nil {
		ac.Log.Handlers = ac.Log.AddHandler(f)
	}

	// add syslog log handler
	if ac.SyslogIP != "" {
		s := utils.NewSyslogHandler(ac.SyslogIP, format, sLevel)
		if s != nil {
			ac.Log.Handlers = ac.Log.AddHandler(s)
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
func defineDefLogFname(workdir string) string {

	defDir := "/var/log/artistic.log"
	if runtime.GOOS == "windows" {
		//defDir = path.Join(aa.WorkDir, "artistic.log")
		defDir = path.Join(workdir, "artistic.log")
	}
	return filepath.Clean(defDir)
}

// Reads a log file and returns it as a list of lines
func readLog(filename string) ([]string, error) {

	// we read a log file
	var err error
	var data []byte
	if data, err = ioutil.ReadFile(filename); err != nil {
		return []string{""}, err
	}

	// now we convert the text into an array of lines
	lines := strings.Split(string(data), "\n")

	// the last line is always an empty one, slice it out...
	return lines[:len(lines)-1], nil
}
