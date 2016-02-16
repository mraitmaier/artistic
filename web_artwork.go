package main

//
//   web_artwork.go
//

import (
	"fmt"
	"github.com/gorilla/mux"
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
				if err = paintingHTTPGetHandler("", w, r, app, user); err != nil {
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
			s.ID = db.MongoStringToId(id)
			err = app.DataProv.UpdatePainting(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating painting '%s'", u.Username, s.Title))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete painting: ID is empty")
		}
		s := db.NewPainting()
		s.ID = db.MongoStringToId(id)
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
func paintingHTTPGetHandler(qry string, w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	a, err := app.DataProv.GetPaintings(qry)
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting paintings from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Paintings []*db.Painting
		Num       int
		Ptype     string
		User      *db.User
	}{a, len(a), "painting", u}
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
				if err = sculptureHTTPGetHandler("", w, r, app, user); err != nil {
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
			s.ID = db.MongoStringToId(id)
			err = app.DataProv.UpdateSculpture(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating sculpture '%s'", u.Username, s.Title))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete sculpture: ID is empty")
		}
		s := db.NewSculpture()
		s.ID = db.MongoStringToId(id)
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
func sculptureHTTPGetHandler(qry string, w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	s, err := app.DataProv.GetSculptures(qry)
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting culptures from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Sculptures []*db.Sculpture
		Num        int
		Ptype      string
		User       *db.User
	}{s, len(s), "sculpture", u}
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
				if err = printHTTPGetHandler("", w, r, app, user); err != nil {
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
			s.ID = db.MongoStringToId(id)
			err = app.DataProv.UpdatePrint(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating graphic print '%s'", u.Username, s.Title))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete graphic print: ID is empty")
		}
		s := db.NewPrint()
		s.ID = db.MongoStringToId(id)
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
func printHTTPGetHandler(qry string, w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	p, err := app.DataProv.GetPrints(qry)
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting graphic prints from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Prints []*db.Print
		Num    int
		Ptype  string
		User   *db.User
	}{p, len(p), "print", u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/print' page", u.Username))
	return renderPage("prints", web, app, w, r)
}

// This is handler that handler the "/building" URL.
func buildingHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = buildingHTTPGetHandler("", w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Building HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = buildingHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Building HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main building page
				http.Redirect(w, r, "/building", http.StatusFound)

			case "DELETE":
				msg := fmt.Sprintf("[%s] Building HTTP DELETE request received. Redirecting to main 'building' page.",
					user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main building page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/building", http.StatusSeeOther)

			case "PUT":
				msg := fmt.Sprintf("[%s] Building HTTP PUT request received. Redirecting to main 'building' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main building page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/building", http.StatusSeeOther)

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

// This is HTTP POST handler for buildings.
func buildingHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "":
		// insert new building, when 'cmd' is empty...
		if s := parseBuildingFormValues(r); s != nil {
			err = app.DataProv.InsertBuilding(s)
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Creating new building '%s'", u.Username, s.Title))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify building: ID is empty")
		}
		if s := parseBuildingFormValues(r); s != nil {
			s.ID = db.MongoStringToId(id)
			err = app.DataProv.UpdateBuilding(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating building '%s'", u.Username, s.Title))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete building: ID is empty")
		}
		s := db.NewBuilding()
		s.ID = db.MongoStringToId(id)
		err = app.DataProv.DeleteBuilding(s)
		app.Log.Info(fmt.Sprintf("[%s] Removing building '%s'", u.Username, s.Title))

	default:
		err = fmt.Errorf("Illegal POST request for building")
	}
	return err
}

// Helper function that parses the '/building' POST request values and creates a new instance of Building
func parseBuildingFormValues(r *http.Request) *db.Building {

	// get POST form values and create a struct
	p := db.NewBuilding()
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

// This is HTTP GET handler for buildings
func buildingHTTPGetHandler(qry string, w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	p, err := app.DataProv.GetBuildings(qry)
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting buildings from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Buildings []*db.Building
		Num       int
		Ptype     string
		User      *db.User
	}{p, len(p), "building", u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/building' page", u.Username))
	return renderPage("buildings", web, app, w, r)
}

// This is handler that handler the "/book" URL.
func bookHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = bookHTTPGetHandler("", w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Book HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = bookHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Book HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main book page
				http.Redirect(w, r, "/book", http.StatusFound)

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

// This is HTTP POST handler for books.
func bookHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "":
		// insert new book , when 'cmd' is empty...
		if b := parseBookFormValues(r); b != nil {
			err = app.DataProv.InsertBook(b)
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Creating new book '%s'", u.Username, b.Title))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify book: ID is empty")
		}
		if b := parseBookFormValues(r); b != nil {
			b.ID = db.MongoStringToId(id)
			err = app.DataProv.UpdateBook(b)
			app.Log.Info(fmt.Sprintf("[%s] Updating book '%s'", u.Username, b.Title))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete book: ID is empty")
		}
		b := db.NewBook()
		b.ID = db.MongoStringToId(id)
		err = app.DataProv.DeleteBook(b)
		app.Log.Info(fmt.Sprintf("[%s] Removing book '%s'", u.Username, b.Title))

	default:
		err = fmt.Errorf("Illegal POST request for book")
	}
	return err
}

