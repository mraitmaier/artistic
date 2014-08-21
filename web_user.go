/*
   web_user.go
*/
package main

import (
	"fmt"
    "strings"
	"net/http"
//	"bitbucket.org/miranr/artistic/core"
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
			log.Error(fmt.Sprintf("Problem getting all users: %s", err.Error()))
			http.Redirect(w, r, "/error404", http.StatusFound)
			return
		}

		// create ad-hoc struct to be sent to page template
		var web = struct {
			User  *utils.User
			Users []utils.User
		} { user, users }

		// render the page
		err = aa.WebInfo.templates.ExecuteTemplate(w, "users", &web)
        if err != nil {
			log.Error("Cannot render the 'users' page.")
		}

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
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
                log.Error(err.Error())
			    http.Redirect(w, r, "/users", http.StatusFound)
            }

        case "POST":
            if err := postUserHandler(w, r, aa); err != nil {
                log.Error(err.Error())
            }
			http.Redirect(w, r, "/users", http.StatusFound)

        case "DELETE":
            id := mux.Vars(r)["id"]
            cmd := mux.Vars(r)["cmd"]
            t := new(utils.User)
            t.Id = db.MongoStringToId(id) // only valid ID needed to delete 
            if err := aa.DataProv.DeleteUser(t); err != nil {
                msg := fmt.Sprintf(
                    "%s user id=%q, DB returned %q.", cmd, id, err)
                log.Error(msg)
                return
            }
            log.Info(fmt.Sprintf("Successfully deleted user %q.", t.Id))
	        http.Redirect(w, r, "/users", http.StatusFound)

        case "PUT":
            fmt.Printf("received PUT request. :)\n")
        }

	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
    }) // return handler closure
}

//  HTTP GET handler for "/user/<cmd>" URLs.
func getUserHandler(w http.ResponseWriter, r *http.Request,
                        aa *ArtisticApp, user *utils.User) error {

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

    case "insert": // do nothing here...

    case "delete":
        s.Id = db.MongoStringToId(id) // only valid ID needed to delete 
        if err = aa.DataProv.DeleteUser(s); err != nil {
            return fmt.Errorf("%s user id=%q, DB returned %q.", cmd, id, err)
        }
        log.Info(fmt.Sprintf("Successfully deleted user %q.", s.Id))
	    http.Redirect(w, r, "/users", http.StatusFound)
        return nil //  this is all about deleting items...

    default:
        return fmt.Errorf("GET User handler: unknown command %q", cmd)
    }

	// create ad-hoc struct to be sent to page template
    var web = struct {
		User  *utils.User
        Cmd   string   // "view", "modify", "insert" or "delete"...
		UserProfile *utils.User
    }{ user, cmd, s }

    // render the page
	err = aa.WebInfo.templates.ExecuteTemplate(w, "user", &web)
    if err != nil {
	    log.Error("Error rendering the 'user' page.")
	}

    return err
}

func postUserHandler(w http.ResponseWriter, r *http.Request,
                            aa *ArtisticApp) error {

    // get data to modify 
    cmd := mux.Vars(r)["cmd"]

    var err error = nil

    switch cmd {

    case "insert":
        if err = insertNewUser(w, r, aa) ; err != nil {
            err = fmt.Errorf("Error creating user: %q.", err)
        }

    case "modify":
        if err = modifyExistingUser(w, r, aa); err != nil {
            err = fmt.Errorf("Modifying user: %q.", err)
        }

    case "changepwd":
        if err = changeUserPassword(w, r, aa); err != nil {
            err = fmt.Errorf("Changing password for user failed: %q.", err)
        }

    default:
        err = fmt.Errorf(
            "Invalid command %q for users. Redirecting to default page.", cmd)
    }

    return err
}

// modify an existing user handler function.
func modifyExistingUser(
            w http.ResponseWriter, r *http.Request, aa *ArtisticApp) error {

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
        return fmt.Errorf("invalid role.")
    }
    t.Id = db.MongoStringToId(id)
    t.Name = full
    t.Role = role
    t.Email = email

    // do it...
    if err = aa.DataProv.UpdateUser(t); err != nil {
        return err
    }
    aa.Log.Info(fmt.Sprintf("Successfully inserted new user %q.", name))
    return err
}

// create new user handler function.
func insertNewUser(
            w http.ResponseWriter, r *http.Request, aa *ArtisticApp) error {

    // get data to modify 
	id  := mux.Vars(r)["id"]

    // get POST form values and create a struct
	name  := strings.TrimSpace(r.FormValue("username"))
	pwd   := strings.TrimSpace(r.FormValue("password"))
	pwd2  := strings.TrimSpace(r.FormValue("password2"))
	role  := strings.TrimSpace(r.FormValue("role"))
	full  := strings.TrimSpace(r.FormValue("fullname"))
	email := strings.TrimSpace(r.FormValue("email"))

    if pwd != pwd2 {
        return fmt.Errorf("passwords do not match.")
    }

    var err error = nil
    // create a user and check passwords
    u, err := utils.NewUser(name, pwd, role);
    if err != nil {
        return err
    }
    u.Id = db.MongoStringToId(id)
    u.Name = full
    u.Role = role
    u.Email = email

    // do it...
    if err = aa.DataProv.InsertUser(u); err != nil {
       return err
    }

    aa.Log.Info(fmt.Sprintf("Successfully modified existing user %q.", name))
    return err
}

// Change password for existing user handler function.
func changeUserPassword(
            w http.ResponseWriter, r *http.Request, aa *ArtisticApp) error {

    var err error = nil

    // get data to modify 
	id  := mux.Vars(r)["id"]

    // get POST form values and create a struct
    old  := strings.TrimSpace(r.FormValue("oldpassword"))
	pwd  := strings.TrimSpace(r.FormValue("newpassword"))
	pwd2 := strings.TrimSpace(r.FormValue("newpassword2"))

	u, err := aa.DataProv.GetUser(id)
	if err != nil {
		return fmt.Errorf("Get user id=%q, DB returned %q.", id, err)
	}

    // check password first and return error if they're not valid
    if pwd != pwd2 {
        return fmt.Errorf("new passwords do not match.")
    }
    if !u.ComparePassword(old) {
        return fmt.Errorf("invalid old password.")
    }

    // now do it...
    u.SetPassword(pwd)
    if err = aa.DataProv.UpdateUser(u); err != nil {
        return err
    }

    aa.Log.Info(fmt.Sprintf(
            "Successfully changed password for existing user %q.", u.Username))
    return err
}

