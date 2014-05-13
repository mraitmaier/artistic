/*
    web.go
 */
package main

import (
    "os"
    "runtime"
    "fmt"
    "errors"
//    "bytes"
//    "encoding/binary"
    "net/http"
    "path/filepath"
    "html/template"
//    "labix.org/v2/mgo/bson"
 //   "labix.org/v2/mgo"
    "github.com/gorilla/sessions"
    "github.com/gorilla/context"
    "bitbucket.org/miranr/artistic/utils"
    dbase "bitbucket.org/miranr/artistic/db"
)

type WebInfo struct {

    // a path to where session files are stored
    SessDir string
}

const (
    // a (quite) random string that is used as a key for sessions
    sessKey = `iufwnwieh3436SiKJSJo90e3jdiejdlje3+'0%$#!)dlkjja(!~§<sdfad$io*"`
)

var (
    // global var holding cached page templates
    templates *template.Template

    // global var holding the web session data
    store = sessions.NewCookieStore([]byte(sessKey))
)

// register web page handler functions
func registerHandlers() {

    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/logout", logoutHandler)
    http.HandleFunc("/index", indexHandler)
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/error404", err404Handler)
    http.HandleFunc("/license", licenseHandler)
    http.HandleFunc("/favicon.ico", faviconHandler)
}

// initializes and starts web server
func webStart(wwwpath string) error {

    aa.WebInfo = new(WebInfo)

    // register handler functions
    registerHandlers()

    // check dir for session files and create it if needed
    if !checkSessDir("") {
        return errors.New("Cannot create session folder; cannot continue.")
    }

    // handle static files
    path := filepath.Join(wwwpath, "static")
    http.Handle("/static/", http.StripPrefix("/static/",
            http.FileServer(http.Dir(path))))

    //web page templates, with defined additional functions
    funcs := template.FuncMap{
        "add": func(x,y int) int { return x + y } }
    t := filepath.Join(wwwpath, "templates", "*.tpl")
    templates = template.Must(template.New("").Funcs(funcs).ParseGlob(t))

    // finally, start web server, we're using HTTPS
//    http.ListenAndServe(":8088", context.ClearHandler(http.DefaultServeMux))
    http.ListenAndServeTLS(":8088", "./web/static/cert.pem",
                           "./web/static/key.pem",
                           context.ClearHandler(http.DefaultServeMux))
    return nil
}

func checkSessDir(path string) bool {

    basedir := path
    // if given base path is empty, default to temp folder
    if path == "" {
        switch runtime.GOOS {
            case "windows": basedir = os.Getenv("TEMP")
            default: basedir = "/tmp"
        }
        aa.WebInfo.SessDir = filepath.Join(basedir, "artistic", "sessions")
    }

    // if path does not exits, create it...
    if err := os.MkdirAll(aa.WebInfo.SessDir, 0755); err != nil  {
        fmt.Println("FATAL: Cannot create path. Cannot continue...")
        fmt.Println(err.Error()) // DEBUG
        return false
    }

    return true
}

// remove all session files (and the session folder itself!) when app is
// terminated.
func cleanSessDir() bool {
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
func err404Handler(w http.ResponseWriter, r *http.Request) {

    if loggedin, user := userIsAuthenticated(r); loggedin {

        // render the page
        if err := templates.ExecuteTemplate(w, "error404", user); err != nil {
            aa.Log.Error("Error rendering the '404' page.")
        }

    } else {
        http.Redirect(w, r, "login",  http.StatusFound)
    }
}

// user admin page handler
func logoutHandler(w http.ResponseWriter, r *http.Request) {

    if loggedin, user := userIsAuthenticated(r); loggedin {

        log := aa.Log

        // render the page
        if err := logout(w, r); err != nil {
           log.Error(err.Error())
        } else {
            log.Info(fmt.Sprintf("Logging out user %q.", user))
        }
    }
    http.Redirect(w, r, "login",  http.StatusFound)
}


// license page handler
func licenseHandler(w http.ResponseWriter, r *http.Request) {

    if loggedin, user := userIsAuthenticated(r); loggedin {

        log := aa.Log

        // render the page
        if err := templates.ExecuteTemplate(w, "license", user); err != nil {
            log.Error("Cannot render the 'license' page.")
        }

    } else {
        http.Redirect(w, r, "login",  http.StatusFound)
    }
}


// Index (home) page handler
func indexHandler(w http.ResponseWriter, r *http.Request) {

    if loggedin, user := userIsAuthenticated(r); loggedin {

        if err := templates.ExecuteTemplate(w, "index", user); err != nil {
        }

    } else {
        http.Redirect(w, r, "login",  http.StatusFound)
    }
}

// user admin page handler
func usersHandler(w http.ResponseWriter, r *http.Request) {

    if loggedin, user := userIsAuthenticated(r); loggedin {

        log := aa.Log

        // get all users from DB
        users, err := dbase.MongoGetAllUsers(aa.DbSess.DB("artistic"))
        if err != nil {
            log.Error(fmt.Sprintf("Problem getting all users: %s", err.Error()))
            http.Redirect(w, r, "error404", http.StatusFound)
            return
        }

// DEBUG
        for _, val := range users {
            fmt.Printf("User: %s\n", val.String()) // DEBUG
        }

        // create temp struct variable to be sent to page template
        web := new(struct{ User *utils.User
                           Users []utils.User })
        web.User = user
        web.Users = users

        // render the page
        if err := templates.ExecuteTemplate(w, "users", web); err != nil {
            log.Error("Cannot render the 'users' page.")
        }

    } else {
        http.Redirect(w, r, "login.html",  http.StatusFound)
    }
}


// login page handler - we must authenticate user 
func loginHandler(w http.ResponseWriter, r *http.Request) {

    log := aa.Log // get logger instance

    switch r.Method {

    // when HTTP POST is received...
    case "POST":

        user := r.FormValue("username")
        pwd := r.FormValue("password")

        // authenticate user
        log.Info(fmt.Sprintf("Trying to authenticate user %q...\n", user))
        status, err := authenticateUser(user, pwd, w, r)
        if !status || err != nil {
            if err = templates.ExecuteTemplate(w, "login", nil); err != nil {
                log.Error(err.Error())
            }
            log.Alert(fmt.Sprintf("User %q NOT authenticated.\n", user))
        }

        // if authenticated, redirect to index page; otherwise display login
        if status {
            http.Redirect(w, r, "index",  http.StatusFound)
        }
        log.Info(fmt.Sprintf("User %q authenticated, OK.\n", user))

    // when HTTP GET is received, just display the default login template
    case "GET":
        if err := templates.ExecuteTemplate(w, "login", nil); err != nil {
            log.Error("Cannot render the 'login' page.")
        }
    }
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "/static/favicon.ico")
}

/*
// Converts 64-bit integer value into byte buffer.
func int64ToBytes(i int64) []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, i)
    return buf.Bytes()
}
*/