// Helper function that parses the '/book' POST request values and creates a new instance of Book
func parseBookFormValues(r *http.Request) *db.Book {

	// get POST form values and create a struct
	b := db.NewBook()
	b.Title = strings.TrimSpace(r.FormValue("title"))
	b.Authors = strings.TrimSpace(r.FormValue("authors"))
	b.Edition = strings.TrimSpace(r.FormValue("edition"))
	b.Publisher = strings.TrimSpace(r.FormValue("publisher"))
	b.Year = strings.TrimSpace(r.FormValue("year"))
	b.Location = strings.TrimSpace(r.FormValue("location"))
	b.ISBN = strings.TrimSpace(r.FormValue("isbn"))
	b.Keywords = strings.TrimSpace(r.FormValue("keywords"))
	b.Created = db.Timestamp(strings.TrimSpace(r.FormValue("created")))
	//b.Notes = strings.TrimSpace(r.FormValue("notes"))
	//b.Picture = strings.TrimSpace(r.FormValue("picture"))
	return b
}

// This is HTTP GET handler for books
func bookHTTPGetHandler(qry string, w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	p, err := app.DataProv.GetBooks(qry)
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting books from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Books []*db.Book
		Num   int
		Ptype string
		User  *db.User
	}{p, len(p), "book", u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/book' page", u.Username))
	return renderPage("books", web, app, w, r)
}

// This is handler that handler the "/article" URL.
func articleHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = articleHTTPGetHandler("", w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Article HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = articleHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Article HTTP POST %s", user.Username, err.Error()))
				}
				http.Redirect(w, r, "/article", http.StatusFound)

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

// This is HTTP POST handler for articles.
func articleHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "":
		// insert new article, when 'cmd' is empty...
		if a := parseArticleFormValues(r); a != nil {
			err = app.DataProv.InsertArticle(a)
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Creating new article '%s'", u.Username, a.Title))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify article: ID is empty")
		}
		if a := parseArticleFormValues(r); a != nil {
			a.ID = db.MongoStringToId(id)
			err = app.DataProv.UpdateArticle(a)
			app.Log.Info(fmt.Sprintf("[%s] Updating article '%s'", u.Username, a.Title))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete article: ID is empty")
		}
		a := db.NewArticle()
		a.ID = db.MongoStringToId(id)
		err = app.DataProv.DeleteArticle(a)
		app.Log.Info(fmt.Sprintf("[%s] Removing article '%s'", u.Username, a.Title))

	default:
		err = fmt.Errorf("Illegal POST request for article")
	}
	return err
}

// Helper function that parses the '/article' POST request values and creates a new instance of Article
func parseArticleFormValues(r *http.Request) *db.Article {

	// get POST form values and create a struct
	a := db.NewArticle()
	a.Title = strings.TrimSpace(r.FormValue("title"))
	a.Authors = strings.TrimSpace(r.FormValue("authors"))
	a.Publication = strings.TrimSpace(r.FormValue("publication"))
	a.Volume = strings.TrimSpace(r.FormValue("volume"))
	a.Issue = strings.TrimSpace(r.FormValue("issue"))
	a.Year = strings.TrimSpace(r.FormValue("year"))
	a.Publisher = strings.TrimSpace(r.FormValue("publisher"))
	a.Location = strings.TrimSpace(r.FormValue("location"))
	a.ISSN = strings.TrimSpace(r.FormValue("issn"))
	a.Keywords = strings.TrimSpace(r.FormValue("keywords"))
	a.Created = db.Timestamp(strings.TrimSpace(r.FormValue("created")))
	//a.Notes = strings.TrimSpace(r.FormValue("notes"))
	a.Link = strings.TrimSpace(r.FormValue("link"))
	return a
}

// This is HTTP GET handler for articles
func articleHTTPGetHandler(qry string, w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	p, err := app.DataProv.GetArticles(qry)
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting articles from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Articles []*db.Article
		Num      int
		Ptype    string
		User     *db.User
	}{p, len(p), "article", u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/article' page", u.Username))
	return renderPage("articles", web, app, w, r)
}
