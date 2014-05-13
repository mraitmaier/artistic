/*
    auth.go
 */
package main

import (
    "os"
    "fmt"
    "net/http"
    "path/filepath"
    "math/rand"
    "crypto/sha512"
    "labix.org/v2/mgo/bson"
 //   "labix.org/v2/mgo"
    "bitbucket.org/miranr/artistic/utils"
    "bitbucket.org/miranr/artistic/db"
)

// authenticate the user with given username and password
func authenticateUser(u, p string,
                      w http.ResponseWriter, r *http.Request) (bool, error) {

    // create new session ID
    id := newSessId()
//    fmt.Printf("DEBUG: session ID=%q\n", id) // DEBUG

    // get information from DB
    userdb := utils.CreateUser("", "")
    query := bson.M{ "username" : u }
    err := aa.DbSess.DB("artistic").C("users").Find(query).One(userdb)
    if err != nil {
        //fmt.Printf("ERROR: user=%s\n", u.String()) // DEBUG
        return false, err
    }
//    fmt.Printf("DEBUG: user=%s\n", userdb.String()) // DEBUG

    // if passwords match....
    if utils.ComparePasswords(userdb.Password, p) {

        // get current session data; create new session with given random ID
        s, err := store.Get(r, "artistic")
        if err != nil { return false, err }
        s.Values["sessid"] = id
        s.Values["user"] = u

        // create a new file in sessions folder to indicate valid session; 
        // we don't care about the descriptor
        _, err = os.Create(filepath.Join(aa.WebInfo.SessDir, id))
        if err != nil { return false, err }

        // save the session data
        s.Save(r, w)

        fmt.Printf("DEBUG creating Session: %v\n", s) // DEBUG

        return true, nil
    }
    return false, nil // no error, but passwords do not match
}

//func logout(u *utils.User, r *http.Request) error {
func logout(r *http.Request) error {

    // get current session data; retrieve session ID
    s, err := store.Get(r, "artistic")
    if err != nil { return err }
    id := s.Values["sessid"]

    // user has a unique session ID and there should be the file with this ID
    // in the sessions folder. 
    // Delete it, if it exists. 
    // If it doesn't exist, there's probably something wrong: do nothing.
    f := filepath.Join(aa.WebInfo.SessDir, id.(string))
    if utils.FileExists(f) {
        os.Remove(f)
    }
    return nil
}

// check if user is already authenticated 
func userIsAuthenticated(r *http.Request) bool {

    s, err := store.Get(r, "artistic")
    if err != nil {
        return false
    }

    fmt.Printf("DEBUG Session: %v\n", s) // DEBUG

    // get a session ID 
    id := s.Values["sessid"]

    f := filepath.Join(aa.WebInfo.SessDir, id.(string))
    if utils.FileExists(f) {
        return true
    }
    return false
}

// get a User instance for current session
func getUser(r *http.Request) *utils.User {
    s, err := store.Get(r, "artistic")
    if err != nil { return nil }

    _db := aa.DbSess.DB("artistic")
    u, err := db.MongoGetUser(_db, s.Values["user"].(string))
    if err != nil { return nil }
    return u
}

// get a username for current session
func getUsername(r *http.Request) string {

    s, err := store.Get(r, "artistic")
    if err != nil { return "Error!" }

    return s.Values["user"].(string)
}

// generate unique session ID; return it as string
func newSessId() string {

    // generate pseudo-random int64
    num := rand.Int63()

    // now hash the random int64 value with SHA512
    hash := sha512.Sum512(int64ToBytes(num))

    return fmt.Sprintf("%x", hash)
}

