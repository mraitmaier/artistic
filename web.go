/*
   web.go
*/
package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	//    "bytes"
	//    "encoding/binary"
	"html/template"
	"net/http"
	"path/filepath"
	//    "labix.org/v2/mgo/bson"
    //   "labix.org/v2/mgo"
	"bitbucket.org/miranr/artistic/core"
	"bitbucket.org/miranr/artistic/utils"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
)

type WebInfo struct {

	// a path to where session files are stored
	SessDir string

    // websocket connection
    WsConn *websocket.Conn
}

const (
	// a (quite) random string that is used as a key for sessions
	sessKey = `iufwnwieh3436SiKJSJo90e3jdiejdlje3+'0%$#!)dlkjja(!~§<sdfad$io*"`

    // context key
    //LoggedUser string = "user"
)

var (
	// global var holding cached page templates
	templates *template.Template

	// global var holding the web session data
	store = sessions.NewCookieStore([]byte(sessKey))

	// Favicon location
	favicon = "/web/static/favicon.ico"
)

// register web page handler functions
func registerHandlers(aa *ArtisticApp) {

	http.Handle("/", indexHandler(aa) )
	http.Handle("/login", loginHandler(aa) )
	http.Handle("/logout", logoutHandler(aa) )
	http.Handle("/index", indexHandler(aa) )
	http.Handle("/users", usersHandler(aa))
	http.Handle("/techniques", techniquesHandler(aa) )
	http.Handle("/styles", stylesHandler(aa) )
	http.Handle("/datings", datingsHandler(aa) )
	http.Handle("/error404", err404Handler(aa) )
	http.Handle("/license", licenseHandler(aa) )
	http.HandleFunc("/favicon.ico", faviconHandler)
    //
    http.Handle("/ws", wsHandler(aa) )
}

// initializes and starts web server
func webStart(aa *ArtisticApp, wwwpath string) error {

	aa.WebInfo = new(WebInfo)

	// register handler functions
	registerHandlers(aa)

	// check dir for session files and create it if needed
	if !checkSessDir("", aa) {
		return errors.New("Cannot create session folder; cannot continue.")
	}

	// handle static files
	path := filepath.Join(wwwpath, "static")
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir(path))))

	//web page templates, with defined additional functions
	funcs := template.FuncMap{
		"add": func(x, y int) int { return x + y }}
	t := filepath.Join(wwwpath, "templates", "*.tpl")
	templates = template.Must(template.New("").Funcs(funcs).ParseGlob(t))

	// finally, start web server, we're using HTTPS
	// http.ListenAndServe(":8088", context.ClearHandler(http.DefaultServeMux))
	http.ListenAndServeTLS(":8088", "./web/static/cert.pem",
		"./web/static/key.pem", context.ClearHandler(http.DefaultServeMux))
	return nil
}

func checkSessDir(path string, aa *ArtisticApp) bool {

	basedir := path
	// if given base path is empty, default to temp folder
	if path == "" {
		switch runtime.GOOS {
		case "windows":
			basedir = os.Getenv("TEMP")
		default:
			basedir = "/tmp"
		}
		aa.WebInfo.SessDir = filepath.Join(basedir, "artistic", "sessions")
	}

	// if path does not exits, create it...
	if err := os.MkdirAll(aa.WebInfo.SessDir, 0755); err != nil {
		fmt.Println("FATAL: Cannot create path. Cannot continue...")
		fmt.Println(err.Error()) // DEBUG
		return false
	}

	return true
}

// remove all session files (and the session folder itself!) when app is
// terminated.
func cleanSessDir(aa *ArtisticApp) bool {

	status := false

	if aa.WebInfo.SessDir != "" {
		if err := os.RemoveAll(aa.WebInfo.SessDir); err != nil {
			aa.Log.Error(err.Error())
			return status
		}
		status = true
	}
	return status
}

// user admin page handler
func err404Handler(aa *ArtisticApp) http.Handler {

    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		// render the page
		if err := templates.ExecuteTemplate(w, "error404", user); err != nil {
			aa.Log.Error("Error rendering the '404' page.")
		}

	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}

    }) // return handler closure
}

// user admin page handler
func logoutHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		log := aa.Log

		// render the page
		if err := logout(aa, w, r); err != nil {
			log.Error(err.Error())
		} else {
			log.Info(fmt.Sprintf("Logging out user %q.", user.Username))
		}
	}
	http.Redirect(w, r, "login", http.StatusFound)
    } ) // return handler closure
}

// license page handler
func licenseHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc(func  (w http.ResponseWriter, r *http.Request) {

	 if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		// render the page
		if err := templates.ExecuteTemplate(w, "license", user); err != nil {
			aa.Log.Error("Cannot render the 'license' page.")
		}

	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}

    }) // return handler closure
}

// Index (home) page handler
func indexHandler(aa *ArtisticApp) http.Handler{
    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

   //     u := context.Get(r, LoggedUser)
   //     fmt.Printf("DEBUG context=%v\n", u) // DEBUG

		if err := templates.ExecuteTemplate(w, "index", user); err != nil {
			aa.Log.Error("Cannot render the 'index' page.")
		}

	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}
    }) // return handler closure
}

