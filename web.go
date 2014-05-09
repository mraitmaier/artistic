/*
    web.go
 */
package main

import (
    "os"
    "runtime"
    "fmt"
    "bytes"
    "encoding/binary"
    "net/http"
    "path/filepath"
    "html/template"
    "labix.org/v2/mgo/bson"
 //   "labix.org/v2/mgo"
    "github.com/gorilla/sessions"
    "github.com/gorilla/context"
    "bitbucket.org/miranr/artistic/utils"
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

    // aux DB type instance
    //DB = aa.dbsess.DB("artistic")
)

// register web page handler functions
func registerHandlers() {

    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/index", indexHandler)
    http.HandleFunc("/users", usersHandler)
    http.HandleFunc("/error404", err404Handler)
    http.HandleFunc("/license", licenseHandler)
    http.HandleFunc("/favicon.ico", faviconHandler)
}

// initializes and starts web server
func webStart(wwwpath string) {

    aa.WebInfo = new(WebInfo)

    // register handler functions
    registerHandlers()

    // check dir for session files and create it if needed
    if !checkSessDir("") {
        aa.Log.Critical("Cannot create session folder; cannot continue...\n")
        return
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

/* this is currently not needed yet...
    if userIsAuthenticated(r) {
        if err := templates.ExecuteTemplate(w, "users", nil); err != nil {
        }
    } else {
        http.Redirect(w, r, "login",  http.StatusFound)
    }
*/
    // render the page
    if err := templates.ExecuteTemplate(w, "error404", nil); err != nil {
    }
}

// license page handler
func licenseHandler(w http.ResponseWriter, r *http.Request) {

/* this is currently not needed yet...
    if userIsAuthenticated(r) {
        if err := templates.ExecuteTemplate(w, "users", nil); err != nil {
        }
    } else {
        http.Redirect(w, r, "login",  http.StatusFound)
    }
*/
    // render the page
    if err := templates.ExecuteTemplate(w, "license", nil); err != nil {
    }
}


// Index (home) page handler
func indexHandler(w http.ResponseWriter, r *http.Request) {

    if userIsAuthenticated(r) {
        if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
        }
    } else {
        http.Redirect(w, r, "login",  http.StatusFound)
    }
}

// user admin page handler
func usersHandler(w http.ResponseWriter, r *http.Request) {

    log := aa.Log
/* this is currently not needed yet...
    if userIsAuthenticated(r) {
        if err := templates.ExecuteTemplate(w, "users", nil); err != nil {
        }
    } else {
        http.Redirect(w, r, "login.html",  http.StatusFound)
    }
*/

    // get all users from DB
    users, err := getAllUsers()
    if err != nil {
        log.Error(fmt.Sprintf("Problem getting all users: %s", err.Error()))
        http.Redirect(w, r, "error404", http.StatusFound)
        return
    }

    for _, val := range users {
        fmt.Printf("User: %s\n", val.String()) // DEBUG
    }

    // render the page
    if err := templates.ExecuteTemplate(w, "users", users); err != nil {
    }
}


// login page handler - we must authenticate user 
func loginHandler(w http.ResponseWriter, r *http.Request) {

//    s, err := store.Get(r, "session")
//    if err != nil {
//    }
    //fmt.Printf("DEBUG DB=%v\n", DB) // DEBUG

    log := aa.Log // get logger instance

    switch r.Method {

    // when HTTP POST is received...
    case "POST":

        user := r.FormValue("username")
        pwd := r.FormValue("password")

        // create new user instance
        u := utils.CreateUser(user, pwd)

        // authenticate user
        log.Info(fmt.Sprintf("Trying to authenticate user %q...\n", u.Username))
        status, err := authenticateUser(u, w, r)
        if !status && err != nil {
            if err = templates.ExecuteTemplate(w, "login", nil); err != nil {
                log.Error(err.Error())
            }
            log.Alert(fmt.Sprintf("User %q NOT authenticated.\n", u.Username))
        }

        // if authenticated, redirect to index page; otherwise display login
        if status {
            http.Redirect(w, r, "index",  http.StatusFound)
        }
        log.Info(fmt.Sprintf("User %q authenticated, OK.\n", u.Username))

    // when HTTP GET is received, just display the default login template
    case "GET":
        if err := templates.ExecuteTemplate(w, "login", nil); err != nil {
            log.Error("")
        }
    }
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "/static/favicon.ico")
}

// Converts 64-bit integer value into byte buffer.
func int64ToBytes(i int64) []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, i)
    return buf.Bytes()
}

// retrieves all users from DoB
func getAllUsers() ([]utils.User, error) {

    db := aa.DbSess.DB("artistic")

    // prepare the empty slice for users
    u := make([]utils.User, 0)

    // get all users from DB
    if err := db.C("users").Find(bson.D{}).All(&u); err != nil {
        return nil, err
    }

    return u, nil
}

