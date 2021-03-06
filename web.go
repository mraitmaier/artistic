package main

//
//   web.go
//

import (
	//"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/mraitmaier/artistic/core"
	"github.com/mraitmaier/artistic/db"
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

	// (error, info etc.) message to be displayed on page
	//Msg WebMessage
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
	r.Handle("/pwd/{id}", pwdHandler(aa))
	r.Handle("/search", searchHandler(aa))
	r.Handle("/user", userHandler(aa))
	r.Handle("/user/{id}/{cmd}", userHandler(aa))
	//	r.Handle("/log", logHandler(aa))
	r.Handle("/technique", techniqueHandler(aa))
	r.Handle("/technique/{id}/{cmd}", techniqueHandler(aa))
	r.Handle("/style", styleHandler(aa))
	r.Handle("/style/{id}/{cmd}", styleHandler(aa))
	r.Handle("/dating", datingHandler(aa))
	r.Handle("/dating/{id}/{cmd}", datingHandler(aa))
	r.Handle("/error404", err404Handler(aa))
	r.Handle("/error", errorHandler(aa))
	r.Handle("/license", licenseHandler(aa))
	r.Handle("/profile", profileHandler(aa))
	r.Handle("/profile/{id}", profileHandler(aa))
	r.Handle("/artist", artistHandler(aa, db.ArtistTypeArtist))
	r.Handle("/painter", artistHandler(aa, db.ArtistTypePainter))
	r.Handle("/sculptor", artistHandler(aa, db.ArtistTypeSculptor))
	r.Handle("/printmaker", artistHandler(aa, db.ArtistTypePrintmaker))
	r.Handle("/architect", artistHandler(aa, db.ArtistTypeArchitect))
	r.Handle("/artist/{id}/{cmd}", artistHandler(aa, db.ArtistTypeArtist))
	r.Handle("/painting", paintingHandler(aa))
	r.Handle("/painting/{id}/{cmd}", paintingHandler(aa))
	r.Handle("/sculpture", sculptureHandler(aa))
	r.Handle("/sculpture/{id}/{cmd}", sculptureHandler(aa))
	r.Handle("/print", printHandler(aa))
	r.Handle("/print/{id}/{cmd}", printHandler(aa))
	r.Handle("/building", buildingHandler(aa))
	r.Handle("/building/{id}/{cmd}", buildingHandler(aa))
	r.Handle("/book", bookHandler(aa))
	r.Handle("/book/{id}/{cmd}", bookHandler(aa))
	r.Handle("/article", articleHandler(aa))
	r.Handle("/article/{id}/{cmd}", articleHandler(aa))
	r.HandleFunc("/favicon.ico", faviconHandler)
	r.NotFoundHandler = err404Handler(aa)

	// Call the default URL router...
	http.Handle("/", r)
}

/*
// ErrorMessage is the structure for displaying messages on web page
type ErrorMessage struct {

	// Code is internal error code
	Code int `json:"code"`
	// Type denotes the message type: warning, danger, success, info
	Type string `json:"type"`
	// Text is the actual text of the message
	Message string `json:"message"`
	// Status is HTTP Error Code: 200, 404 etc.
	Status int `json:"status"`
}

// A handy string represenation of the ErrorMessage instance.
func (m *ErrorMessage) String() string {
	return fmt.Sprintf("%d (%s): %s", m.Code, m.Type, m.Message)
}

// MarshalJSON marshals the error message into JSON-encoded text.
func (m *ErrorMessage) MarshalJSON() (string, error) {
	b, err := json.Marshal(m)
	return string(b[:]), err
}

// UnmarshalJSON unmarshals the JSON-encoded text into error message.
func (m *ErrorMessage) UnmarshalJSON(j string) error {
	err := json.Unmarshal([]byte(j), m)
	return err
}
*/

