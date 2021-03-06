package main

// auth.go

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"github.com/mraitmaier/artistic/db"
	"github.com/mraitmaier/artistic/utils"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// authenticate the user with given username and password
func authenticateUser(u, p string, aa *ArtisticApp, w http.ResponseWriter, r *http.Request) (bool, error) {

	// create new session ID
	id := newSessID()
	//    fmt.Printf("DEBUG: session ID=%q\n", id) // DEBUG

	// get information from DB
	userdb, err := aa.DataProv.GetUserByUsername(u)
	if err != nil {
		return false, err
	}

	// fmt.Printf("DEBUG authenticateUser(): %v\n", userdb) // DEBUG
	// if passwords match....
	if userdb != nil && userdb.ComparePassword(p) {

		// get current session data; create new session with given random ID
		s, err := aa.WebInfo.store.Get(r, "artistic")
		if err != nil {
			return false, err
		}
		s.Values["sessid"] = id
		s.Values["user"] = u

		// create a new file in sessions folder to indicate valid session;
		// we don't care about the descriptor
		f, err := os.Create(filepath.Join(aa.WebInfo.sessDir, id))
		if err != nil {
			return false, err
		}
		f.Close()

		// save the session data
		s.Save(r, w)

		return true, nil
	}
	return false, fmt.Errorf("Passwords do not match.")
}

// Perform a user logout.
func logout(aa *ArtisticApp, w http.ResponseWriter, r *http.Request) error {

	// get current session data; retrieve session ID
	s, err := aa.WebInfo.store.Get(r, "artistic")
	if err != nil {
		return err
	}
	id := s.Values["sessid"].(string) // get session ID
	//   name := s.Values["user"].(string) // get username

	// user has a unique session ID and there should be the file with this ID in the sessions folder.
	// Delete it, if it exists. If it doesn't exist, there's probably something wrong: do nothing.
	f := filepath.Join(aa.WebInfo.sessDir, id)
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
func userIsAuthenticated(aa *ArtisticApp, r *http.Request) (bool, *db.User) {

	s, err := aa.WebInfo.store.Get(r, "artistic")
	if err != nil {
		return false, nil
	}

	// get a session ID
	id := s.Values["sessid"]

	// if only the id value is not empty...
	if id != nil {
		f := filepath.Join(aa.WebInfo.sessDir, id.(string))
		if utils.FileExists(f) {

			// get user information
			user, err := aa.DataProv.GetUserByUsername(s.Values["user"].(string))
			if err != nil { // something is not OK...
				return false, nil
			}

			return true, user
		}
	}
	return false, nil
}

// generate unique session ID; return it as string
func newSessID() string {

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
