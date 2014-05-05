/*
    web.go
 */
package main

import (
    "fmt"
    "bytes"
    "encoding/binary"
    "net/http"
    "path/filepath"
    "html/template"
    "math/rand"
    "crypto/sha512"
    "labix.org/v2/mgo/bson"
 //   "labix.org/v2/mgo"
    "github.com/gorilla/sessions"
    "github.com/gorilla/context"
    "bitbucket.org/miranr/artistic/utils"
)

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
    DB = ac.dbsess.DB("artistic")
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

    // register handler functions
    registerHandlers()

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

    log := ac.log
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
    fmt.Printf("DEBUG DB=%v\n", DB) // DEBUG

    switch r.Method {

    // when HTTP POST is received...
    case "POST":

        user := r.FormValue("username")
        pwd := r.FormValue("password")

        // create new user instance
        u := utils.CreateUser(user, pwd)

        // authenticate user
        status, err := authenticateUser(u, w, r)
        if !status && err != nil {
            if err = templates.ExecuteTemplate(w, "login", nil); err != nil {
            }
        }

        // if authenticated, redirect to index page; otherwise display login
        if status {
            http.Redirect(w, r, "index",  http.StatusFound)
        }

    // when HTTP GET is received, just display the default login template
    case "GET":
        if err := templates.ExecuteTemplate(w, "login", nil); err != nil {
        }
    }
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "/static/favicon.ico")
}

// authenticate the user with given username and password
func authenticateUser(u *utils.User,
                      w http.ResponseWriter, r *http.Request) (bool, error) {

    // create new session ID
    id := newSessId()
//    fmt.Printf("DEBUG: session ID=%q\n", id) // DEBUG

    // get information from DB
    querys := fmt.Sprintf("{ username: %q }", u.Username)
    err := ac.dbsess.DB("artistic").C("users").Find(querys).One(u)
    if err != nil {
    fmt.Printf("ERROR: user=%s\n", u.String()) // DEBUG
        return false, err
    //if cnt, err := ac.dbsess.DB("artistic").C("users").Count(); err != nil {
    //    fmt.Printf("DEBUG count=%d\n", cnt)
    //} else {
    //    fmt.Printf("DEBUG found user=%v\n", u) // DEBUG
    }
    fmt.Printf("DEBUG: user=%s\n", u.String()) // DEBUG

    // get current session data; will create new session with given random ID
    s, err := store.Get(r, id)
    if err != nil { return false, err }


    // save the session data
    s.Save(r, w)

    return true, nil
}

// check if user is already authenticated 
func userIsAuthenticated(r *http.Request) bool {

    s, err := store.Get(r, "session")
    if err != nil {
    }

    // get a session ID 
    sessid := s.Values["session-id"]
    fmt.Printf("Session ID: %v\n", sessid)

    //return false
    return true
}

// generate unique session ID; return it as string
func newSessId() string {

    // generate pseudo-random int64
    num := rand.Int63()

    // now hash the random int64 value with SHA512
    hash := sha512.Sum512(int64ToBytes(num))

    return fmt.Sprintf("%x", hash)
}

// Converts 64-bit integer value into byte buffer.
func int64ToBytes(i int64) []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, i)
    return buf.Bytes()
}

// retrieves all users from DoB
func getAllUsers() ([]utils.User, error) {
//func getAllUsers() ([]interface{}, error) {

    db := ac.dbsess.DB("artistic")

    // prepare the empty slice for users
    //u := make([]interface{}, 0)
    u := make([]utils.User, 0)

    // get all users from DB
    if err := db.C("users").Find(bson.D{}).All(&u); err != nil {
    //if err := DB.C("users").Find(bson.D{}).All(&u); err != nil {
        return nil, err
    }

    return u, nil
}