// The webStart function initializes and starts the web server.
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
		"add":                 func(x, y int) int { return x + y },
		"length":              func(list []string) int { return len(list) },
		"allowedroles":        func() []string { return db.AllRoles },
		"get_artist_type":     func(t db.ArtistType) string { return t.String() },
		"get_datings":         func() []string { return getDatingNames(aa) },
		"get_styles":          func() []string { return getStyleNames(aa) },
		"get_techniques":      func() []string { return getTechniqueNames(aa) },
		"get_technique_types": func() []core.TechniqueType { return core.GetTechniqueTypes() },
		"totitle":             func(s string) string { return strings.Title(s) },
		"toupper":             func(s string) string { return strings.ToUpper(s) },
		"tolower":             func(s string) string { return strings.ToLower(s) }}
	t := filepath.Join(wwwpath, "templates", "*.tpl")
	aa.WebInfo.templates = template.Must(template.New("").Funcs(funcs).ParseGlob(t))

	// finally, start web server, we're using HTTPS
	// http.ListenAndServe(":8088", context.ClearHandler(http.DefaultServeMux))
	http.ListenAndServeTLS(":8088", "./web/static/cert.pem", "./web/static/key.pem", nil)
	//		"./web/static/key.pem", context.ClearHandler(http.DefaultServeMux))
	return nil
}

// Helper function that checks path and creates it if needed.
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

// Remove all session files (and the session folder itself!) when app is terminated.
func cleanSessDir(aa *ArtisticApp) bool {

	status := false

	if aa.WebInfo.sessDir != "" {
		if err := os.RemoveAll(aa.WebInfo.sessDir); err != nil {
			aa.Log.Error(fmt.Sprintf("Remove session directory %q", err.Error()))
			return status
		}
		status = true
	}
	return status
}

// Aux function redirecting to login page.
func redirectToLoginPage(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) {
	aa.Log.Warning("User not authenticated")
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
		http.Redirect(w, r, "/err404", http.StatusFound)
	}
	return err
}

// HTTP error 404 page handler
func err404Handler(aa *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if loggedin, user := userIsAuthenticated(aa, r); loggedin {

			web := struct {
				Ptype string
				User  *db.User
			}{"", user}
			if err := renderPage("error404", web, aa, w, r); err != nil {
				aa.Log.Error(fmt.Sprintf("[%s] Cannot render the 'error404' template: %q.", user.Username, err.Error()))
				return
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
			web := struct {
				Ptype string
				User  *db.User
			}{"error", user}
			if err := renderPage("error", web, aa, w, r); err != nil {
				aa.Log.Error(fmt.Sprintf("[%s] Cannot render the 'error404' template: %q.", user.Username, err.Error()))
				return
			}
			aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

		} else {
			redirectToLoginPage(w, r, aa)
		}

	}) // return handler closure
}

// logout handler
func logoutHandler(aa *ArtisticApp) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if loggedin, user := userIsAuthenticated(aa, r); loggedin {

			if err := logout(aa, w, r); err != nil {
				aa.Log.Error(fmt.Sprintf("[%s] Logging out %q", user.Username, err.Error()))
			} else {
				aa.Log.Info(fmt.Sprintf("[%s] User logged out.", user.Username))
			}
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}) // return handler closure
}

// license page handler
func licenseHandler(aa *ArtisticApp) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if loggedin, user := userIsAuthenticated(aa, r); loggedin {

			web := struct {
				Ptype string
				User  *db.User
			}{"license", user}

			if err := renderPage("license", web, aa, w, r); err != nil {
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

			web := struct {
				Ptype string
				User  *db.User
			}{"index", user}

			if err := renderPage("index", web, aa, w, r); err != nil {
				aa.Log.Error(fmt.Sprintf("[%s] Cannot render the 'index' template: %q.", user.Username, err.Error()))
				return
			}
			aa.Log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))
		} else {
			redirectToLoginPage(w, r, aa)
		}
	}) // return handler closure
}

