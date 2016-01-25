package main

//
//   web.go
//

import (
	"github.com/mraitmaier/artistic/core"
	"github.com/mraitmaier/artistic/db"
	//"github.com/mraitmaier/artistic/utils"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// WebInfo is the main type containing the data related to display web pages.
type WebInfo struct {

	// a path to where session files are stored
	sessDir string

	// cached page templates
	templates *template.Template

	// web session cookie store
	store *sessions.CookieStore

	// websocket connection
	//wsConn *websocket.Conn

	// (error, info etc.) message to be displayed on page
	Msg WebMessage
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
	r.Handle("/", indexHandler(aa))
	r.Handle("/login", loginHandler(aa))
	r.Handle("/logout", logoutHandler(aa))
	r.Handle("/index", indexHandler(aa))
	r.Handle("/users", usersHandler(aa))
	r.Handle("/user/{cmd}/{id}", userHandler(aa))
	r.Handle("/user/{cmd}/", userHandler(aa))
	r.Handle("/log", logHandler(aa))
	r.Handle("/technique", techniqueHandler(aa))
	r.Handle("/technique/{id}/{cmd}", techniqueHandler(aa))
	//	r.Handle("/technique/{cmd}/", techniqueHandler(aa))
	r.Handle("/styles", stylesHandler(aa))
	r.Handle("/style/{cmd}/{id}", styleHandler(aa))
	r.Handle("/style/{cmd}/", styleHandler(aa))
	r.Handle("/dating", datingHandler(aa))
	r.Handle("/dating/{id}/{cmd}", datingHandler(aa))
	r.Handle("/error404", err404Handler(aa))
	r.Handle("/error", errorHandler(aa))
	r.Handle("/license", licenseHandler(aa))
	r.Handle("/userprofile", profileHandler(aa))
	r.Handle("/userprofile/{cmd}", profileHandler(aa))
	r.Handle("/userprofile/{cmd}/{id}", profileHandler(aa))
	r.Handle("/artists", artistsHandler(aa, db.ArtistTypeArtist))
	r.Handle("/painters", artistsHandler(aa, db.ArtistTypePainter))
	r.Handle("/sculptors", artistsHandler(aa, db.ArtistTypeSculptor))
	r.Handle("/printmakers", artistsHandler(aa, db.ArtistTypePrintmaker))
	r.Handle("/ceramicists", artistsHandler(aa, db.ArtistTypeCeramicist))
	r.Handle("/architects", artistsHandler(aa, db.ArtistTypeArchitect))
	r.Handle("/artist/{cmd}/{id}", artistHandler(aa))
	r.Handle("/artist/{cmd}/", artistHandler(aa))
	r.HandleFunc("/favicon.ico", faviconHandler)
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
		"add":             func(x, y int) int { return x + y },
		"length":          func(list []string) int { return len(list) },
		"allowedroles":    func() []string { return db.AllRoles },
		"get_artist_type": func(t db.ArtistType) string { return t.String() },
		"totitle":         func(s string) string { return strings.Title(s) },
		"toupper":         func(s string) string { return strings.ToUpper(s) },
		"tolower":         func(s string) string { return strings.ToLower(s) }}
	t := filepath.Join(wwwpath, "templates", "*.tpl")
	aa.WebInfo.templates = template.Must(template.New("").Funcs(funcs).ParseGlob(t))

	// finally, start web server, we're using HTTPS
	// http.ListenAndServe(":8088", context.ClearHandler(http.DefaultServeMux))
	http.ListenAndServeTLS(":8088", "./web/static/cert.pem", "./web/static/key.pem", nil)
	//		"./web/static/key.pem", context.ClearHandler(http.DefaultServeMux))
	return nil
}

// SetMessage resets the contents of the page message (that is to be displayed on page).
func SetMessage(wi *WebInfo, msgtype, msg string) {
	wi.Msg.MsgType = msgtype
	wi.Msg.MsgText = msg
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

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
	}) // return handler closure
}

// license page handler
func licenseHandler(aa *ArtisticApp) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if loggedin, user := userIsAuthenticated(aa, r); loggedin {
			if err := renderPage("license", user, aa, w, r); err != nil {
				aa.Log.Error(fmt.Sprintf("[%s] Cannot render the %q template: %q.", user.Username, "license", err.Error()))
				return
			}
			aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))
		} else {
			redirectToLoginPage(w, r, aa)
		}

	}) // return handler closure
}

