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

/*
/////#################################################################
// Create a mongoDB URL from given username/password and other input info.
func CreateUrl(host string, port int,
               username, passwd string, dbname string) string {
    s := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", username, passwd,
                                                 host, port, dbname)
    return s
}

// Connect to mongoDB with given URL and timeout (in seconds) to stop trying
// when server is not available.
func Connect(url string, timeout time.Duration) (*mgo.Session, error) {
   return mgo.DialWithTimeout(url, timeout)
}

// Close the mongoDB connection. 
func Close(dbsess *mgo.Session) {
    if dbsess != nil {
        dbsess.Close()
    }
}

// Insert a new user into the DB. 
func MongoInsertUser(db *mgo.Database, u *utils.User) error {
    // TODO
    return nil
}

// Update an existing user in the DB. 
func MongoUpdateUser(db *mgo.Database, u *utils.User) error {
    // TODO
    return nil
}
*/

// Get a single user from the DB: we need a username. 
func (m * MongoDbConn) GetUser(username string) (*utils.User, error) {

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)

    if db != nil {
        u := utils.CreateUser("", "") // create empty user

        // get all users from DB
        err := db.C("users").Find(bson.M{"username": username }).One(&u)
        if err != nil { return nil, err }

        return u, nil
    }
    return nil,  errors.New("MongoDB descriptor empty.")
}

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

