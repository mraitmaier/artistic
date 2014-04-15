/*
 */
package main

import (
    "fmt"
    "net/http"
    "path/filepath"
    "html/template"
    "github.com/gorilla/sessions"
    "github.com/gorilla/context"
)

// global var holding cached page templates
var templates *template.Template
// global var holdiung the web session data
var store = sessions.NewCookieStore(
                        []byte("Something a bit more secret than default"))

func registerHandlers() {

    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/index", indexHandler)
    http.HandleFunc("/favicon.ico", faviconHandler)
}

/* initializes and starts web server */
func webStart(ac *ArtisticCtrl, wwwpath string) {

    // register handler functions
    registerHandlers()

    // handle static files
    path := filepath.Join(wwwpath, "static")
    http.Handle("/static/", http.StripPrefix("/static/",
            http.FileServer(http.Dir(path))))

    //web page templates
    t := filepath.Join(wwwpath, "templates", "*.tpl")
    templates = template.Must(template.ParseGlob(t))

    // finally, start web server
//    http.ListenAndServe(":8088", context.ClearHandler(http.DefaultServeMux))
    http.ListenAndServeTLS(":8088", "./web/static/cert.pem",
                           "./web/static/key.pem",
                           context.ClearHandler(http.DefaultServeMux))
}

/* Index (home) page handler */
func indexHandler(w http.ResponseWriter, r *http.Request) {
    if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
    }
}

/* login page handler - we must authenticate user */
func loginHandler(w http.ResponseWriter, r *http.Request) {

    s, err := store.Get(r, "session")
    if err != nil {
    }
    /* get a session ID */
    sessid := s.Values["session-id"]
    fmt.Printf("Session ID: %v\n", sessid)

    if err := templates.ExecuteTemplate(w, "login", nil); err != nil {
    }
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "/static/favicon.ico")
}

/* check if user is already authenticated */
func UserIsAuthenticated(username, pwd string, sess *sessions.Session) bool {
    return false
}

