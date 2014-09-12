//
//   web.go
//
package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
    "strings"
	"html/template"
	"net/http"
	"path/filepath"
	"bitbucket.org/miranr/artistic/core"
	"bitbucket.org/miranr/artistic/utils"
	"bitbucket.org/miranr/artistic/db"
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

    // (error) message to be displayed on page
    LastMsg string
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
	r.Handle("/user/{cmd}/{id}", userHandler(aa) )
	r.Handle("/user/{cmd}/", userHandler(aa) )
	r.Handle("/techniques", techniquesHandler(aa) )
	r.Handle("/technique/{cmd}/{id}", techniqueHandler(aa) )
	r.Handle("/technique/{cmd}/", techniqueHandler(aa) )
	r.Handle("/styles", stylesHandler(aa) )
	r.Handle("/style/{cmd}/{id}", styleHandler(aa) )
	r.Handle("/style/{cmd}/", styleHandler(aa) )
	r.Handle("/datings", datingsHandler(aa) )
	r.Handle("/dating/{cmd}/{id}", datingHandler(aa) )
	r.Handle("/error404", err404Handler(aa) )
	r.Handle("/error", errorHandler(aa) )
	r.Handle("/license", licenseHandler(aa) )
    r.Handle("/userprofile", profileHandler(aa) )
    r.Handle("/userprofile/{cmd}", profileHandler(aa) )
    r.Handle("/userprofile/{cmd}/{id}", profileHandler(aa) )
    r.Handle("/artists", artistsHandler(aa, core.ArtistTypeArtist) )
    r.Handle("/painters", artistsHandler(aa, core.ArtistTypePainter) )
    r.Handle("/sculptors", artistsHandler(aa, core.ArtistTypeSculptor) )
    r.Handle("/printmakers", artistsHandler(aa, core.ArtistTypePrintmaker) )
    r.Handle("/ceramicists", artistsHandler(aa, core.ArtistTypeCeramicist) )
    r.Handle("/architects", artistsHandler(aa, core.ArtistTypeArchitect) )
    r.Handle("/artist/{cmd}/{id}", artistHandler(aa) )
    r.Handle("/artist/{cmd}/", artistHandler(aa) )
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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))

	//web page templates, with defined additional functions
	funcs := template.FuncMap{
		"add": func(x, y int) int { return x + y },
        "allowedroles": func() []string { return utils.AllowedRoles },
        "get_artist_type": func(t core.ArtistType) string { return t.String() },
        "totitle": func(s string) string { return strings.Title(s) },
        "toupper": func(s string) string { return strings.ToUpper(s) },
        "tolower": func(s string) string { return strings.ToLower(s) }}
	t := filepath.Join(wwwpath, "templates", "*.tpl")
	aa.WebInfo.templates = template.Must(template.New("").Funcs(funcs).ParseGlob(t))

	// finally, start web server, we're using HTTPS
	// http.ListenAndServe(":8088", context.ClearHandler(http.DefaultServeMux))
	http.ListenAndServeTLS(":8088", "./web/static/cert.pem", "./web/static/key.pem", nil)
//		"./web/static/key.pem", context.ClearHandler(http.DefaultServeMux))
	return nil
}

func SetMessage(wi *WebInfo, msg string) {
    wi.LastMsg = msg
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

// aux function redirecting to login page.
func redirectToLoginPage(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) {

    aa.Log.Warning("User not authenticated. ")
	http.Redirect(w, r, "/login", http.StatusFound)
}

/*
// aux function redirecting to error page.
func redirectToErrorPage(name string, user *utils.User, w http.ResponseWriter, r *http.Request, aa *ArtisticApp) {

	aa.Log.Error(fmt.Sprintf("[%s] Cannot render the %q template.", name, user.Username))
	http.Redirect(w, r, "/error", http.StatusFound)
}
*/

// Aux function that renders the page (template!) with given (template) name.
// Input parameters are:
// - name - name of the template to render
// - user - instance of User type
// - web  - ptr to ad-hoc web struct that is used by template to fill in the data on page
// - aa   - instance of ArtisticApp type
// - w    - ptr to the ResponseWriter type instance
// - r    - ptr to the (HTTP) Request type instance
func renderPage(name string, web interface{}, aa *ArtisticApp, w http.ResponseWriter, r *http.Request) error {

    var err error
    if err = aa.WebInfo.templates.ExecuteTemplate(w, name, web); err != nil {
        http.Redirect(w, r, "/error", http.StatusFound)
    }
    return err
}

// HTTP error 404 page handler
func err404Handler(aa *ArtisticApp) http.Handler {

    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

        if err := renderPage("error404", user, aa, w, r); err != nil {
	        aa.Log.Error(fmt.Sprintf("[%s] Cannot render the 'error404' template: %q.", user.Username, err.Error()))
        }
        aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

	} else {
		redirectToLoginPage(w, r, aa)
	}

    }) // return handler closure
}

