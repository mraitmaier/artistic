package main

//
//   web_user.go
//

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mraitmaier/artistic/db"
	"net/http"
	"strings"
)

// user admin page handler
func usersHandler(aa *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if loggedin, user := userIsAuthenticated(aa, r); loggedin {

			log := aa.Log

			// get all users from DB
			users, err := aa.DataProv.GetAllUsers()
			if err != nil {
				log.Error(fmt.Sprintf("[%s] Problem getting all users: %q", user.Username, err.Error()))
				http.Redirect(w, r, "/error", http.StatusFound)
				return
			}

			// create ad-hoc struct to be sent to page template
			var web = struct {
				User  *db.User
				Users []*db.User
			}{user, users}

			// render the page
			if err = renderPage("users", &web, aa, w, r); err != nil {
				log.Error(fmt.Sprintf("[%s] Rendering the 'users' page: %q", user.Username, err.Error()))
			}
			log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

		} else {
			redirectToLoginPage(w, r, aa)
		}
	}) // return handler closure
}

// a single user handler
func userHandler(aa *ArtisticApp) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if loggedin, user := userIsAuthenticated(aa, r); loggedin {

			log := aa.Log // get logger instance

			switch r.Method {

			case "GET":
				if err := getUserHandler(w, r, aa, user); err != nil {
					log.Error(fmt.Sprintf("[%s] User GET handler: %q.", user.Username, err.Error()))
					break // force break!
				}
				log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

			case "POST":
				if err := postUserHandler(w, r, aa, user); err != nil {
					log.Error(fmt.Sprintf("[%s] User POST handler: %q.", user.Username, err.Error()))
				}
				http.Redirect(w, r, "/users", http.StatusFound)

			case "DELETE":
				id := mux.Vars(r)["id"]
				cmd := mux.Vars(r)["cmd"]
				t := new(db.User)
				t.Id = db.MongoStringToId(id) // only valid ID needed to delete
				if err := aa.DataProv.DeleteUser(t); err != nil {
					log.Error(fmt.Sprintf("[%s] %s user id=%q, DB returned %q.", user.Username, cmd, id, err))
					return
				}
				log.Info(fmt.Sprintf("[%s] Successfully deleted user %q.", user, t.Username))
				http.Redirect(w, r, "/users", http.StatusFound)

			case "PUT":
				fmt.Printf("received PUT request. :)\n")
			}

		} else {
			redirectToLoginPage(w, r, aa)
		}
	}) // return handler closure
}

//  HTTP GET handler for "/user/<cmd>" URLs.
func getUserHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	log := aa.Log
	var err error
	s := new(db.User)

	switch cmd {

	case "view", "modify", "changepwd":

		// get a user from DB
		s, err = aa.DataProv.GetUser(id)
		if err != nil {
			return fmt.Errorf("%s user id=%q, DB returned %q", cmd, id, err)
		}
		log.Info(fmt.Sprintf("[%s] %s user=%q Success.", user.Username, strings.ToUpper(cmd), s.Username)) // OK log message

	case "insert": // do nothing here...

	case "delete":
		s.Id = db.MongoStringToId(id) // only valid ID needed to delete
		if err = aa.DataProv.DeleteUser(s); err != nil {
			return fmt.Errorf("%s user id=%q, DB returned %q", cmd, id, err)
		}
		log.Info(fmt.Sprintf("[%s] Successfully deleted user %q.", user.Username, s.Username))
		http.Redirect(w, r, "/users", http.StatusFound)
		return nil //  this is all about deleting items...

	default:
		return fmt.Errorf("unknown command %q", cmd)
	}

	// create ad-hoc struct to be sent to page template
	var web = struct {
		User        *db.User
		Cmd         string // "view", "modify", "insert" or "delete"...
		UserProfile *db.User
	}{user, cmd, s}

	return renderPage("user", &web, aa, w, r)
}

func postUserHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *db.User) error {

	// get data to modify
	cmd := mux.Vars(r)["cmd"]

	var err error
	var username string

	switch cmd {

	case "insert":
		if username, err = insertNewUser(w, r, aa); err != nil {
			return fmt.Errorf("Create (%q)", err.Error())
		}
		aa.Log.Info(fmt.Sprintf("[%s] Successfully inserted user %q.", user.Username, username))

	case "modify":
		if username, err = modifyExistingUser(w, r, aa); err != nil {
			return fmt.Errorf("Modify (%q)", err.Error())
		}
		aa.Log.Info(fmt.Sprintf("[%s] Successfully miodified user %q.", user.Username, username))

	case "changepwd":
		if username, err = changeUserPassword(w, r, aa); err != nil {
			return fmt.Errorf("Change password (%q)", err.Error())
		}
		aa.Log.Info(fmt.Sprintf("[%s] Successfully changed password for user %q.", user.Username, username))

	default:
		err = fmt.Errorf("Unknown command %q", cmd)
	}

	return err
}

// modify an existing user handler function.
func modifyExistingUser(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) (string, error) {

	// get data to modify
	id := mux.Vars(r)["id"]

	// get POST form values and create a struct
	name := strings.TrimSpace(r.FormValue("username"))
	pwd := strings.TrimSpace(r.FormValue("password"))
	role := strings.TrimSpace(r.FormValue("urole"))
	full := strings.TrimSpace(r.FormValue("fullname"))
	email := strings.TrimSpace(r.FormValue("email"))
	phone := strings.TrimSpace(r.FormValue("phone"))
	disabled := strings.ToLower(strings.TrimSpace(r.FormValue("disabled")))
	change := strings.ToLower(strings.TrimSpace(r.FormValue("change")))

	var err error

	// create a user and check passwords
	t := db.CreateUser(name, pwd, role, false)
	t.Id = db.MongoStringToId(id)
	t.Fullname = full
	t.Email = email
	t.Phone = phone
	if disabled == "no" {
		t.Disabled = false
	} else {
		t.Disabled = true
	}
	if change == "no" {
		t.MustChangePassword = false
	} else {
		t.MustChangePassword = true
	}

	// do it...
	if err = aa.DataProv.UpdateUser(t); err != nil {
		return "", err
	}
	return t.Username, err
}

// create new user handler function.
func insertNewUser(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) (string, error) {

	// get data to modify
	id := mux.Vars(r)["id"]

	// get POST form values and create a struct
	name := strings.TrimSpace(r.FormValue("username"))
	pwd := strings.TrimSpace(r.FormValue("password"))
	role := strings.TrimSpace(r.FormValue("urole"))
	full := strings.TrimSpace(r.FormValue("fullname"))
	email := strings.TrimSpace(r.FormValue("email"))
	phone := strings.TrimSpace(r.FormValue("phone"))
	disabled := strings.ToLower(strings.TrimSpace(r.FormValue("disabled")))
	change := strings.ToLower(strings.TrimSpace(r.FormValue("change")))

	var err error

	// create a user and check passwords
	t := db.CreateUser(name, pwd, role, true)
	t.Id = db.MongoStringToId(id)
	t.Fullname = full
	t.Email = email
	t.Phone = phone
	if disabled == "no" {
		t.Disabled = false
	} else {
		t.Disabled = true
	}
	if change == "no" {
		t.MustChangePassword = false
	} else {
		t.MustChangePassword = true
	}

	// do it...
	if err = aa.DataProv.InsertUser(t); err != nil {
		return "", err
	}
	return t.Username, err
}

