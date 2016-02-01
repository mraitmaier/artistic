package main

//
//   web_artist.go
//

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mraitmaier/artistic/core"
	"github.com/mraitmaier/artistic/db"
	"net/http"
	"strings"
)

// This is handler that handler the "/artist" URL.
func artistHandler(app *ArtisticApp, t db.ArtistType) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = artistHTTPGetHandler(w, r, app, user, t); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Artist HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = artistHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Artist HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main artist page
				http.Redirect(w, r, "/artist", http.StatusFound)

			case "DELETE":
				msg := fmt.Sprintf("[%s] Artist HTTP DELETE request received. Redirecting to main 'artist' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main artist page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/artist", http.StatusSeeOther)

			case "PUT":
				msg := fmt.Sprintf("[%s] Artist HTTP PUT request received. Redirecting to main 'artist' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main artist page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/artist", http.StatusSeeOther)

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

// This is HTTP POST handler for artists.
func artistHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "":
		// insert new artist, when 'cmd' is empty...
		if s := parseArtistFormValues(r); s != nil {
			err = app.DataProv.InsertArtist(s)
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Creating new Artist '%s'", u.Username, s.Name))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify artist: ID is empty")
		}
		if s := parseArtistFormValues(r); s != nil {
			s.Id = db.MongoStringToId(id)
			err = app.DataProv.UpdateArtist(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating Artist '%s'", u.Username, s.Name))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete artist: ID is empty")
		}
		s := db.NewArtist()
		s.Id = db.MongoStringToId(id)
		err = app.DataProv.DeleteArtist(s)
		app.Log.Info(fmt.Sprintf("[%s] Removing artist '%s'", u.Username, s.Name))

	default:
		err = fmt.Errorf("Illegal POST request for artist")
	}
	return err
}

// Helper function that parses the '/artist' POST request values and creates a new instance of Artist
func parseArtistFormValues(r *http.Request) *db.Artist {

	// get POST form values and create a struct
	first := strings.TrimSpace(r.FormValue("first"))
	middle := strings.TrimSpace(r.FormValue("middle"))
	last := strings.TrimSpace(r.FormValue("last"))
	rfirst := strings.TrimSpace(r.FormValue("realfirst"))
	rmiddle := strings.TrimSpace(r.FormValue("realmiddle"))
	rlast := strings.TrimSpace(r.FormValue("reallast"))
	born := strings.TrimSpace(r.FormValue("born"))
	died := strings.TrimSpace(r.FormValue("died"))
	nation := strings.TrimSpace(r.FormValue("nationality"))
	painter := strings.TrimSpace(r.FormValue("painter"))
	sculptor := strings.TrimSpace(r.FormValue("sculptor"))
	printmaker := strings.TrimSpace(r.FormValue("printmaker"))
	architect := strings.TrimSpace(r.FormValue("architect"))
	ceramicist := strings.TrimSpace(r.FormValue("ceramicist"))
	bio := strings.TrimSpace(r.FormValue("biography"))
    created := strings.TrimSpace(r.FormValue("created"))

	// create an Artist instance
    a := db.NewArtist()
	a.Name = core.CreateName(first, middle, last)
    a.RealName = core.CreateName(rfirst, rmiddle, rlast)
	a.Born = born
	a.Died = died
	a.Nationality = nation
	a.Biography = bio
    a.Created = db.Timestamp(created)
	if painter == "yes" {
		a.IsPainter = true
	}
	if sculptor == "yes" {
		a.IsSculptor = true
	}
	if ceramicist == "yes" {
		a.IsCeramicist = true
	}
	if printmaker == "yes" {
		a.IsPrintmaker = true
	}
	if architect == "yes" {
		a.IsArchitect = true
	}
	return a
}

// This is HTTP GET handler for artists
func artistHTTPGetHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User, t db.ArtistType) error {

	a, err := app.DataProv.GetAllArtists(t)
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting artists from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Artists []*db.Artist
		Num    int
		Type    db.ArtistType
		User   *db.User
	}{a, len(a), t, u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/artist' page", u.Username))
    return renderPage("artists", web, app, w, r) 
}

