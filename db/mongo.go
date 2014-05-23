//
// mongo.go -
//
package db

import (
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

///////////////////////////// Users

// Retrieves all users from database.
func (m *MongoDbConn) GetAllUsers() ([]utils.User, error) {

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)

    if db != nil {
        // prepare the empty slice for users
        u := make([]utils.User, 0)

        // get all users from DB
        if err := db.C("users").Find(bson.D{}).All(&u); err != nil {
            return nil, err
        }

        return u, nil
    }
    return nil,  errors.New("MongoDB descriptor empty.")
}

// Get a single user from the DB: we need a username. 
func (m * MongoDbConn) GetUser(username string) (*utils.User, error) {

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)

    if db != nil {
        u := utils.CreateUser("", "") // create empty user

        // get all users from DB
        err := db.C("users").Find(bson.M{ "username": username }).One(&u)
        if err != nil { return nil, err }

        return u, nil
    }
    return nil,  errors.New("MongoDB descriptor empty.")
}

// Update a single user in DB. 
func (m * MongoDbConn) UpdateUser(u *utils.User) error {
    return nil
}

///////////////////////////// Datings

// Retrieves all datings from database.
func (m *MongoDbConn) GetAllDatings() ([]core.Dating, error) {

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)

    if db != nil {
        // prepare the empty slice for users
        d := make([]core.Dating, 0)

        // get all users from DB
        if err := db.C("datings").Find(bson.D{}).All(&d); err != nil {
            return nil, err
        }

        return d, nil
    }
    return nil,  errors.New("MongoDB descriptor empty.")
}

// Retrieve a single Technique record.
func (m *MongoDbConn) GetDating(name string) (*core.Dating, error) {

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)

    if db != nil {

        t := new(core.Dating)

        err := db.C("datings").Find(bson.M{ "dating" : name }).One(&t)
        if err != nil { return nil, err }

        return t, nil
    }
    return nil, errors.New("MongoDb descriptor empty.")
}

// Update a single user in DB. 
func (m * MongoDbConn) UpdateDating(d *core.Dating) error {
    return nil
}

///////////////////////////// Styles

// Retrieves all styles from DB with given DB descriptor.
func (m *MongoDbConn) GetAllStyles() ([]core.Style, error) {

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)

    if db != nil {
        // prepare the empty slice for users
        s := make([]core.Style, 0)

        // get all users from DB
        if err := db.C("styles").Find(bson.D{}).All(&s); err != nil {
            return nil, err
        }

        return s, nil
    }
    return nil,  errors.New("MongoDB descriptor empty.")
}

// Retrieve a single Style record.
func (m *MongoDbConn) GetStyle(name string) (*core.Style, error) {

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)

    if db != nil {

        s := new(core.Style)

        err := db.C("styles").Find(bson.M{ "name" : name }).One(&s)
        if err != nil { return nil, err }

        return s, nil
    }
    return nil, errors.New("MongoDb descriptor empty.")
}

///////////////////////////// Techniques

// Retrieves all techniques from DB with given DB descriptor.
func (m *MongoDbConn) GetAllTechniques() ([]core.Technique, error) {

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)

    if db != nil {
        // prepare the empty slice for users
        t := make([]core.Technique, 0)

        // get all users from DB
        if err := db.C("techniques").Find(bson.D{}).All(&t); err != nil {
            return nil, err
        }

        return t, nil
    }
    return nil,  errors.New("MongoDB descriptor empty.")
}

// Retrieve a single Technique record.
func (m *MongoDbConn) GetTechnique(name string) (*core.Technique, error) {

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)

    if db != nil {

        t := new(core.Technique)

        err := db.C("techniques").Find(bson.M{ "name" : name }).One(&t)
        if err != nil { return nil, err }

        return t, nil
    }
    return nil, errors.New("MongoDb descriptor empty.")
}

