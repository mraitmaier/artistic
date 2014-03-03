/*
    web.go -
 */
package main

import (
    "fmt"
    "net/http"
    "path/filepath"
    "html/template"
)

type webCtrl struct {

    // cached templates
    templ *template.Template
}

func webStart(ac *ArtisticCtrl, wwwpath string) *webCtrl {

    web := new(webCtrl)

    // register handler functions
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/index", indexHandler)

    // handle static files
    http.Handle(wwwpath,
                http.StripPrefix(wwwpath, http.FileServer(http.Dir("web"))))
//    http.Handle(wwwpath, http.FileServer(http.Dir("web")))

    //web page templates
    t := filepath.Join(wwwpath, "templates", "*.tpl")
    web.templ = template.Must(template.ParseGlob(t))

    // finally, start web server
    http.ListenAndServe(":8088", nil)

    return web
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Artistic Index Web Page")
}