// Index (home) page handler
func indexHandler(aa *ArtisticApp) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
				if err != nil {
					log.Error(err.Error())
				}
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

// log admin page handler
func logHandler(aa *ArtisticApp) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if loggedin, user := userIsAuthenticated(aa, r); loggedin {
			switch r.Method {

			case "POST":
				aa.Log.Clear() // clear the log contents
				//			http.Redirect(w, r, "/log", http.StatusFound)
				//          return

			case "GET":
				// do nothing...
			}

			// read a log file as a slice of lines
			var err error
			var contents []string
			if contents, err = readLog(aa.LogFname); err != nil {
				aa.Log.Error(fmt.Sprintf("[%s] Problem reading log file: %q", user.Username, err.Error()))
				http.Redirect(w, r, "/error", http.StatusFound)
				return
			}

			// create ad-hoc struct to be sent to page template
			var web = struct {
				User     *db.User
				Contents []string // a list of log messages
				PerPage  int      // how many log messages per page is displayed...
			}{user, contents, 25}

			if err = renderPage("log", &web, aa, w, r); err != nil {
				aa.Log.Error(fmt.Sprintf("[%s] Cannot render the 'log' template: %q.", user.Username, err.Error()))
				return
			}
			aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

		} else {
			redirectToLoginPage(w, r, aa)
		}
	}) // return handler closure
}

// favincon handler
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, favicon)
}

// styles page handler
func stylesHandler(aa *ArtisticApp) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
				User   *db.User
				Styles []*db.Style
			}{user, styles}

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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
				t := new(db.Style)
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
	aa *ArtisticApp, user *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	log := aa.Log
	var err error
	s := db.NewStyle()

	switch cmd {

	case "view", "modify":
		// get a style from DB
		s, err = aa.DataProv.GetStyle(id)
		if err != nil {
			return fmt.Errorf("%s Style id=%q: %s", strings.ToUpper(cmd), id, err.Error())
		}

	case "insert":
		// do nothing here...

	case "delete":
		s.Id = db.MongoStringToId(id) // only valid ID needed to delete
		if err = aa.DataProv.DeleteStyle(s); err != nil {
			return fmt.Errorf("DELETE Style id=%q: %s", id, err.Error())
		}
		log.Info(fmt.Sprintf("Successfully deleted style %q.", s.Name))
		http.Redirect(w, r, "/styles", http.StatusFound)
		return nil //  this is all about deleting items...

	default:
		return fmt.Errorf("Unknown command %q", cmd)
	}

	// create ad-hoc struct to be sent to page template
	var web = struct {
		User  *db.User
		Cmd   string // "view", "modify", "create" or "delete"...
		Style *db.Style
	}{user, cmd, s}

	return renderPage("style", &web, aa, w, r)
}

func postStyleHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *db.User) error {

	// get data to modify
	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	// get POST form values and create a struct
	name := strings.TrimSpace(r.FormValue("style-name"))
	descr := strings.TrimSpace(r.FormValue("style-description"))
	ts := db.NewTimestamp()
	t := &db.Style{Id: db.MongoStringToId(id), Style: core.Style{Name: name, Description: descr}, Created: ts, Modified: ts}

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

// This is handler that handler the "/dating" URL.
func datingHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = datingHTTPGetHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("Dating HTTP GET %s", err.Error()))
				}

			case "POST":
				if err = datingHTTPPostHandler(w, r, app); err != nil {
					app.Log.Error(fmt.Sprintf("Dating HTTP POST %s", err.Error()))
				}
				// unconditionally reroute to main dating page
				http.Redirect(w, r, "/dating", http.StatusFound)

			case "DELETE":
				app.Log.Info("Dating HTTP DELETE request received. Redirecting to main 'dating' page.")
				// unconditionally reroute to main test cases page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/dating", http.StatusSeeOther)

			case "PUT":
				app.Log.Info("Dating HTTP PUT request received. Redirecting to main 'dating' page.")
				// unconditionally reroute to main test cases page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/dating", http.StatusSeeOther)

			default:
				// otherwise just display main 'index' page
				if err := renderPage("index", nil, app, w, r); err != nil {
					app.Log.Error(fmt.Sprintf("Index HTTP GET %s", err.Error()))
					return
				}
			}

		} else {
			// if user not authenticated
			redirectToLoginPage(w, r, app)
		}
	})
}

