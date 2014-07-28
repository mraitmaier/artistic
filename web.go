/*
   web.go
*/
package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
    //    "strconv"
	//    "bytes"
	//    "encoding/binary"
	"html/template"
	"net/http"
	"path/filepath"
	"bitbucket.org/miranr/artistic/core"
	"bitbucket.org/miranr/artistic/utils"
//	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/gorilla/mux"
)

type WebInfo struct {

	// a path to where session files are stored
	sessDir string

	// cached page templates
	templates *template.Template

	// web session cookie store
	store *sessions.CookieStore

    // websocket connection
    //wsConn *websocket.Conn
}

const (
	// a (quite) random string that is used as a key for sessions
	sessKey = `iufwnwieh3436SiKJSJo90e3jdiejdlje3+'0%$#!)dlkjja(!~§<sdfad$io*"`

    // context key
    //LoggedUser string = "user"
)

var (
	// Favicon location
	favicon = "web/static/favicon.ico"
)

// register web page handler functions
func registerHandlers(aa *ArtisticApp) {

    r := mux.NewRouter()
	r.Handle("/", indexHandler(aa) )
	r.Handle("/login", loginHandler(aa) )
	r.Handle("/logout", logoutHandler(aa) )
	r.Handle("/index", indexHandler(aa) )
	r.Handle("/users", usersHandler(aa))
	r.Handle("/techniques", techniquesHandler(aa) )
	r.Handle("/styles", stylesHandler(aa) )
	r.Handle("/datings", datingsHandler(aa) )
	r.Handle("/dating/{id}/{cmd}", datingHandler(aa) )
	r.Handle("/error404", err404Handler(aa) )
	r.Handle("/license", licenseHandler(aa) )
	r.HandleFunc("/favicon.ico", faviconHandler)
    // websocket handler
    //r.Handle("/ws", wsHandler(aa) )
    //r.Handle("/wss", wsHandler(aa) )
    r.NotFoundHandler = err404Handler(aa)

    // Call the default URL router...
    http.Handle("/", r)
}

// initializes and starts web server
func webStart(aa *ArtisticApp, wwwpath string) error {

	aa.WebInfo = new(WebInfo)

    // create new session cookie store
	aa.WebInfo.store = sessions.NewCookieStore([]byte(sessKey))

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
	aa.WebInfo.templates = template.Must(
            template.New("").Funcs(funcs).ParseGlob(t))

	// finally, start web server, we're using HTTPS
	// http.ListenAndServe(":8088", context.ClearHandler(http.DefaultServeMux))
	http.ListenAndServeTLS(":8088", "./web/static/cert.pem",
		"./web/static/key.pem", nil)
//		"./web/static/key.pem", context.ClearHandler(http.DefaultServeMux))
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
		aa.WebInfo.sessDir = filepath.Join(basedir, "artistic", "sessions")
	}

	// if path does not exits, create it...
	if err := os.MkdirAll(aa.WebInfo.sessDir, 0755); err != nil {
		fmt.Println("FATAL: Cannot create path. Cannot continue...")
		//fmt.Println(err.Error()) // DEBUG
		return false
	}

	return true
}

// remove all session files (and the session folder itself!) when app is
// terminated.
func cleanSessDir(aa *ArtisticApp) bool {

	status := false

	if aa.WebInfo.sessDir != "" {
		if err := os.RemoveAll(aa.WebInfo.sessDir); err != nil {
			aa.Log.Error(err.Error())
			return status
		}
		status = true
	}
	return status
}

// HTTP error 404 page handler
func err404Handler(aa *ArtisticApp) http.Handler {

    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		// render the page
		err := aa.WebInfo.templates.ExecuteTemplate(w, "error404", user)
        if err != nil {
			aa.Log.Error("Error rendering the '404' page.")
		}

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

    }) // return handler closure
}

// logout handler
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
	http.Redirect(w, r, "/login", http.StatusFound)
    } ) // return handler closure
}

// license page handler
func licenseHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc(func  (w http.ResponseWriter, r *http.Request) {

	 if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		// render the page
		err := aa.WebInfo.templates.ExecuteTemplate(w, "license", user)
        if err != nil {
			aa.Log.Error("Cannot render the 'license' page.")
		}

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

    } ) // return handler closure
}