// helper function that displays the /index (or /) page, this is done quite a lot...
func displayIndexPage(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	// create ad-hoc struct to be sent to page template
	var web = struct {
		Ptype string
		User  *db.User
	}{"", u}
	return renderPage("index", web, app, w, r)
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
				log.Alert(fmt.Sprintf("User %q NOT authenticated.", user))
				err = aa.WebInfo.templates.ExecuteTemplate(w, "login", nil)
				aa.Log.Info(fmt.Sprintf("Displaying the %q page.", r.RequestURI))
				if err != nil {
					log.Error(err.Error())
				}
			}

			// if authenticated, redirect to index page; otherwise display login
			if status {
				log.Info(fmt.Sprintf("[%s] User authenticated, OK.", user))
				http.Redirect(w, r, "/index", http.StatusFound)
			}

		// when HTTP GET is received, just display the default login template
		case "GET":
			err := aa.WebInfo.templates.ExecuteTemplate(w, "login", nil)
			if err != nil {
				log.Error(fmt.Sprintf("Error rendering the 'login' template: %q.", err.Error()))
			}
			aa.Log.Info(fmt.Sprintf("Displaying the %q page.", r.RequestURI))
		}
	}) // return handler closure
}

/*
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
				Ptype    string
				Contents []string // a list of log messages
				PerPage  int      // how many log messages per page is displayed...
			}{user, "log", contents, 25}

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
*/

// favincon handler
func faviconHandler(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, favicon) }

// This is handler that handler the "/dating" URL.
func datingHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = datingHTTPGetHandler("", w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Dating HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = datingHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Dating HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main dating page
				http.Redirect(w, r, "/dating", http.StatusFound)

			case "DELETE":
				msg := fmt.Sprintf("[%s] Dating HTTP DELETE request received. Redirecting to main 'dating' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main dating page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/dating", http.StatusSeeOther)

			case "PUT":
				msg := fmt.Sprintf("[%s] Dating HTTP PUT request received. Redirecting to main 'dating' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main dating page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/dating", http.StatusSeeOther)

			default:
				// otherwise just display main 'index' page
				if err := renderPage("index", nil, app, w, r); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Index HTTP GET %s", user.Username, err.Error()))
					return
				}
			}

		} else {
			redirectToLoginPage(w, r, app) // if user not authenticated
		}
	})
}

// This is HTTP POST handler for datings.
func datingHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "put":
		if id == "" {
			return fmt.Errorf("Modify dating: ID is empty")
		}
		if d := parseDatingFormValues(r); d != nil {
			d.ID = db.MongoStringToId(id)
			if err = updateCachedDating(d, app); err != nil {
				return err
			}
			err = app.DataProv.UpdateDating(d)
			app.Log.Info(fmt.Sprintf("[%s] Updating Dating '%s'", u.Username, d.Dating))
		}

	default:
		err = fmt.Errorf("Illegal POST request for dating")
	}
	return err
}

// Helper function that updates the cached list of datings.
func updateCachedDating(d *db.Dating, app *ArtisticApp) error {

	var err error
	if app.Cached.Datings != nil {

		for _, val := range app.Cached.Datings {
			if val.ID == d.ID {
				val.Dating = d.Dating
				val.Description = d.Description
				val.Modified = db.NewTimestamp()
			}
		}
	} else {
		err = fmt.Errorf("Datings cache empty?")
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
func datingHTTPGetHandler(qry string, w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	var d []*db.Dating
	var err error

	if app.Cached.Datings == nil {
		d, err = app.DataProv.GetDatings(qry)
		if err != nil {
			http.Redirect(w, r, "/err404", http.StatusFound)
			return fmt.Errorf("Problem getting datings from DB: '%s'", err.Error())
		}
	} else {
		d = app.Cached.Datings
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Datings []*db.Dating
		Num     int
		Ptype   string
		User    *db.User
	}{d, len(d), "dating", u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/dating' page", u.Username))
	return renderPage("datings", web, app, w, r)
}

// This is handler that handler the "/technique" URL.
func techniqueHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = techniqueHTTPGetHandler("", w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Technique HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = techniqueHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Technique HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main technique page
				http.Redirect(w, r, "/technique", http.StatusFound)

			case "DELETE":
				msg := fmt.Sprintf("[%s] Technique HTTP DELETE request received. Redirecting to main 'technique' page.",
					user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main technique page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/technique", http.StatusSeeOther)

			case "PUT":
				msg := fmt.Sprintf("[%s] Technique HTTP PUT request received. Redirecting to main 'technique' page.",
					user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main technique page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/technique", http.StatusSeeOther)

			default:
				// otherwise just display main 'index' page
				if err := renderPage("index", nil, app, w, r); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Index HTTP GET %s", user.Username, err.Error()))
					return
				}
			}

		} else {
			// if user not authenticated
			redirectToLoginPage(w, r, app)
		}
	})
}

