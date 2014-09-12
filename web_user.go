//
//   web_user.go
//
package main

import (
	"fmt"
    "strings"
	"net/http"
	"bitbucket.org/miranr/artistic/utils"
	"bitbucket.org/miranr/artistic/db"
	"github.com/gorilla/mux"
)

// user admin page handler
func usersHandler(aa *ArtisticApp) http.Handler {

    return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

		log := aa.Log

		// get all users from DB
		users, err := aa.DataProv.GetAllUsers()
		if err != nil {
			log.Error(fmt.Sprintf("[%s] Problem getting all users: %q", user, err.Error()))
			http.Redirect(w, r, "/error", http.StatusFound)
			return
		}

		// create ad-hoc struct to be sent to page template
		var web = struct {
			User  *utils.User
			Users []utils.User
		} { user, users }

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
    return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

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
                log.Error(fmt.Sprintf("[%s] User POST handler: %q.", user, err.Error()))
            }
			http.Redirect(w, r, "/users", http.StatusFound)

        case "DELETE":
            id := mux.Vars(r)["id"]
            cmd := mux.Vars(r)["cmd"]
            t := new(utils.User)
            t.Id = db.MongoStringToId(id) // only valid ID needed to delete 
            if err := aa.DataProv.DeleteUser(t); err != nil {
                log.Error(fmt.Sprintf("[%s] %s user id=%q, DB returned %q.", user, cmd, id, err))
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
func getUserHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *utils.User) error {

    id := mux.Vars(r)["id"]
    cmd := mux.Vars(r)["cmd"]

    log := aa.Log
    var err error
    s := new(utils.User)

    switch cmd {

    case "view", "modify", "changepwd":

	    // get a user from DB
	    s, err = aa.DataProv.GetUser(id)
	    if err != nil {
		    err = fmt.Errorf("%s user id=%q, DB returned %q.", cmd, id, err)
            return err
	    }
        log.Info(fmt.Sprintf("[%s] %s user=%q Success.", user, strings.ToUpper(cmd), s.Username)) // OK log message

    case "insert": // do nothing here...

    case "delete":
        s.Id = db.MongoStringToId(id) // only valid ID needed to delete 
        if err = aa.DataProv.DeleteUser(s); err != nil {
            return fmt.Errorf("%s user id=%q, DB returned %q.", cmd, id, err)
        }
        log.Info(fmt.Sprintf("[%s] Successfully deleted user %q.", user, s.Username))
	    http.Redirect(w, r, "/users", http.StatusFound)
        return nil //  this is all about deleting items...

    default:
        return fmt.Errorf("unknown command %q", cmd)
    }

	// create ad-hoc struct to be sent to page template
    var web = struct {
		User  *utils.User
        Cmd   string   // "view", "modify", "insert" or "delete"...
		UserProfile *utils.User
    }{ user, cmd, s }

    return renderPage("user", &web, aa, w, r)
}

func postUserHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *utils.User) error {

    // get data to modify 
    cmd := mux.Vars(r)["cmd"]

    var err error = nil
    var username string

    switch cmd {

    case "insert":
        if username, err = insertNewUser(w, r, aa) ; err != nil {
            return fmt.Errorf("Create (%q)", err.Error())
        }
        aa.Log.Info(fmt.Sprintf("[%s] Successfully inserted user %q.", user, username))

    case "modify":
        if username, err = modifyExistingUser(w, r, aa); err != nil {
            return fmt.Errorf("Modify (%q)", err.Error())
        }
        aa.Log.Info(fmt.Sprintf("[%s] Successfully miodified user %q.", user, username))

    case "changepwd":
        if username, err = changeUserPassword(w, r, aa); err != nil {
            return fmt.Errorf("Change password (%q)", err.Error())
        }
        aa.Log.Info(fmt.Sprintf("[%s] Successfully changed password for user %q.", user, username))

    default:
        err = fmt.Errorf("Unknown command %q", cmd)
    }

    return err
}

// modify an existing user handler function.
func modifyExistingUser(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) (string, error) {

    // get data to modify 
	id  := mux.Vars(r)["id"]

    // get POST form values and create a struct
    name  := strings.TrimSpace(r.FormValue("username"))
    pwd   := strings.TrimSpace(r.FormValue("password"))
    role  := strings.TrimSpace(r.FormValue("role"))
    full  := strings.TrimSpace(r.FormValue("fullname"))
    email := strings.TrimSpace(r.FormValue("email"))

    var err error = nil

    // create a user and check passwords
    t := utils.CreateUser(name, pwd)
    if err = t.SetRole(role); err != nil {
        return "", fmt.Errorf("invalid role")
    }
    t.Id = db.MongoStringToId(id)
    t.Name = full
    t.Role = role
    t.Email = email

    // do it...
    if err = aa.DataProv.UpdateUser(t); err != nil {
        return "", err
    }
    //aa.Log.Info(fmt.Sprintf("[%s] Successfully updated user %q.", user, t.Username))
    return t.Username, err
}