// user admin page handler
func usersHandler(aa *ArtisticApp) http.Handler {

    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		log := aa.Log

		// get all users from DB
		users, err := aa.DataProv.GetAllUsers()
		if err != nil {
			log.Error(fmt.Sprintf("Problem getting all users: %s", err.Error()))
			http.Redirect(w, r, "error404", http.StatusFound)
			return
		}

		/* DEBUG
		   for _, val := range users {
		       fmt.Printf("User: %s\n", val.String()) // DEBUG
		   }
		*/

		// create temp struct variable to be sent to page template
		web := new(struct {
			User  *utils.User
			Users []utils.User
		})
		web.User = user
		web.Users = users

		// render the page
		if err := templates.ExecuteTemplate(w, "users", web); err != nil {
			log.Error("Cannot render the 'users' page.")
		}

	} else {
		http.Redirect(w, r, "login.html", http.StatusFound)
	}
    }) // return handler closure
}

// login page handler - we must authenticate user
func loginHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

	log := aa.Log // get logger instance

	switch r.Method {

	// when HTTP POST is received...
	case "POST":

		user := r.FormValue("username")
		pwd := r.FormValue("password")

		// authenticate user
		log.Info(fmt.Sprintf("Trying to authenticate user %q...", user))
		status, err := authenticateUser(user, pwd, aa, w, r)
		if !status || err != nil {
			if err = templates.ExecuteTemplate(w, "login", nil); err != nil {
				log.Error(err.Error())
			}
			log.Alert(fmt.Sprintf("User %q NOT authenticated.", user))
		}

		// if authenticated, redirect to index page; otherwise display login
		if status {

           // context.Set(r, LoggedUser, user)
			http.Redirect(w, r, "index", http.StatusFound)
		}
		log.Info(fmt.Sprintf("User %q authenticated, OK.", user))

	// when HTTP GET is received, just display the default login template
	case "GET":
		if err := templates.ExecuteTemplate(w, "login", nil); err != nil {
			log.Error("Cannot render the 'login' page.")
		}
	}

    }) // return handler closure
}

// favincon handler
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, favicon)
}

// datings page handler
func datingsHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		log := aa.Log // get logger instance

		// get all datings from DB
		//datings, err := dbase.MongoGetAllDatings(aa.DbSess.DB("artistic"))
		datings, err := aa.DataProv.GetAllDatings()
		if err != nil {
			log.Error(fmt.Sprintf("Problem getting all datings: %s",
				err.Error()))
			http.Redirect(w, r, "error404", http.StatusFound)
			return
		}

		// create temp struct variable to be sent to page template
		web := new(struct {
			User    *utils.User
			Datings []core.Dating
		})
		web.User = user
		web.Datings = datings

		// render the page
		if err := templates.ExecuteTemplate(w, "datings", web); err != nil {
			aa.Log.Error("Error rendering the 'datings' page.")
		}

	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}
    }) // return handler closure
}

// styles page handler
func stylesHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		log := aa.Log // get logger instance

		// get all styles from DB
		//styles, err := dbase.MongoGetAllStyles(aa.DbSess.DB("artistic"))
		styles, err := aa.DataProv.GetAllStyles()
		if err != nil {
			log.Error(fmt.Sprintf("Problem getting all styles: %s", err.Error()))
			http.Redirect(w, r, "error404", http.StatusFound)
			return
		}

		// create temp struct variable to be sent to page template
		web := new(struct {
			User   *utils.User
			Styles []core.Style
		})
		web.User = user
		web.Styles = styles

		// render the page
		if err := templates.ExecuteTemplate(w, "styles", web); err != nil {
			aa.Log.Error("Error rendering the 'styles' page.")
		}

	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}
    }) // return handler closure
}

// techniques page handler
func techniquesHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		log := aa.Log // get logger instance

		// get all techniques from DB
		tech, err := aa.DataProv.GetAllTechniques()
		if err != nil {
			log.Error(fmt.Sprintf("Problem getting all techniques: %s",
				err.Error()))
			http.Redirect(w, r, "error404", http.StatusFound)
			return
		}

		// create temp struct variable to be sent to page template
		web := new(struct {
			User       *utils.User
			Techniques []core.Technique
		})
		web.User = user
		web.Techniques = tech

		// render the page
		if err := templates.ExecuteTemplate(w, "techniques", web); err != nil {
			aa.Log.Error("Error rendering the 'techniques' page.")
		}

	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}
    }) // return handler closure
}

//
const wsBuffer int = 1024
func wsHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	//if loggedin, user := userIsAuthenticated(aa, r); loggedin {
	if loggedin, _ := userIsAuthenticated(aa, r); loggedin {

        if r.Method != "GET" {
            http.Error(w, "Method not allowed", 405)
            return
        }
        if r.Header.Get("Origin") != "http://" + r.Host {
            http.Error(w, "Origin not allowed", 403)
            return
        }

		log := aa.Log // get logger instance

        ws, err := websocket.Upgrade(w, r, nil, wsBuffer, wsBuffer)
        if _, ok := err.(websocket.HandshakeError); ok {
            http.Error(w, "Not a websocket handshake", 400)
            return
        } else if err != nil {
            log.Error(err.Error())
            return
        }
        aa.WebInfo.WsConn = ws

        fmt.Printf("DEBUG websocket: %v\n", ws)

	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}

    }) // return handler closure
}