// Index (home) page handler
func indexHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

   //     u := context.Get(r, LoggedUser)
   //     fmt.Printf("DEBUG context=%v\n", u) // DEBUG

		err := aa.WebInfo.templates.ExecuteTemplate(w, "index", user)
        if err != nil {
			aa.Log.Error("Cannot render the 'index' page.")
		}

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
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
			http.Redirect(w, r, "/error404", http.StatusFound)
			return
		}

		/* DEBUG
		   for _, val := range users {
		       fmt.Printf("User: %s\n", val.String()) // DEBUG
		   }
		*/

		// create ad-hoc struct to be sent to page template
		var web = struct {
			User  *utils.User
			Users []utils.User
		} { user, users }

		// render the page
		err = aa.WebInfo.templates.ExecuteTemplate(w, "users", &web)
        if err != nil {
			log.Error("Cannot render the 'users' page.")
		}

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
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
			err = aa.WebInfo.templates.ExecuteTemplate(w, "login", nil)
            if err != nil { log.Error(err.Error()) }
			log.Alert(fmt.Sprintf("User %q NOT authenticated.", user))
		}

		// if authenticated, redirect to index page; otherwise display login
		if status {

           // context.Set(r, LoggedUser, user)
			http.Redirect(w, r, "/index", http.StatusFound)
		}
		log.Info(fmt.Sprintf("User %q authenticated, OK.", user))

	// when HTTP GET is received, just display the default login template
	case "GET":
		err := aa.WebInfo.templates.ExecuteTemplate(w, "login", nil)
        if err != nil {
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
		datings, err := aa.DataProv.GetAllDatings()
		if err != nil {
			log.Error(fmt.Sprintf("Problem getting all datings: %s",
				err.Error()))
			http.Redirect(w, r, "/error404", http.StatusFound)
			return
		}

 //       ghhjhg, _ := core.DatingsToJson(datings)

		// create ad-hoc struct to be sent to page template
        var web = struct {
			User    *utils.User
			Datings []core.Dating
//            Json string
 //       } { user, datings, ghhjhg}
        } { user, datings}

		// render the page
		err = aa.WebInfo.templates.ExecuteTemplate(w, "datings", &web)
        if err != nil {
			aa.Log.Error("Error rendering the 'datings' page.")
		}

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
    }) // return handler closure
}

// Handle request for single dating: view & modify operations.
func datingHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		log := aa.Log // get logger instance

        id := mux.Vars(r)["id"]
        cmd := mux.Vars(r)["cmd"]

		// get a dating with given ID from DB
		dating, err := aa.DataProv.GetDating(id)
		if err != nil {
			log.Error(fmt.Sprintf("Problem getting a dating %q: %q",
				id, err.Error()))
			http.Redirect(w, r, "/error404", http.StatusFound)
			return
		}

		// create ad-hoc struct to be sent to page template
        var web = struct {
			User    *utils.User
            Cmd     string // "view" or "modify"; we don't allow "delete"...
			Dating  *core.Dating
        } { user, cmd, dating }

		// render the page
        //fmt.Printf("DEBUG r='%v'\n", r) // DEBUG
		err = aa.WebInfo.templates.ExecuteTemplate(w, "dating", &web)
        if err != nil {
            msg := fmt.Sprintf("Error rendering the 'dating' page: %q.\n",
                err.Error())
			aa.Log.Error(msg)
		}

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
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
			log.Error(
                fmt.Sprintf("Problem getting all styles: %s", err.Error()))
			http.Redirect(w, r, "/error404", http.StatusFound)
			return
		}

		// create ad-hoc struct to be sent to page template
		var web = struct {
			User   *utils.User
			Styles []core.Style
		} { user, styles }

		// render the page
		err = aa.WebInfo.templates.ExecuteTemplate(w, "styles", &web)
        if err != nil {
			aa.Log.Error("Error rendering the 'styles' page.")
		}

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
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
			http.Redirect(w, r, "/error404", http.StatusFound)
			return
		}

		// create ad-hoc struct to be sent to page template
        var web = struct {
			User       *utils.User
			Techniques []core.Technique
        }{ user, tech }

		// render the page
		//err = aa.WebInfo.templates.ExecuteTemplate(w, "techniques", web)
		err = aa.WebInfo.templates.ExecuteTemplate(w, "techniques", &web)
        if err != nil {
			aa.Log.Error("Error rendering the 'techniques' page.")
		}

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
    }) // return handler closure
}

/*
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
        aa.WebInfo.wsConn = ws

        fmt.Printf("DEBUG websocket: %v\n", ws)
	} else {
		http.Redirect(w, r, "login", http.StatusFound)
	}

    }) // return handler closure
}
*/