// Change password for existing user handler function.
func changeUserPassword(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) (string, error) {

	// get data to modify
	id := mux.Vars(r)["id"]

	// get POST form values and create a struct
	old := strings.TrimSpace(r.FormValue("oldpassword"))
	pwd := strings.TrimSpace(r.FormValue("newpassword"))
	pwd2 := strings.TrimSpace(r.FormValue("newpassword2"))

	var err error

	u, err := aa.DataProv.GetUser(id)
	if err != nil {
		return "", fmt.Errorf("id=%q, DB returned %q", id, err)
	}

	// check password first and return error if they're not valid
	if err = u.ChangePassword(old, pwd, pwd2); err != nil {
		return "", err
	}

	/*
		if pwd != pwd2 {
			return "", fmt.Errorf("new passwords do not match")
		}
		if !u.ComparePassword(old) {
			return "", fmt.Errorf("invalid old password")
		}
	*/

	// now do it...
	if err = aa.DataProv.UpdateUser(u); err != nil {
		return "", err
	}

	return u.Username, err
}

// Parse HTTP form values for user; this is used by /user and /profile
func parseUserFormValues(r *http.Request) *db.User {

	// get data to modify
	id := mux.Vars(r)["id"]

	// get POST form values and create a struct
	name := strings.TrimSpace(r.FormValue("username"))
	pwd := strings.TrimSpace(r.FormValue("password"))
	role := strings.TrimSpace(r.FormValue("role"))
	full := strings.TrimSpace(r.FormValue("fullname"))
	email := strings.TrimSpace(r.FormValue("email"))
	phone := strings.TrimSpace(r.FormValue("phone"))
	disabled := strings.ToLower(strings.TrimSpace(r.FormValue("disabled")))
	mustchange := strings.ToLower(strings.TrimSpace(r.FormValue("mustchange")))
	created := strings.TrimSpace(r.FormValue("created"))

	// create a user and check passwords
	u := db.CreateUser(name, pwd, role, false)
	u.Id = db.MongoStringToId(id)
	u.Fullname = full
	u.Email = email
	u.Phone = phone
	u.Created = db.Timestamp(created)
	if disabled == "no" {
		u.Disabled = false
	} else {
		u.Disabled = true
	}
	if mustchange == "no" {
		u.MustChangePassword = false
	} else {
		u.MustChangePassword = true
	}
	return u
}

// This is handler that handler the "/profile" URL.
func profileHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = profileHTTPGetHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Profile HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = profileHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Profile HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main style page
				http.Redirect(w, r, "/profile", http.StatusFound)

			case "DELETE":
				msg := fmt.Sprintf("[%s] Profile HTTP DELETE request received. Redirecting to main 'profile' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main profile page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/profile", http.StatusSeeOther)

			case "PUT":
				msg := fmt.Sprintf("[%s] Profile HTTP PUT request received. Redirecting to main 'profile' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main profile page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/profile", http.StatusSeeOther)

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

// This is HTTP POST handler for profile
func profileHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "put":
		if id == "" {
			return fmt.Errorf("Modify profile: ID is empty")
		}
		if s := parseUserFormValues(r); s != nil {
			s.Id = db.MongoStringToId(id)
			err = app.DataProv.UpdateUser(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating Profile '%s'", s.Username, s.Fullname))
		}

	case "changepwd":
		if id == "" {
			return fmt.Errorf("Change password for profile: ID is empty")
		}
		if username, err := changeUserPassword(w, r, app); err != nil {
			return fmt.Errorf("Change password (%q)", err.Error())
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Successfully changed password for user %q.", u.Username, username))
		}

	default:
		err = fmt.Errorf("Illegal POST request for style")
	}
	return err
}

// This is HTTP GET handler for styles
func profileHTTPGetHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	s, err := app.DataProv.GetAllUsers()
	if err != nil {
		http.Redirect(w, r, "/err404", http.StatusFound)
		return fmt.Errorf("Problem getting users from DB: '%s'", err.Error())
	}
	// create ad-hoc struct to be sent to page template
	var web = struct {
		Users []*db.User
		Num   int
		User  *db.User
	}{s, len(s), u}
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/profile' page", u.Username))
	return renderPage("profile", web, app, w, r)
}
