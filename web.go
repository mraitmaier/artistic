/*
 */
package main

import (
   // "fmt"
    "net/http"
    "path/filepath"
    "html/template"
)

// global var holding cached page templates
var templates *template.Template

func registerHandlers() {

    http.HandleFunc("/", indexHandler)
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
    http.ListenAndServe(":8088", nil)
}

/* Index (home) page handler */
func indexHandler(w http.ResponseWriter, r *http.Request) {
    if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
    }
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
 //   http.ServeFile(
}