// error page handler
// this one should be displayed when template  cannot be rendered or in case when unknown error occurs.
func errorHandler(aa *ArtisticApp) http.Handler {

    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		// render the page
		if err := aa.WebInfo.templates.ExecuteTemplate(w, "error", user); err != nil {
			aa.Log.Error(fmt.Sprintf("[%s] Rendering the 'error' template.", user.Username))
            http.Redirect(w, r, "/error404", http.StatusFound) // This is really worst-case scenario...
		}

	} else {
        redirectToLoginPage(w, r, aa)
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
        if err := renderPage("license", user, aa, w, r); err != nil {
	        aa.Log.Error(fmt.Sprintf("[%s] Cannot render the %q template: %q.", user.Username, "license", err.Error()))
            return
        }
        aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))
	} else {
        redirectToLoginPage(w, r, aa)
	}

    } ) // return handler closure
}

// Index (home) page handler
func indexHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

        if err := renderPage("index", user, aa, w, r); err != nil {
	        aa.Log.Error(fmt.Sprintf("[%s] Cannot render the 'index' template: %q.", user.Username, err.Error()))
            return
        }
        aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))
    } else {
        redirectToLoginPage(w, r, aa)
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
            aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user, r.RequestURI))
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
			log.Error("Rendering the 'login' template.")
		}
        aa.Log.Info(fmt.Sprintf("Displaying the %q page.", r.RequestURI))
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
			log.Error(fmt.Sprintf("[%s] Problem getting all datings: %s", user.Username, err.Error()))
			http.Redirect(w, r, "/error", http.StatusFound)
			return
		}

		// create ad-hoc struct to be sent to page template
        var web = struct {
			User    *utils.User
			Datings []core.Dating
        } { user, datings }

        if err := renderPage("datings", &web, aa, w, r); err != nil {
	        aa.Log.Error(fmt.Sprintf("[%s] Cannot render the 'datings' template: %q.", user.Username, err.Error()))
            return
        }
        aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

    } else {
        redirectToLoginPage(w, r, aa)
	}
    }) // return handler closure
}

// POST request handler for datings (private).
func postDatingHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) error {

    // get data to modify 
	id := mux.Vars(r)["id"]
	name := r.FormValue("dating-name")
	descr := r.FormValue("dating-description")

    d := &core.Dating {db.MongoStringToId(id), name, descr}

    if err := aa.DataProv.UpdateDating(d); err != nil {
        return err
    }
    aa.Log.Info(fmt.Sprintf("Successfully updated dating %q\n", name))

    return nil
}

// GET request handler for datings (private).
func getDatingHandler(w http.ResponseWriter, r *http.Request,
                    aa *ArtisticApp, user *utils.User) error {

    id := mux.Vars(r)["id"]
    cmd := mux.Vars(r)["cmd"]

	// get a dating with given ID from DB
	dating, err := aa.DataProv.GetDating(id)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusFound)
		return fmt.Errorf("Problem getting a dating from DB: %q", err.Error())
	}

	// create ad-hoc struct to be sent to page template
    var web = struct {
		User    *utils.User
        Cmd     string // "view" or "modify"; we don't allow "delete"...
		Dating  *core.Dating
    } { user, cmd, dating }

    return renderPage("dating", &web, aa, w, r)
}

// Handle request for single dating: view & modify operations.
func datingHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

        log := aa.Log // get logger instance

        switch r.Method {

        case "GET":
            if err := getDatingHandler(w, r, aa, user); err != nil {
			    log.Error(fmt.Sprintf("[%s] Dating GET handler: %s.", user.Username, err.Error()))
                break // force break!
            }
            aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

        case "POST":
            if err := postDatingHandler(w, r, aa); err != nil{
			    log.Error(fmt.Sprintf("[%s] Dating POST handler: %s.", user.Username, err.Error()))
            }
		    http.Redirect(w, r, "/datings", http.StatusFound)

        } // switch

	} else {
        redirectToLoginPage(w, r, aa)
	}
    }) // return handler closure
}

