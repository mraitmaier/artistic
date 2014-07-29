//
// mongo.go -
//
package db

import (
    "fmt"
    "time"
    "errors"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
    "bitbucket.org/miranr/artistic/utils"
    "bitbucket.org/miranr/artistic/core"
)

// Type representing the MongoDB Session. Implements the DbConnector interface.
type MongoDbConn struct {

    // name of the database
    name string

    // open connection (session) to MongoDB 
    Sess *mgo.Session
}

/* DbConnector interface implementation */

// Connect to mongoDB with given URL and timeout (in seconds) to stop trying
// when server is not available.
func (m *MongoDbConn) Connect(url string, timeout time.Duration) (e error) {

    m.Sess, e = mgo.DialWithTimeout(url, timeout)

    return e
}

// Close the mongoDB connection. 
func (m *MongoDbConn) Close() {
    if m.Sess != nil {
        m.Sess.Close()
    }
}



/* DataProvider interface implementation */

// Convert BSON Object ID to string hex representation.
func MongoIdToString(id bson.ObjectId) string {
    return id.Hex()
}

// Convert ID from hex string representation to BSON Object ID.
func MongoStringToId(id string) bson.ObjectId {
    return bson.ObjectIdHex(id)
}

///////////////////////////// Users

// Retrieves all users from database.
func (m *MongoDbConn) GetAllUsers() ([]utils.User, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil, errors.New("Mongo descriptor empty.") }

    // create channel
    ch := make(chan []utils.User)

    // start a new goroutine to get users from DB
    go func(ch chan []utils.User) {

        // check channel
        if ch == nil { return }

        // prepare the empty slice for users
        users := make([]utils.User, 0)

        // get all users from DB
        if err := db.C("users").Find(bson.D{}).All(&users); err != nil {
            return
        }

        // write the users to the channel
        ch <- users

    }(ch)

    // read the answer from channel
    users := <-ch

    return users, nil // OK
}

// Get a single user from the DB: we need a username. 
func (m * MongoDbConn) GetUser(username string) (*utils.User, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil, errors.New("MongoDB descriptor empty.") }

    // prepare channel
    ch := make(chan *utils.User)

    // start goroutine to get a user
    go func(username string, ch chan *utils.User) {

        u := utils.CreateUser("", "") // create empty user

        // get a user from DB
        err := db.C("users").Find(bson.M{ "username": username }).One(&u)
        if err != nil {
            return
        }

        // write a user to channel
        ch <- u
    }(username, ch)

    // read user from channel
    user := <-ch

    return user, nil // all OK
}

// Update a single user in DB. 
func (m * MongoDbConn) UpdateUser(u *utils.User) error {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return errors.New("MongoDB descriptor empty.") }

    // FIXME: update, not get...
    err := db.C("users").Find(bson.M{ "username": u.Username }).One(&u)

    if err != nil {
        return err
    }
    return nil
}

///////////////////////////// Datings

// Retrieves all datings from database.
func (m *MongoDbConn) GetAllDatings() ([]core.Dating, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil {
        return nil,  errors.New("Get all datings: MongoDB descriptor empty.")
    }

    // prepare the empty slice for users
    d := make([]core.Dating, 0)

    // get all users from DB
    if err := db.C("datings").Find(bson.D{}).All(&d); err != nil {
        return nil, err
    }
    //fmt.Printf("DEBUG: %v\n", d) // DEBUG
    return d, nil
}

// Retrieve a single Dating record by ID.
func (m *MongoDbConn) GetDating(id string) (*core.Dating, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil {
        return nil, errors.New("Get a dating: MongoDb descriptor empty.")
    }

    t := new(core.Dating)

    //err := db.C("datings").Find(bson.M{ "_id": bson.ObjectIdHex(id) }).One(&t)
    err := db.C("datings").Find(bson.M{ "_id": MongoStringToId(id) }).One(&t)
    if err != nil { return nil, err }

    return t, nil
}

// Update a single dating in DB. 
func (m * MongoDbConn) UpdateDating(d *core.Dating) error {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    if d == nil {
        return fmt.Errorf("Update a dating: cannot update empty dating.")
    }

    db := m.Sess.DB(m.name)
    if db == nil {
        return  fmt.Errorf("Update a dating: MongoDB descriptor empty.")
    }

    // update the dating in DB
    if err := db.C("datings").Update(bson.M{ "_id": d.Id }, d); err != nil {
        return err
    }

    return nil
}

///////////////////////////// Styles

// Retrieves all styles from DB with given DB descriptor.
func (m *MongoDbConn) GetAllStyles() ([]core.Style, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil,  errors.New("MongoDB descriptor empty.") }

    // prepare the empty slice for users
    s := make([]core.Style, 0)

    // get all users from DB
    if err := db.C("styles").Find(bson.D{}).All(&s); err != nil {
        return nil, err
    }
    return s, nil
}

// Retrieve a single Style record.
func (m *MongoDbConn) GetStyle(name string) (*core.Style, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil, errors.New("MongoDb descriptor empty.") }

    s := new(core.Style)

    err := db.C("styles").Find(bson.M{ "name" : name }).One(&s)
    if err != nil { return nil, err }

    return s, nil
}

///////////////////////////// Techniques

// Retrieves all techniques from DB with given DB descriptor.
func (m *MongoDbConn) GetAllTechniques() ([]core.Technique, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil,  errors.New("MongoDB descriptor empty.") }

    // prepare the empty slice for users
    t := make([]core.Technique, 0)

    // get all users from DB
    if err := db.C("techniques").Find(bson.D{}).All(&t); err != nil {
        return nil, err
    }

    return t, nil
}

// Retrieve a single Technique record.
func (m *MongoDbConn) GetTechnique(name string) (*core.Technique, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil, errors.New("MongoDb descriptor empty.") }

    t := new(core.Technique)

    err := db.C("techniques").Find(bson.M{ "name" : name }).One(&t)
    if err != nil { return nil, err }

    return t, nil
}

