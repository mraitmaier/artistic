package main

//
//   web_artwork.go
//

import (
	"fmt"
	"github.com/gorilla/mux"
	//	"github.com/mraitmaier/artworkic/core"
	"github.com/mraitmaier/artistic/db"
	"net/http"
	"strings"
)

// This is handler that handler the "/artwork" URL.
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
				// unconditionally reroute to main artwork page
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
		// insert new artwork, when 'cmd' is empty...
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

// Helper function that parses the '/artwork' POST request values and creates a new instance of Artist
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

// This is HTTP GET handler for paintings
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

// This is handler that handler the "/sculpture" URL.
func sculptureHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = sculptureHTTPGetHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Sculpture HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = sculptureHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Sculpture HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main artwork page
				http.Redirect(w, r, "/sculpture", http.StatusFound)

			case "DELETE":
				msg := fmt.Sprintf("[%s] Sculpture HTTP DELETE request received. Redirecting to main 'sculpture' page.",
					user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main painting page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/sculpture", http.StatusSeeOther)

			case "PUT":
				msg := fmt.Sprintf("[%s] Sculpture HTTP PUT request received. Redirecting to main 'sculpture' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main sculpture page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/sculpture", http.StatusSeeOther)

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

// This is HTTP POST handler for sculptures.
func sculptureHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "":
		// insert new artwork, when 'cmd' is empty...
		if s := parseSculptureFormValues(r); s != nil {
			err = app.DataProv.InsertSculpture(s)
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Creating new sculpture '%s'", u.Username, s.Title))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify sculpture: ID is empty")
		}
		if s := parseSculptureFormValues(r); s != nil {
			s.Id = db.MongoStringToId(id)
			err = app.DataProv.UpdateSculpture(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating sculpture '%s'", u.Username, s.Title))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete sculpture: ID is empty")
		}
		s := db.NewSculpture()
		s.Id = db.MongoStringToId(id)
		err = app.DataProv.DeleteSculpture(s)
		app.Log.Info(fmt.Sprintf("[%s] Removing sculpture '%s'", u.Username, s.Title))

	default:
		err = fmt.Errorf("Illegal POST request for sculpture")
	}
	return err
}

// Helper function that parses the '/sculpture' POST request values and creates a new instance of Sculpture
func parseSculptureFormValues(r *http.Request) *db.Sculpture {

	// get POST form values and create a struct
	p := db.NewSculpture()
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

// This is HTTP GET handler for sculptures
func sculptureHTTPGetHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	s, err := app.DataProv.GetAllSculptures()
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting culptures from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Sculptures []*db.Sculpture
		Num        int
		User       *db.User
	}{s, len(s), u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/sculpture' page", u.Username))
	return renderPage("sculptures", web, app, w, r)
}

// This is handler that handler the "/print" URL.
func printHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = printHTTPGetHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Print HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = printHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Print HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main print page
				http.Redirect(w, r, "/print", http.StatusFound)

			case "DELETE":
				msg := fmt.Sprintf("[%s] Print HTTP DELETE request received. Redirecting to main 'print' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main print page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/print", http.StatusSeeOther)

			case "PUT":
				msg := fmt.Sprintf("[%s] Print HTTP PUT request received. Redirecting to main 'print' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main print page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/print", http.StatusSeeOther)

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

// This is HTTP POST handler for graphic prints.
func printHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "":
		// insert new print, when 'cmd' is empty...
		if s := parsePrintFormValues(r); s != nil {
			err = app.DataProv.InsertPrint(s)
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Creating new graphic print '%s'", u.Username, s.Title))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify graphic print: ID is empty")
		}
		if s := parsePrintFormValues(r); s != nil {
			s.Id = db.MongoStringToId(id)
			err = app.DataProv.UpdatePrint(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating graphic print '%s'", u.Username, s.Title))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete graphic print: ID is empty")
		}
		s := db.NewPrint()
		s.Id = db.MongoStringToId(id)
		err = app.DataProv.DeletePrint(s)
		app.Log.Info(fmt.Sprintf("[%s] Removing graphic print '%s'", u.Username, s.Title))

	default:
		err = fmt.Errorf("Illegal POST request for graphic print")
	}
	return err
}

// Helper function that parses the '/print' POST request values and creates a new instance of Print
func parsePrintFormValues(r *http.Request) *db.Print {

	// get POST form values and create a struct
	p := db.NewPrint()
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

// This is HTTP GET handler for graphic prints
func printHTTPGetHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	p, err := app.DataProv.GetAllPrints()
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting graphic prints from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Prints []*db.Print
		Num    int
		User   *db.User
	}{p, len(p), u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/print' page", u.Username))
	return renderPage("prints", web, app, w, r)
}
