//
//   web_artist.go
//
package main

import (
	"fmt"
    "strings"
	"net/http"
	"bitbucket.org/miranr/artistic/core"
	"bitbucket.org/miranr/artistic/utils"
	"bitbucket.org/miranr/artistic/db"
	"github.com/gorilla/mux"
)

// The artists page handler: one handler factory function to rule them all
func artistsHandler(aa *ArtisticApp, t core.ArtistType) http.Handler {

    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		log := aa.Log
        var err error

        // get all painters from DB
        artists, err := aa.DataProv.GetAllArtists(t)
        if err != nil {
            log.Error(fmt.Sprintf("[%s] Getting artists from DB: %q", user.Username, err.Error()))
            http.Redirect(w, r, "/error", http.StatusFound)
            return
        }

	    // create ad-hoc struct to be sent to page template
        var web = struct {
		    User  *utils.User
            Type core.ArtistType
            Artists []core.Artist
        } { user, t, artists }

        if err = renderPage("artists", &web, aa, w, r); err != nil {
			log.Error(fmt.Sprintf("[%s] Rendering the 'artists' page (%s)", user.Username, err.Error()))
		}
        log.Info(fmt.Sprintf("[%s] Displaying the %s page.", user.Username, r.RequestURI))

	} else {
        redirectToLoginPage(w, r, aa)
	}
    }) // return handler closure
}

//
func artistHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

	    log := aa.Log // get logger instance

        switch r.Method {

        case "GET":
            if err := getArtistHandler(w, r, aa, user); err != nil {
                log.Error(fmt.Sprintf("[%s] Artist GET handler: %q.", user.Username, err.Error()))
			    //http.Redirect(w, r, "/artists", http.StatusFound)
                break // force break
            }
            log.Info(fmt.Sprintf("[%s] Displaying the %q page", user.Username, r.RequestURI))

        case "POST":
            if err := postArtistHandler(w, r, aa, user); err != nil {
                log.Error(fmt.Sprintf("[%s] Artist POST handler: %q.", user.Username, err.Error()))
            }
			http.Redirect(w, r, "/artists", http.StatusFound)

        case "DELETE":
            log.Warning("received DELETE request. :)")

        case "PUT":
            log.Warning("received PUT request. :)")
        }

	} else {
        redirectToLoginPage(w, r, aa)
	}
    }) // return handler closure
}

//  HTTP GET handler for "/artist/<cmd>" URLs.
func getArtistHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *utils.User) error {
//                      at core.ArtistType, user *utils.User) error {

//    id := mux.Vars(r)["id"]
    cmd := mux.Vars(r)["cmd"]

//    log := aa.Log
    //var err error

    // create new artist instance
    a := core.CreateArtist()

    switch cmd {

    case "view", "modify", "changepwd":

	    // get a user from DB
    /*
	    s, err = aa.DataProv.GetArtist(id)
	    if err != nil {
		    err = fmt.Errorf("%s user id=%q, DB returned %q.", cmd, id, err)
            return err
	    }
    */

    case "insert": // do nothing here...

    case "delete":
    /*
        s.Id = db.MongoStringToId(id) // only valid ID needed to delete 
        if err = aa.DataProv.DeleteUser(s); err != nil {
            return fmt.Errorf("%s user id=%q, DB returned %q.", cmd, id, err)
        }
        log.Info(fmt.Sprintf("Successfully deleted artist %q.", s.Name))
	    http.Redirect(w, r, "/artists", http.StatusFound)
        return nil //  this is all about deleting items...
    */

    default:
        return fmt.Errorf("Unknown command %q.", cmd)
    }

	// create ad-hoc struct to be sent to page template
    var web = struct {
		User  *utils.User
        Cmd   string   // "view", "modify", "insert" or "delete"...
		Artist *core.Artist
    }{ user, cmd, a }

    return renderPage("artist", &web, aa, w, r)
}

// HTTP POST handler for "/painter/<cmd>" URLs.
func postArtistHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *utils.User) error {

    // get data to modify 
    cmd := mux.Vars(r)["cmd"]

    var err error = nil
    var name string

    switch cmd {

    case "insert":
        if name, err = insertNewArtist(w, r, aa) ; err != nil {
            return fmt.Errorf("INSERT Failed (%s)", err.Error())
        }
        aa.Log.Info(fmt.Sprintf("[%s] Artist %q INSERT Success.", user.Username, name))

    case "modify":
        if name, err = modifyExistingArtist(w, r, aa); err != nil {
            err = fmt.Errorf("MODIFY Failed (%s)", err.Error())
        }
        aa.Log.Info(fmt.Sprintf("[%s] Artist %q MODIFY Success.", user.Username, name))

    default:
        err = fmt.Errorf("Unknown command %q", cmd)
    }

    return err
}

// Modify an existing user handler function.
// Return the name of the modified artist and error code. Return empty string for name when error occurs. 
func modifyExistingArtist(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) (string, error) {

    // get data to modify 
	id  := mux.Vars(r)["id"]

    // get POST form values and create a struct
    a, err := parseFormValues(r)
    if err != nil {
        return "", err
    }
    a.Id = db.MongoStringToId(id)

    // do it...
    if err = aa.DataProv.UpdateArtist(a); err != nil {
        return "", err
    }

    return a.Name.String(), err
}

// Create new user handler function.
// Return the name of the modified artist and error code. Return empty string for name when error occurs. 
func insertNewArtist(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) (string, error) {

    // get data to modify 
	id  := mux.Vars(r)["id"]

    a, err := parseFormValues(r)
    if err != nil {
        return "", err
    }
    a.Id = db.MongoStringToId(id)

    // do it...
    if err = aa.DataProv.InsertArtist(a); err != nil {
       return "", err
    }

    //aa.Log.Info(fmt.Sprintf("Successfully created new artist %q.", a.Name))
    return a.Name.String(), err
}

// aux function that parses the HTTP POST form data and creates an Artist instance
func parseFormValues(r *http.Request) (a *core.Artist, err error) {

    // get POST form values and create a struct
    first := strings.TrimSpace(r.FormValue("first"))
    middle := strings.TrimSpace(r.FormValue("middle"))
    last := strings.TrimSpace(r.FormValue("last"))
    born := strings.TrimSpace(r.FormValue("born"))
    died := strings.TrimSpace(r.FormValue("died"))
    nation := strings.TrimSpace(r.FormValue("nationality"))

    //if r.FormValue != nil {
    painter := strings.TrimSpace(r.FormValue("painter"))
    //}
    sculptor := strings.TrimSpace(r.FormValue("sculprtor"))
    printmaker := strings.TrimSpace(r.FormValue("printmaker"))
    architect := strings.TrimSpace(r.FormValue("architect"))
    ceramicist := strings.TrimSpace(r.FormValue("ceramicist"))

    // create an Artist instance
    a = core.CreateArtist()
    a.Name = core.CreateName(first, middle, last)
    a.Born = born
    a.Died = died
    a.Nationality = nation

    if painter == "yes" { a.IsPainter = true }
    if sculptor == "yes" { a.IsSculptor = true }
    if ceramicist == "yes" { a.IsCeramicist = true }
    if printmaker == "yes" { a.IsPrintmaker = true }
    if architect == "yes" { a.IsArchitect = true }

    return
}