// styles page handler
func stylesHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		log := aa.Log // get logger instance

		// get all styles from DB
		styles, err := aa.DataProv.GetAllStyles()
		if err != nil {
			log.Error(fmt.Sprintf("Problem getting all styles: %s.", err.Error()))
			http.Redirect(w, r, "/error", http.StatusFound)
			return
		}

		// create ad-hoc struct to be sent to page template
		var web = struct {
			User   *utils.User
			Styles []core.Style
		} { user, styles }

		// render the page
        if err = renderPage("styles", &web, aa, w, r); err != nil {
	        aa.Log.Error(fmt.Sprintf("[%s] Cannot render the 'styles' template: %q.", user.Username, err.Error()))
            return
        }
        aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

	} else {
        redirectToLoginPage(w, r, aa)
	}
    }) // return handler closure
}

// a single style handler
func styleHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

	    log := aa.Log // get logger instance

        switch r.Method {

        case "GET":
            if err := getStyleHandler(w, r, aa, user); err != nil {
                log.Error(fmt.Sprintf("[%s] Style GET handler: %s.", user.Username, err.Error()))
                break // force break!
            }
            log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

        case "POST":
            if err := postStyleHandler(w, r, aa, user); err != nil {
                log.Error(fmt.Sprintf("[%s] Style POST handler: %s.", user.Username, err.Error()))
            }
			http.Redirect(w, r, "/styles", http.StatusFound)

        case "DELETE":
            id := mux.Vars(r)["id"]
            //cmd := mux.Vars(r)["cmd"]
            t := new(core.Style)
            t.Id = db.MongoStringToId(id) // only valid ID needed to delete 
            if err := aa.DataProv.DeleteStyle(t); err != nil {
                log.Error(fmt.Sprintf("[%s] DELETE Style id=%q (%s).", user.Username, id, err.Error()))
                return
            }
            log.Info(fmt.Sprintf("[%s] Successfully deleted style %q.", user.Username, t.Name))
	        http.Redirect(w, r, "/styles", http.StatusFound)

        case "PUT":
            log.Warning(fmt.Sprintf("[%s] Received PUT request. :).", user.Username))
        }

	} else {
        redirectToLoginPage(w, r, aa)
	}
    }) // return handler closure
}

func getStyleHandler(w http.ResponseWriter, r *http.Request,
                        aa *ArtisticApp, user *utils.User) error {

    id := mux.Vars(r)["id"]
    cmd := mux.Vars(r)["cmd"]

    log := aa.Log
    var err error
    s := new(core.Style)

    switch cmd {

    case "view", "modify":
	    // get a style from DB
	    s, err = aa.DataProv.GetStyle(id)
	    if err != nil {
		    return fmt.Errorf("%s Style id=%q: %s.", strings.ToUpper(cmd), id, err.Error())
	    }

    case "insert":
        // do nothing here...

    case "delete":
        s.Id = db.MongoStringToId(id) // only valid ID needed to delete 
        if err = aa.DataProv.DeleteStyle(s); err != nil {
            return fmt.Errorf("DELETE Style id=%q: %s.", cmd, id, err.Error())
        }
        log.Info(fmt.Sprintf("Successfully deleted style %q.", s.Name))
	    http.Redirect(w, r, "/styles", http.StatusFound)
        return nil //  this is all about deleting items...

    default:
        return fmt.Errorf("Unknown command %q", cmd)
    }

	// create ad-hoc struct to be sent to page template
    var web = struct {
		User  *utils.User
        Cmd   string        // "view", "modify", "create" or "delete"...
		Style *core.Style
    }{ user, cmd, s }

    return renderPage("style", &web, aa, w, r)
}

func postStyleHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *utils.User) error {

    // get data to modify 
	id := mux.Vars(r)["id"]
    cmd := mux.Vars(r)["cmd"]

    // get POST form values and create a struct
	name := strings.TrimSpace(r.FormValue("style-name"))
	descr := strings.TrimSpace(r.FormValue("style-description"))
    t := &core.Style{ db.MongoStringToId(id), name, descr }

    var err error

    switch cmd {

    case "insert":
        if err = aa.DataProv.InsertStyle(t); err != nil {
            return err
        }
        aa.Log.Info(fmt.Sprintf("[%s] Style INSERT %q Success.", user.Username, name))

    case "modify":
        if err = aa.DataProv.UpdateStyle(t); err != nil {
            return err
        }
        aa.Log.Info(fmt.Sprintf("[%s] Style MODIFY %q Success.", user.Username, name))

    default:
	    http.Redirect(w, r, "/styles", http.StatusFound)
        err = fmt.Errorf("Unknown command %q", strings.ToUpper(cmd))
    }

    return err
}

