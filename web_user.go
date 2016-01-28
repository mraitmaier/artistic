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

// Parse HTTP form values for user; this is used by /user and /profile
// The 'encryptPwd' flag denotes whether password should be encrypted (insert new user) or not (modify existing user)..
func parseUserFormValues(r *http.Request, encryptPwd bool) *db.User {

	// get POST form values and create a struct
	name := strings.TrimSpace(r.FormValue("username"))
	pwd := strings.TrimSpace(r.FormValue("password"))
	role := strings.TrimSpace(r.FormValue("urole"))
	full := strings.TrimSpace(r.FormValue("fullname"))
	email := strings.TrimSpace(r.FormValue("email"))
	phone := strings.TrimSpace(r.FormValue("phone"))
	disabled := strings.ToLower(strings.TrimSpace(r.FormValue("disabled")))
	mustchange := strings.ToLower(strings.TrimSpace(r.FormValue("mustchange")))
	created := strings.TrimSpace(r.FormValue("created"))

	// create a user and check passwords
	u := db.CreateUser(name, pwd, role, encryptPwd)
	// note: ID is handled in parent function
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

	var err error
	id := mux.Vars(r)["id"]

	if id == "" {
		return fmt.Errorf("Modify profile: ID is empty")
	}
	if s := parseUserFormValues(r, false); s != nil {
		s.Id = db.MongoStringToId(id)
		err = app.DataProv.UpdateUser(s)
		app.Log.Info(fmt.Sprintf("[%s] Updating Profile '%s'", s.Username, s.Fullname))
	}

	return err
}

// This is HTTP GET handler for user profile
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

// This is handler that handler the "/user" URL.
func userHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			var err error

			switch r.Method {

			case "GET":
				if err = userHTTPGetHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] User HTTP GET %s", user.Username, err.Error()))
				}

			case "POST":
				if err = userHTTPPostHandler(w, r, app, user); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] User HTTP POST %s", user.Username, err.Error()))
				}
				// unconditionally reroute to main user page
				http.Redirect(w, r, "/user", http.StatusFound)

			case "DELETE":
				msg := fmt.Sprintf("[%s] User HTTP DELETE request received. Redirecting to main 'user' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main user page
				// Use HTTP 303 (see other) to force GET to redirect as DELETE request is normally
				// followed by another DELETE
				http.Redirect(w, r, "/user", http.StatusSeeOther)

			case "PUT":
				msg := fmt.Sprintf("[%s] User HTTP PUT request received. Redirecting to main 'user' page.", user.Username)
				app.Log.Info(msg)
				// unconditionally reroute to main user page
				// Use HTTP 303 (see other) to force GET to redirect as PUT request is normally followed by
				// another PUT
				http.Redirect(w, r, "/user", http.StatusSeeOther)

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

// This is HTTP POST handler for users.
func userHTTPPostHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

	id := mux.Vars(r)["id"]
	cmd := mux.Vars(r)["cmd"]

	var err error
	switch strings.ToLower(cmd) {

	case "":
		// insert new user, when 'cmd' is empty...
		if s := parseUserFormValues(r, true); s != nil {
			err = app.DataProv.InsertUser(s)
		} else {
			app.Log.Info(fmt.Sprintf("[%s] Creating new User '%s (%s)'", u.Username, s.Fullname, s.Username))
		}

	case "put":
		if id == "" {
			return fmt.Errorf("Modify user: ID is empty")
		}
		if s := parseUserFormValues(r, false); s != nil {
			s.Id = db.MongoStringToId(id)
			err = app.DataProv.UpdateUser(s)
			app.Log.Info(fmt.Sprintf("[%s] Updating User '%s (%s)'", u.Username, s.Fullname, s.Username))
		}

	case "delete":
		if id == "" {
			return fmt.Errorf("Delete user: ID is empty")
		}
		s := db.NewUser()
		s.Id = db.MongoStringToId(id)
		err = app.DataProv.DeleteUser(s)
		app.Log.Info(fmt.Sprintf("[%s] Removing user '%s (%s)'", u.Username, s.Fullname, s.Username))

	default:
		err = fmt.Errorf("Illegal POST request for user")
	}
	return err
}

// This is HTTP GET handler for users
func userHTTPGetHandler(w http.ResponseWriter, r *http.Request, app *ArtisticApp, u *db.User) error {

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
	app.Log.Info(fmt.Sprintf("[%s] Displaying '/user' page", u.Username))
	return renderPage("users", web, app, w, r)
}

// Handle password change for users
func pwdHandler(app *ArtisticApp) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if user is autheticated, display the appropriate page
		if loggedin, user := userIsAuthenticated(app, r); loggedin {

			switch r.Method {

			case "POST":
				// get the previous page information from form values for redirection
				prev_page := strings.TrimSpace(r.FormValue("prev"))
				if uname, err := changeUserPassword(w, r, app); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] %s", user.Username, err.Error()))
				} else {
					app.Log.Info(fmt.Sprintf("[%s] Successfully changed password for user %q.", user.Username, uname))
				}
				// unconditionally reroute to previous page
				http.Redirect(w, r, fmt.Sprintf("/%s", prev_page), http.StatusFound)

			default:
				// otherwise just display main 'index' page
				if err := renderPage("index", nil, app, w, r); err != nil {
					app.Log.Error(fmt.Sprintf("[%s] Index HTTP GET %s", user.Username, err.Error()))
				}
			}

		} else {
			redirectToLoginPage(w, r, app) // if user not authenticated...
		}
	})
}

// Change password for existing user handler function.
func changeUserPassword(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) (string, error) {

	// get data to modify
	id := mux.Vars(r)["id"]
	if id == "" {
		return "", fmt.Errorf("Changing Password: user ID is empty")
	}

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

	// now do it... but refresh modified timestamp before
	u.Modified = db.NewTimestamp()
	if err = aa.DataProv.UpdateUser(u); err != nil {
		return "", err
	}
	return u.Username, err
}
