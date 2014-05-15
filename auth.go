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
    "bytes"
    "encoding/binary"
    "time"
//    "labix.org/v2/mgo/bson"
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
    userdb, err := db.MongoGetUser(aa.DbSess.DB("artistic"), u)
    if err != nil {
        return false, err
    }

    // if passwords match....
    if utils.ComparePasswords(userdb.Password, p) {

        // get current session data; create new session with given random ID
        s, err := store.Get(r, "artistic")
        if err != nil { return false, err }
        s.Values["sessid"] = id
        s.Values["user"] = u

        // create a new file in sessions folder to indicate valid session; 
        // we don't care about the descriptor
        f, err := os.Create(filepath.Join(aa.WebInfo.SessDir, id))
        if err != nil { return false, err }
        f.Close()

        // save the session data
        s.Save(r, w)

        fmt.Printf("DEBUG creating Session: %v\n", s) // DEBUG

        return true, nil
    }
    return false, nil // no error, but passwords do not match
}

func logout(w http.ResponseWriter, r *http.Request) error {

    // get current session data; retrieve session ID
    s, err := store.Get(r, "artistic")
    if err != nil {
        return err
    }
    id := s.Values["sessid"].(string) // get session ID
 //   name := s.Values["user"].(string) // get username

    // user has a unique session ID and there should be the file with this ID
    // in the sessions folder. 
    // Delete it, if it exists. 
    // If it doesn't exist, there's probably something wrong: do nothing.
    f := filepath.Join(aa.WebInfo.SessDir, id)
    if utils.FileExists(f) {
        os.Remove(f)
        // remove values from session and save
        delete(s.Values, "user")
        delete(s.Values, "sessid")
        s.Save(r, w)
    }
    return nil
}

// check if user is already authenticated 
func userIsAuthenticated(r *http.Request) (bool, *utils.User) {

    s, err := store.Get(r, "artistic")
    if err != nil {
        return false, nil
    }

    // get a session ID 
    id := s.Values["sessid"]

    // if only the id value is not empty...
    if id != nil {
        f := filepath.Join(aa.WebInfo.SessDir, id.(string))
        if utils.FileExists(f) {

            // get user information
            user, err := db.MongoGetUser(aa.DbSess.DB("artistic"),
                                  s.Values["user"].(string))
            if  err != nil { // something is not OK...
                return false, nil
            }

            return true, user
        }
    }
    return false, nil
}

/*
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
*/

// generate unique session ID; return it as string
func newSessId() string {

    // generate pseudo-random int64, seed is current time in nanoseconds 
    rand.Seed(time.Now().UnixNano())
    num := rand.Int63()

    // now hash the random int64 value with SHA512
    hash := sha512.Sum512(int64ToBytes(num))

    return fmt.Sprintf("%x", hash)
}

// Converts 64-bit integer value into byte buffer.
func int64ToBytes(i int64) []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, i)
    return buf.Bytes()
}