// This is HTTP POST handler for techniques
func techniqueHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]
	var err error

	switch strings.ToLower(cmd) {

	case "":
		// insert new technique, when 'cmd' is empty...
		if t := parseTechniqueFormValues(r); t != nil {
			if app.Cached.Techniques != nil {
				app.Cached.Techniques = append(app.Cached.Techniques, t)
			}
			err = app.DataProv.InsertTechnique(t)
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Creating new Technique '%s'", u.Username, t.Name))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify technique: ID is empty")
		}
		if t := parseTechniqueFormValues(r); t != nil {
			t.ID = db.MongoStringToId(id)
			if err = updateCachedTechniques(t, app); err != nil {
				return err
			}
			err = app.DataProv.UpdateTechnique(t)
			app.Log.Info(fmt.Sprintf("[%s] Updating Technique '%s'", u.Username, t.Name))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete technique: ID is empty")
		}
		t := db.NewTechnique()
		t.ID = db.MongoStringToId(id)
		if err = deleteCachedTechnique(t, app); err != nil {
			return err
		}
		err = app.DataProv.DeleteTechnique(t)
		app.Log.Info(fmt.Sprintf("[%s] Removing Technique '%s'", u.Username, t.Name))

	default:
		err = fmt.Errorf("Illegal POST request for technique")
	}
	return err
}

// Helper function that updates the cached list of techniques.
func updateCachedTechniques(t *db.Technique, app *ArtisticApp) error {

	var err error
	if app.Cached.Techniques != nil {

		for _, val := range app.Cached.Techniques {
			if val.ID == t.ID {
				val.Name = t.Name
				val.Description = t.Description
				val.Type = t.Type
				val.Modified = db.NewTimestamp()
			}
		}
	} else {
		err = fmt.Errorf("Updating Techniques: techniques cache empty?")
	}
	return err
}

// Helper function that deletes the cached technique.
func deleteCachedTechnique(t *db.Technique, app *ArtisticApp) error {

	var err error
	if app.Cached.Techniques != nil {
		ix := 0
		for _, val := range app.Cached.Techniques {
			if val.ID == t.ID {
				newtech := make([]*db.Technique, len(app.Cached.Techniques)-1)
				copy(newtech, app.Cached.Techniques[:ix])
				copy(newtech, app.Cached.Techniques[ix+1:])
				app.Cached.Techniques = newtech
				break
			}
			ix++
		}
	} else {
		err = fmt.Errorf("Removing Technique: techniques cache empty?")
	}
	return err
}

// Helper function that parses the '/technique' POST request values and creates a new instance of Technique
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