// create new user handler function.
func insertNewUser(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) (string, error) {

    // get data to modify 
	id  := mux.Vars(r)["id"]

    // get POST form values and create a struct
    name  := strings.TrimSpace(r.FormValue("username"))
    pwd   := strings.TrimSpace(r.FormValue("password"))
    role  := strings.TrimSpace(r.FormValue("role"))
    full  := strings.TrimSpace(r.FormValue("fullname"))
    email := strings.TrimSpace(r.FormValue("email"))

    var err error = nil

    // create a user and check passwords
    t := utils.CreateUser(name, pwd)

    if err = t.SetRole(role); err != nil {
        return "", fmt.Errorf("invalid role")
    }
    t.Id = db.MongoStringToId(id)
    t.Name = full
    t.Role = role
    t.Email = email

    // do it...
    if err = aa.DataProv.InsertUser(t); err != nil {
        return "", err
    }
    return t.Username, err
}

// Change password for existing user handler function.
func changeUserPassword(w http.ResponseWriter, r *http.Request, aa *ArtisticApp) (string, error) {

    var err error = nil

    // get data to modify 
	id  := mux.Vars(r)["id"]

    // get POST form values and create a struct
    old  := strings.TrimSpace(r.FormValue("oldpassword"))
	pwd  := strings.TrimSpace(r.FormValue("newpassword"))
	pwd2 := strings.TrimSpace(r.FormValue("newpassword2"))

	u, err := aa.DataProv.GetUser(id)
	if err != nil {
		return "", fmt.Errorf("id=%q, DB returned %q", id, err)
	}

    // check password first and return error if they're not valid
    if pwd != pwd2 {
        return "", fmt.Errorf("new passwords do not match")
    }
    if !u.ComparePassword(old) {
        return "", fmt.Errorf("invalid old password")
    }

    // now do it...
    u.SetPassword(pwd)
    if err = aa.DataProv.UpdateUser(u); err != nil {
        return "", err
    }

    return u.Username, err
}

// a user profile handler
func profileHandler(aa *ArtisticApp) http.Handler {
    return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

	if loggedin, user := userIsAuthenticated(aa, r); loggedin {

	    log := aa.Log // get logger instance

        switch r.Method {

        case "GET":
            if err := getProfileHandler(w, r, aa, user); err != nil {
                log.Error(fmt.Sprintf("[%s] Profile GET handler: %q.", user, err.Error()))
			    //http.Redirect(w, r, "/users", http.StatusFound)
                break // force break!
            }
            log.Info(fmt.Sprintf("[%s] Displaying the %q page.", user.Username, r.RequestURI))

        case "POST":
            if err := postProfileHandler(w, r, aa, user); err != nil {
                log.Error(fmt.Sprintf("[%s] Profile POST handler: %q.", user, err.Error()))
            }
			http.Redirect(w, r, "/users", http.StatusFound)
        }

	} else {
        redirectToLoginPage(w, r, aa)
	}
    }) // return handler closure
}

//  HTTP GET handler for "/userprofile/<cmd>" URLs.
func getProfileHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *utils.User) error {

    cmd := mux.Vars(r)["cmd"]

    //log := aa.Log
    //var err error

	// create ad-hoc struct to be sent to page template
    var web = struct {
		User  *utils.User
        Cmd   string   // "view", "modify", "changepwd"...
    }{ user, cmd }

	return renderPage("userprofile", &web, aa, w, r)
}

func postProfileHandler(w http.ResponseWriter, r *http.Request, aa *ArtisticApp, user *utils.User) error {

    // get data to modify 
    cmd := mux.Vars(r)["cmd"]

    var err error = nil
    //var username string

    switch cmd {

    case "modify":
        //if username, err = modifyExistingUser(w, r, aa); err != nil {
        if _, err = modifyExistingUser(w, r, aa); err != nil {
            err = fmt.Errorf("[%s] MODIFY Profile: %q", user. Username, err.Error())
        }

    case "changepwd":
        //if username, err = changeUserPassword(w, r, aa); err != nil {
        if _, err = changeUserPassword(w, r, aa); err != nil {
            err = fmt.Errorf("[%s] CHANGEPWD Profile: %q", user. Username, err.Error())
        }

    default:
        err = fmt.Errorf("Unknown command %q", cmd)
    }

    return err
}

