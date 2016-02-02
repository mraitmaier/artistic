package main

//
//   web_artist.go
//

import (
	"fmt"
	"github.com/gorilla/mux"
	//	"github.com/mraitmaier/artistic/core"
	"github.com/mraitmaier/artistic/db"
	"net/http"
	"strings"
)

// This is handler that handler the "/artist" URL.
func paintingHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = paintingHTTPGetHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Painting HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = paintingHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Painting HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main artist page
				http.Redirect(w, r, "/painting", http.StatusFound)

			case "DELETE":
				msg := fmt.Sprintf("[%s] Painting HTTP DELETE request received. Redirecting to main 'painting' page.",
					user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main painting page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/painting", http.StatusSeeOther)

			case "PUT":
				msg := fmt.Sprintf("[%s] Painting HTTP PUT request received. Redirecting to main 'painting' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main painting page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/painting", http.StatusSeeOther)

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

// This is HTTP POST handler for paintings.
func paintingHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "":
		// insert new artist, when 'cmd' is empty...
		if s := parsePaintingFormValues(r); s != nil {
			err = app.DataProv.InsertPainting(s)
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Creating new painting '%s'", u.Username, s.Title))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify painting: ID is empty")
		}
		if s := parsePaintingFormValues(r); s != nil {
			s.Id = db.MongoStringToId(id)
			err = app.DataProv.UpdatePainting(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating painting '%s'", u.Username, s.Title))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete painting: ID is empty")
		}
		s := db.NewPainting()
		s.Id = db.MongoStringToId(id)
		err = app.DataProv.DeletePainting(s)
		app.Log.Info(fmt.Sprintf("[%s] Removing painting '%s'", u.Username, s.Title))

	default:
		err = fmt.Errorf("Illegal POST request for painting")
	}
	return err
}

// Helper function that parses the '/artist' POST request values and creates a new instance of Artist
func parsePaintingFormValues(r *http.Request) *db.Painting {

	// get POST form values and create a struct
	p := db.NewPainting()
	p.Title = strings.TrimSpace(r.FormValue("title"))
	p.Artist = strings.TrimSpace(r.FormValue("artist"))
	p.Style = strings.TrimSpace(r.FormValue("artstyle"))
	p.Technique = strings.TrimSpace(r.FormValue("technique"))
	p.Size = strings.TrimSpace(r.FormValue("size"))
	p.Dating = r.FormValue("dating")
	p.TimeOfCreation = strings.TrimSpace(r.FormValue("timecreat"))
	p.Motive = strings.TrimSpace(r.FormValue("motive"))
	p.Signature = strings.TrimSpace(r.FormValue("signature"))
	p.Place = strings.TrimSpace(r.FormValue("place"))
	p.Location = strings.TrimSpace(r.FormValue("location"))
	p.Provenance = strings.TrimSpace(r.FormValue("provenance"))
	p.Condition = strings.TrimSpace(r.FormValue("condition"))
	p.ConditionDescription = strings.TrimSpace(r.FormValue("conddescription"))
	p.Description = strings.TrimSpace(r.FormValue("description"))
	//p.Exhibitions = strings.TrimSace(r.FormValue("exhibitions"))
	//p.Sources = strings.TrimSpace(r.FormValue("sources"))
	//p.Notes = strings.TrimSpace(r.FormValue("notes"))
	//p.Picture = strings.TrimSpace(r.FormValue("picture"))
	p.Created = db.Timestamp(strings.TrimSpace(r.FormValue("created")))
	return p
}

// This is HTTP GET handler for artists
func paintingHTTPGetHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	a, err := app.DataProv.GetAllPaintings()
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting paintings from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Paintings []*db.Painting
		Num       int
		User      *db.User
	}{a, len(a), u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/painting' page", u.Username))
	return renderPage("paintings", web, app, w, r)
}