// This is HTTP GET handler for techniques
func techniqueHTTPGetHandler(qry string, w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	var t []*db.Technique
	var err error

	if app.Cached.Techniques == nil {
		t, err = app.DataProv.GetTechniques(qry)
		if err != nil {
			http.Redirect(w, r, "/err404", http.StatusFound)
			return fmt.Errorf("Problem getting techniques from DB: '%s'", err.Error())
		}
	} else {
		t = app.Cached.Techniques
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Techniques []*db.Technique
		Num        int
		Ptype      string
		User       *db.User
	}{t, len(t), "technique", u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/technique' page", u.Username))
	return renderPage("techniques", web, app, w, r)
}

// This is handler that handler the "/style" URL.
func styleHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = styleHTTPGetHandler("", w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Style HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = styleHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Style HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main style page
				http.Redirect(w, r, "/style", http.StatusFound)

			case "DELETE":
				msg := fmt.Sprintf("[%s] Style HTTP DELETE request received. Redirecting to main 'style' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main style page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/style", http.StatusSeeOther)

			case "PUT":
				msg := fmt.Sprintf("[%s] Style HTTP PUT request received. Redirecting to main 'style' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main style page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/style", http.StatusSeeOther)

			default:
				// otherwise just display main 'index' page
				if err := renderPage("index", nil, app, w, r); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Index HTTP GET %s", user.Username, err.Error()))
					return
				}
			}

		} else {
			// if user not authenticated
			redirectToLoginPage(w, r, app)
		}
	})
}

// This is HTTP POST handler for styles.
func styleHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "":
		// insert new style, when 'cmd' is empty...
		if s := parseStyleFormValues(r); s != nil {
			if app.Cached.Styles != nil {
				app.Cached.Styles = append(app.Cached.Styles, s)
			}
			err = app.DataProv.InsertStyle(s)
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Creating new Style '%s'", u.Username, s.Name))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify style: ID is empty")
		}
		if s := parseStyleFormValues(r); s != nil {
			s.ID = db.MongoStringToId(id)
			if err = updateCachedStyle(s, app); err != nil {
				return err
			}
			err = app.DataProv.UpdateStyle(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating Style '%s'", u.Username, s.Name))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete style: ID is empty")
		}
		s := db.NewStyle()
		s.ID = db.MongoStringToId(id)
		if err = deleteCachedStyle(s, app); err != nil {
			return err
		}
		err = app.DataProv.DeleteStyle(s)
		app.Log.Info(fmt.Sprintf("[%s] Removing style '%s'", u.Username, s.Name))

	default:
		err = fmt.Errorf("Illegal POST request for style")
	}
	return err
}

// Helper function that updates the cached list of styles.
func updateCachedStyle(s *db.Style, app *ArtisticApp) error {

	var err error
	if app.Cached.Styles != nil {

		for _, val := range app.Cached.Styles {
			if val.ID == s.ID {
				val.Name = s.Name
				val.Description = s.Description
				val.Modified = db.NewTimestamp()
			}
		}
	} else {
		err = fmt.Errorf("Styles cache empty?")
	}
	return err
}

// Helper function that deletes the cached styles.
func deleteCachedStyle(s *db.Style, app *ArtisticApp) error {

	var err error
	if app.Cached.Styles != nil {
		ix := 0
		for _, val := range app.Cached.Styles {
			if val.ID == s.ID {
				newst := make([]*db.Style, len(app.Cached.Styles)-1)
				copy(newst, app.Cached.Styles[:ix])
				copy(newst, app.Cached.Styles[ix+1:])
				app.Cached.Styles = newst
				break
			}
			ix++
		}
	} else {
		err = fmt.Errorf("Removing Technique: techniques cache empty?")
	}
	return err
}

// Helper function that parses the '/style' POST request values and creates a new instance of Style
func parseStyleFormValues(r *http.Request) *db.Style {

	name := strings.TrimSpace(r.FormValue("name"))
	desc := strings.TrimSpace(r.FormValue("description"))
	created := strings.TrimSpace(r.FormValue("created"))

	s := db.NewStyle()
	s.Name = name
	s.Description = desc
	s.Created = db.Timestamp(created)
	return s
}