// This is HTTP POST handler for datings.
func datingHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "put":
		if id == "" {
			return fmt.Errorf("Modify dating: ID is empty")
		}
		if d := parseDatingFormValues(r); d != nil {
			d.Id = db.MongoStringToId(id)
			err = app.DataProv.UpdateDating(d)
            app.Log.Info(fmt.Sprintf("Updating Dating '%s'", d.Dating))
		}

	default:
		err = fmt.Errorf("Illegal POST request for dating")
	}
	return err
}

// Helper function that parses the '/dating' POST request values and creates a new instance of Dating.
func parseDatingFormValues(r *http.Request) *db.Dating {

	name := strings.TrimSpace(r.FormValue("dating"))
	desc := strings.TrimSpace(r.FormValue("description"))
	created := strings.TrimSpace(r.FormValue("created"))

	d := db.NewDating(&core.Dating{name, desc})
	d.Created = db.Timestamp(created)
	return d
}

// This is HTTP GET handler for datings.
func datingHTTPGetHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	datings, err := app.DataProv.GetAllDatings()
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting datings from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Datings []*db.Dating
		Num     int
		User    *db.User
	}{datings, len(datings), u}
    app.Log.Info("Displaying '/dating' page")
	return renderPage("datings", web, app, w, r)
}

// WebMessage is the structure for displaying messages on web page
type WebMessage struct {

	// message type: warning, danger, success, info
	MsgType string

	// actual text message
	MsgText string
}

func (m *WebMessage) String() string {
	return fmt.Sprintf("%s: %s", m.MsgType, m.MsgText)
}

// This is handler that handler the "/technique" URL.
func techniqueHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = techniqueHTTPGetHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("Technique HTTP GET %s", err.Error()))
				}

			case "POST":
				if err = techniqueHTTPPostHandler(w, r, app); err != nil {
					app.Log.Error(fmt.Sprintf("Technique HTTP POST %s", err.Error()))
				}
				// unconditionally reroute to main technique page
				http.Redirect(w, r, "/technique", http.StatusFound)

			case "DELETE":
				app.Log.Info("Technique HTTP DELETE request received. Redirecting to main 'technique' page.")
				// unconditionally reroute to main test cases page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/technique", http.StatusSeeOther)

			case "PUT":
				app.Log.Info("Technique HTTP PUT request received. Redirecting to main 'technique' page.")
				// unconditionally reroute to main test cases page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/technique", http.StatusSeeOther)

			default:
				// otherwise just display main 'index' page
				if err := renderPage("index", nil, app, w, r); err != nil {
					app.Log.Error(fmt.Sprintf("Index HTTP GET %s", err.Error()))
					return
				}
			}

		} else {
			// if user not authenticated
			redirectToLoginPage(w, r, app)
		}
	})
}

// This is HTTP POST handler for datings.
func techniqueHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "":
		// insert new technique, when 'cmd' is empty...
		if t := parseTechniqueFormValues(r); t != nil {
			err = app.DataProv.InsertTechnique(t)
		} else {
			app.Log.Info(fmt.Sprintf("Creating new Technique '%s'", t.Name))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify technique: ID is empty")
		}
		if d := parseTechniqueFormValues(r); d != nil {
			d.Id = db.MongoStringToId(id)
			err = app.DataProv.UpdateTechnique(d)
			app.Log.Info(fmt.Sprintf("Updating Technique '%s'", d.Name))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete technique: ID is empty")
		}
		t := db.NewTechnique()
		t.Id = db.MongoStringToId(id)
		err = app.DataProv.DeleteTechnique(t)
		app.Log.Info(fmt.Sprintf("Removing Technique '%s'", t.Name))

	default:
		err = fmt.Errorf("Illegal POST request for technique")
	}
	return err
}

// Helper function that parses the '/dating' POST request values and creates a new instance of Dating.
func parseTechniqueFormValues(r *http.Request) *db.Technique {

	name := strings.TrimSpace(r.FormValue("name"))
	desc := strings.TrimSpace(r.FormValue("description"))
	created := strings.TrimSpace(r.FormValue("created"))

	t := db.NewTechnique()
	t.Name = name
	t.Description = desc
	t.Created = db.Timestamp(created)
	return t
}

// This is HTTP GET handler for datings.
func techniqueHTTPGetHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	t, err := app.DataProv.GetAllTechniques()
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting techniques from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Techniques []*db.Technique
		Num        int
		User       *db.User
	}{t, len(t), u}
	app.Log.Info("Displaying '/technique' page")
	return renderPage("techniques", web, app, w, r)
}