// techniques page handler
func techniquesHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		log := aa.Log // get logger instance

		// get all techniques from DB
		tech, err := aa.DataProv.GetAllTechniques()
		if err != nil {
			log.Error(fmt.Sprintf("[%s] Problem getting all techniques: %s.", user.Username, err.Error()))
			http.Redirect(w, r, "/error", http.StatusFound)
			return
		}

		// create ad-hoc struct to be sent to page template
        var web = struct {
			User       *utils.User
			Techniques []core.Technique
        }{ user, tech }

		// render the page
		//err = aa.WebInfo.templates.ExecuteTemplate(w, "techniques", &web)
        if err = renderPage("techniques", &web, aa, w, r); err != nil {
			log.Error(fmt.Sprintf("[%s] Rendering the 'techniques' template: %s.", user.Username, err.Error()))
            return
		}
        log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

	} else {
        redirectToLoginPage(w, r, aa)
	}
    }) // return handler closure
}


// a single technique handler
func techniqueHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

	    log := aa.Log // get logger instance

        switch r.Method {

        case "GET":
            if err := getTechniqueHandler(w, r, aa, user); err != nil {
                log.Error(fmt.Sprintf("[%s] Technique GET handler: %s.", user.Username, err.Error()))
                break // force break!
            }
            log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

        case "POST":
            if err := postTechniqueHandler(w, r, aa, user); err != nil {
                log.Error(fmt.Sprintf("[%s] Technique POST handler: %s.", user.Username, err.Error()))
            }
			http.Redirect(w, r, "/techniques", http.StatusFound)

        case "DELETE":
            id := mux.Vars(r)["id"]
            //cmd := mux.Vars(r)["cmd"]
            t := new(core.Technique)
            t.Id = db.MongoStringToId(id) // only valid ID needed to delete 
            if err := aa.DataProv.DeleteTechnique(t); err != nil {
                msg := fmt.Sprintf("[%s] DELETE Technique id=%q: %q.", user.Username, id, err.Error())
                log.Error(msg)
                return
            }
            log.Info(fmt.Sprintf("[%s] DELETE Technique %q Success.", t.Name))
	        http.Redirect(w, r, "/techniques", http.StatusFound)

        case "PUT":
            log.Warning("Received HTTP PUT request. :).")
        }

	} else {
		redirectToLoginPage(w, r, aa)
	}
    }) // return handler closure
}

func getTechniqueHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *utils.User) error {

    id := mux.Vars(r)["id"]
    cmd := mux.Vars(r)["cmd"]

    log := aa.Log
    var err error
    tech := new(core.Technique)

    switch cmd {

    case "view", "modify":
	    // get a technique from DB
	    tech, err = aa.DataProv.GetTechnique(id)
	    if err != nil {
		    return fmt.Errorf("%s Technique id=%q: %s", strings.ToUpper(cmd), id, err)
	    }

    case "insert":
        // do nothing here...

    case "delete":
        tech.Id = db.MongoStringToId(id) // only valid ID needed to delete 
        if err = aa.DataProv.DeleteTechnique(tech); err != nil {
            return fmt.Errorf("DELETE Technique id=%q: %q.", id, err.Error())
        }
        log.Info(fmt.Sprintf("[%s] DELETE Technique %q Success.", tech.Name))
	    http.Redirect(w, r, "/techniques", http.StatusFound)
        return nil //  this is all about deleting items...

    default:
        return fmt.Errorf("Unknown command %q", cmd)
    }

	// create ad-hoc struct to be sent to page template
    var web = struct {
		User      *utils.User
        Cmd       string        // "view", "modify", "insert" or "delete"...
		Technique *core.Technique
    }{ user, cmd, tech }

    return renderPage("technique", &web, aa, w, r)
}

func postTechniqueHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *utils.User) error {

    // get data to modify 
	id := mux.Vars(r)["id"]
    cmd := mux.Vars(r)["cmd"]

    // get POST form values and create a struct
	name := strings.TrimSpace(r.FormValue("technique-name"))
	descr := strings.TrimSpace(r.FormValue("technique-description"))
    t := &core.Technique {db.MongoStringToId(id), name, descr}

    var err error = nil

    switch cmd {

    case "insert":
        if err = aa.DataProv.InsertTechnique(t); err != nil {
            return err
        }
        aa.Log.Info(fmt.Sprintf("[%s] INSERT Technique %q Success.", user.Username, name))

    case "modify":
        if err = aa.DataProv.UpdateTechnique(t); err != nil {
            return err
        }
        aa.Log.Info(fmt.Sprintf("[%s] MODIFY Technique %q Success.", user.Username, name))

    default:
	    http.Redirect(w, r, "/techniques", http.StatusFound)
        err = fmt.Errorf("Unknown command %q", cmd)
    }

    return err
}