// This is HTTP GET handler for styles
func styleHTTPGetHandler(qry string, w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	var s []*db.Style
	var err error

	if app.Cached.Styles == nil {
		s, err = app.DataProv.GetStyles(qry) // empty query  retrieves all records
		if err != nil {
			http.Redirect(w, r, "/err404", http.StatusFound)
			return fmt.Errorf("Problem getting styles from DB: '%s'", err.Error())
		}
	} else {
		s = app.Cached.Styles
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Styles []*db.Style
		Num    int
		Ptype  string
		User   *db.User
	}{s, len(s), "style", u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/style' page", u.Username))
	return renderPage("styles", web, app, w, r)
}

// This is handler that handler the "/search" URL. It accepts only POST requests.
func searchHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "POST":
				if err = searchHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Search HTTP POST %s", user.Username, err.Error()))
				}

			default:
				// otherwise just display main 'index' page
				if err := renderPage("index", nil, app, w, r); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Index HTTP GET %s", user.Username, err.Error()))
					return
				}
			}

		} else {
			// if user not authenticated
			redirectToLoginPage(w, r, app)
		}
	})
}

// This is HTTP POST handler for searches.
func searchHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	qry := strings.TrimSpace(r.FormValue("search-string"))
	ptype := strings.TrimSpace(r.FormValue("search-type"))

	// if type is empty, we cannot do anything with it, just redirect to index page
	if ptype == "" {
		app.Log.Error(fmt.Sprintf("[%s] HTTP POST Search: unknown data type, redirecting to /index page", u.Username))
		return displayIndexPage(w, r, app, u)
	}
	return resolveURL(ptype, qry, w, r, app, u)
}

// The helper function that resolves the proper search context and calls the appropriate GET handler.
func resolveURL(ptype, qry string, w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	var err error

	switch strings.ToLower(ptype) {
	case "style":
		err = styleHTTPGetHandler(qry, w, r, app, u)
	case "technique":
		err = techniqueHTTPGetHandler(qry, w, r, app, u)
	case "dating":
		err = datingHTTPGetHandler(qry, w, r, app, u)
	case "artist":
		err = artistHTTPGetHandler(qry, w, r, app, u, db.ArtistTypeArtist)
	case "painter":
		err = artistHTTPGetHandler(qry, w, r, app, u, db.ArtistTypePainter)
	case "sculptor":
		err = artistHTTPGetHandler(qry, w, r, app, u, db.ArtistTypeSculptor)
	case "printmaker":
		err = artistHTTPGetHandler(qry, w, r, app, u, db.ArtistTypePrintmaker)
	case "architect":
		err = artistHTTPGetHandler(qry, w, r, app, u, db.ArtistTypeArchitect)
	case "painting":
		err = paintingHTTPGetHandler(qry, w, r, app, u)
	case "building":
		err = buildingHTTPGetHandler(qry, w, r, app, u)
	case "sculpture":
		err = sculptureHTTPGetHandler(qry, w, r, app, u)
	case "print":
		err = printHTTPGetHandler(qry, w, r, app, u)
	case "user":
		err = userHTTPGetHandler(qry, w, r, app, u)
	case "book":
		err = bookHTTPGetHandler(qry, w, r, app, u)
	case "article":
		err = articleHTTPGetHandler(qry, w, r, app, u)
	default:
		// just render the /index page
		return displayIndexPage(w, r, app, u)
	}
	return err
}

// aux function that gets the list of dating names
func getDatingNames(app *ArtisticApp) []string {

	d, err := app.DataProv.GetDatingNames()
	if err != nil {
		app.Log.Error(fmt.Sprintf("Error getting a list of datings: %q", err.Error()))
	}
	return d
}

// aux function that gets the list of style names
func getStyleNames(app *ArtisticApp) []string {

	s, err := app.DataProv.GetStyleNames()
	if err != nil {
		app.Log.Error(fmt.Sprintf("Error getting a list of styles: %q", err.Error()))
	}
	return s
}

// aux function that gets the list of technique names
func getTechniqueNames(app *ArtisticApp) []string {

	t, err := app.DataProv.GetTechniqueNames()
	if err != nil {
		app.Log.Error(fmt.Sprintf("Error getting a list of techniques: %q", err.Error()))
	}
	return t
}
