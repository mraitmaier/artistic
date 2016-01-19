package main
//
// main.go -
//

import (
	"flag"
	"fmt"
	"os"
	"time"
	"os/signal"
	"github.com/mraitmaier/artistic/db"
	"github.com/mraitmaier/artistic/core"
)

const (

	// DefConfigFile designates a default path to config file
	DefConfigFile string = "./artistic.cfg"

	// DatabaseTimeout defines a default timeout for DB connect
	DatabaseTimeout time.Duration = 5 * time.Second

	// DefWebRoot defines a default web server root
	DefWebRoot = "./web/"

	// DBName defines a Database name
	DBName = "artistic"

	// DBUser defines the default Database user name
	DBUser = "artistic"

	// DBPwd is the default password for the DB access
	DBPwd = "My9xpk$!W"
)

// create a new Artistic application instance
//var aa = new(ArtisticApp)

// Parse the CLI arguments.
func parseArgs(ac *ArtisticApp, cfgfile *string) {

	if ac == nil {
		fmt.Println("FATAL: The main control structure is NOT defined...")
		os.Exit(1)
	}
	flag.StringVar(&ac.LogFname, "l", "", "define the custom log file path (absolute, please!)")
	flag.StringVar(&ac.SyslogIP, "s", "", "IP address of the Syslog server")
	flag.StringVar(cfgfile, "c", DefConfigFile, "define custom path for config file")
	flag.BoolVar(&ac.Debug, "d", false, "enable debug mode (only for testing!)")
	flag.Parse()
}

// MAIN: this is where all begins...
func main() {

	// create a new Artistic application instance
	var aa = new(ArtisticApp)

	// parse the CLI arguments
	configfile := ""
	parseArgs(aa, &configfile)

	// handle config file
	if err := aa.HandleConfigFile(configfile); err != nil {
		fmt.Println("FATAL: cannot read config file. Cannot continue...")
		os.Exit(1)
	}

	// set working directory
	if err := aa.SetWorkDir(); err != nil {
		fmt.Println("FATAL: cannot create working folder, cannot continue...")
		fmt.Printf("%s\n", err.Error())
		return
	}

	// create the logger
	aa.createLogs()

	// deferring the the cleanup procedure when app is terminated normally
	defer aa.Cleanup()

	var err error
	var url string
	// connect to database (FIXME: currently hardcoded, should be read from
	// config file in the final version)
	url, aa.DbSess, aa.DataProv, err = db.InitDb(db.MongoDB, "localhost", 27017, DBUser, DBPwd, DBName)
	if err = aa.DbSess.Connect(url, DatabaseTimeout); err != nil {
		aa.Log.Critical("Connection to database cannot be established.")
		fmt.Println("Connection to database cannot be established. Exiting...")
		return
	}
	aa.Log.Info("Connection to MongoDB established.")

    // now initialize DB at startup
    if err = initDB(aa); err != nil {
        msg := fmt.Sprintf("Fatal error initializing database: '%s'. Exiting", err.Error())
		aa.Log.Critical(msg)
		fmt.Printf(msg)
		return
    }

	// catch CTRL-C signal and perform cleanup before app is terminated
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		aa.Log.Info("Received a CTRL-C signal to terminate.")
		aa.Cleanup()
		os.Exit(0) // CTRL-C is clean exit for this app...
	}()

	// start web interface
	fmt.Println("Serving application on 'localhost:8088'...")
	aa.Log.Info("Serving application on 'localhost:8088'...")
	if err = webStart(aa, DefWebRoot); err != nil {
		aa.Log.Error(err.Error())
	}
}

// The initDB function checks some stuff at startup and creates some records if needed. 
func initDB(aa *ArtisticApp) error {

    var err error
    var status bool

	// Check number of users defined in DB and create a default one if needed
	if status, err = handleUsers(aa.DataProv); err != nil {
		return fmt.Errorf("Cannot check users in database: '%s'", err.Error())
	}
    if status {
        aa.Log.Info("Default user created: 'admin/admin123!'")
        fmt.Println("IMPORTANT: Default user created: 'admin/admin123!'; please change at first login!")
    }

    // Check the number of Datings in DB and create the records if needed
    if status, err = handleDatings(aa.DataProv); err != nil {
		return fmt.Errorf("Cannot check datings in database: '%s'", err.Error())
    }
    if status {
	    aa.Log.Info("Dating records created")
    }

    return err
}

// The handleUsers function checks the number of users in DB at startup and creates a default user when none is present. 
// Otherwise noone could login into app and do anything... 
func handleUsers(dp db.DataProvider) (bool, error) {

	var cnt int
	var err error
    status := false

	if cnt, err = dp.CountUsers(); err != nil {
		return status, err
	}

	// if database contains no users, create a default one...
	if cnt == 0 {
		u := db.CreateUser(db.DefAppUsername, db.DefAppPasswd, "admin", true)
		u.Fullname = "Change Myname"
		u.Email = "change_me@somewhere.org"
		err = dp.InsertUser(u)
        status = true
	}
	return status, err
}

// The handleDatings function checks the number of Datings in DB at startup and creates the records when empty.
// Since there are only 8 records and those are fixed, there's no need to mess things up in application. Records can
// otherwise be updated in application, but that all.
func handleDatings(dp db.DataProvider) (bool, error) {

	var cnt int
	var err error
    status := false

	if cnt, err = dp.CountDatings(); err != nil {
		return status, err
	}

	if cnt == 0 {
		var datings []*db.Dating
		datings = append(datings, db.NewDating(&core.Dating{"L", "L description"}))
		datings = append(datings, db.NewDating(&core.Dating{"S", "S description"}))
		datings = append(datings, db.NewDating(&core.Dating{"A", "A description"}))
		datings = append(datings, db.NewDating(&core.Dating{"a.q.", "a.q. description"}))
		datings = append(datings, db.NewDating(&core.Dating{"a.q.n.", "a.q.n. description"}))
		datings = append(datings, db.NewDating(&core.Dating{"p.q.", "p.q. description"}))
		datings = append(datings, db.NewDating(&core.Dating{"p.q.n.", "p.q.n. description"}))
		datings = append(datings, db.NewDating(&core.Dating{"unknown", "Unknown dating"}))

		err = dp.InsertDatings(datings)
        status = true
	}
	return status, err
}

