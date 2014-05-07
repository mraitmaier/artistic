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
)

const (
    // a (quite) random string that is used as a key for sessions
    sessKey = `iufwnwieh3436SiKJSJo90e3jdiejdlje3+'0%$#!)dlkjja(!~§<sdfad$io*"`
)

// authenticate the user with given username and password
func authenticateUser(u *utils.User,
                      w http.ResponseWriter, r *http.Request) (bool, error) {

    // create new session ID
    id := newSessId()
//    fmt.Printf("DEBUG: session ID=%q\n", id) // DEBUG

    // get information from DB
    query := bson.M{ "username" : u.Username }
    err := ac.dbsess.DB("artistic").C("users").Find(query).One(u)
    if err != nil {
fmt.Printf("ERROR: user=%s\n", u.String()) // DEBUG
        return false, err
    //if cnt, err := ac.dbsess.DB("artistic").C("users").Count(); err != nil {
    //if err := db.C("users").Find(bson.D{}).All(&u); err != nil {
    //    fmt.Printf("DEBUG count=%d\n", cnt)
    //} else {
    //    fmt.Printf("DEBUG found user=%v\n", u) // DEBUG
    }
    fmt.Printf("DEBUG: user=%s\n", u.String()) // DEBUG

    // get current session data; will create new session with given random ID
    s, err := store.Get(r, "artistic")
    if err != nil { return false, err }
    s.Values["sessid"] = id

    // create a new file in sessions folder to indicate valid session; we don't
    // care about the descriptor
    _, err = os.Create(filepath.Join(ac.sessDir, id))
    if err != nil { return false, err }

    // save the session data
    s.Save(r, w)

    return true, nil
}

func logout(u *utils.User, r *http.Request) error {

    // get current session data; retrieve session ID
    s, err := store.Get(r, "artistic")
    if err != nil { return err }
    id := s.Values["sessid"]

    // user has a unique session ID and there should be the file with this ID
    // in the sessions folder. 
    // Delete it, if it exists. 
    // If it doesn't exist, there's probably something wrong: do nothing.
    f := filepath.Join(ac.sessDir, id.(string))
    if utils.FileExists(f) {
        os.Remove(f)
    }
    return nil
}

// check if user is already authenticated 
func userIsAuthenticated(r *http.Request) bool {

    s, err := store.Get(r, "session")
    if err != nil {
    }

    // get a session ID 
    sessid := s.Values["session-id"]
    fmt.Printf("Session ID: %v\n", sessid)

    //return false
    return true
}

// generate unique session ID; return it as string
func newSessId() string {

    // generate pseudo-random int64
    num := rand.Int63()

    // now hash the random int64 value with SHA512
    hash := sha512.Sum512(int64ToBytes(num))

    return fmt.Sprintf("%x", hash)
}

