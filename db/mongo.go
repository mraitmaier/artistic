//
// mongo.go -
//
package db

import (
    "fmt"
    "time"
    "errors"
//    "labix.org/v2/mgo"
 //   "labix.org/v2/mgo/bson"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
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
    if id != "" {
        return bson.ObjectIdHex(id)
    } else {
        return bson.ObjectId(id)
    }
}

// Create new ObjectId
func NewId() bson.ObjectId {
    return bson.NewObjectId()
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
func (m *MongoDbConn) UpdateUser(u *utils.User) error {
    return m.adminUser(DBCmdUpdate, u)
}

// Create a new user in DB. 
func (m *MongoDbConn) CreateUser(u *utils.User) error {
    return m.adminUser(DBCmdCreate, u)
}

// Delete a single user in DB. 
func (m *MongoDbConn) DeleteUser(u *utils.User) error {
    return m.adminUser(DBCmdDelete, u)
}
// Aux method that administers the user records in DB
func (m *MongoDbConn) adminUser(cmd DbCommand, u *utils.User) error {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    coll := m.Sess.DB(m.name).C("users")
    if coll == nil {
        return  fmt.Errorf("Handling a user: MongoDB descriptor empty.")
    }

    if u == nil {
        return fmt.Errorf("Handling a user: cannot create empty style.")
    }

    var err error
    switch cmd {

        case DBCmdUpdate:
            err = coll.Update(bson.M { "_id" : u.Id }, u)

        case DBCmdCreate:
            err = coll.Insert(u)

        case DBCmdDelete:
             err = coll.RemoveId(u.Id)

        default:
            err = fmt.Errorf("Handling styles: Unknown command.")
    }

    return err
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
        return nil, errors.New("Get all datings: MongoDB descriptor empty.")
    }

    // prepare the empty slice for users
    d := make([]core.Dating, 0)

    // get all users from DB
    if err := db.C("datings").Find(bson.D{}).All(&d); err != nil {
        return nil, err
    }

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
func (m *MongoDbConn) GetStyle(id string) (*core.Style, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil {
        return nil, errors.New("Update a style: MongoDb descriptor empty.")
    }

    s := core.NewStyle("", "")
    err := db.C("styles").Find(bson.M{ "_id" : MongoStringToId(id) }).One(&s)
    if err != nil {
        return nil, err
    }

    return s, nil
}

// Update a single style in DB. 
func (m * MongoDbConn) UpdateStyle(s *core.Style) error {
    return m.adminStyle(DBCmdUpdate, s)
}

// Create a new style in DB. 
func (m * MongoDbConn) CreateStyle(s *core.Style) error {
    // check the ID of the item to be inserted into DB
    if s.Id == "" {
        s.Id = NewId()
    }
    return m.adminStyle(DBCmdCreate, s)
}

// Delete a single style in DB
func (m *MongoDbConn) DeleteStyle(s *core.Style) error {
    return m.adminStyle(DBCmdDelete, s)
}

// Aux method that administers the style records in DB
func (m *MongoDbConn) adminStyle(cmd DbCommand, s *core.Style) error {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    coll := m.Sess.DB(m.name).C("styles")
    if coll == nil {
        return  fmt.Errorf("Handling a style: MongoDB descriptor empty.")
    }

    if s == nil {
        return fmt.Errorf("Handling a style: cannot create empty style.")
    }

    var err error
    switch cmd {

        case DBCmdUpdate:
            err = coll.UpdateId(s.Id, s)

        case DBCmdCreate:
            err = coll.Insert(s)

        case DBCmdDelete:
             err = coll.RemoveId(s.Id)

        default:
            err = fmt.Errorf("Handling styles: Unknown command.")
    }

    return err
}

///////////////////////////// Techniques

// Retrieves all techniques from DB with given DB descriptor.
func (m *MongoDbConn) GetAllTechniques() ([]core.Technique, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil {
        return nil, errors.New("Getting techniques: MongoDB descriptor empty.")
    }

    // prepare the empty slice for users
    t := make([]core.Technique, 0)

    // get all users from DB
    if err := db.C("techniques").Find(bson.D{}).All(&t); err != nil {
        return nil, err
    }

    return t, nil
}

// Retrieve a single Technique record.
func (m *MongoDbConn) GetTechnique(id string) (*core.Technique, error) {

    dblock.Lock()
    defer dblock.Unlock()

    db := m.Sess.DB(m.name)
    if db == nil {
        return nil, errors.New("Getting a technique: MongoDb descriptor empty.")
    }

    t := new(core.Technique)
    err := db.C("techniques").Find(bson.M{ "_id": MongoStringToId(id) }).One(&t)
    if err != nil {
        return nil, err
    }

    return t, nil
}

// Update a single technique in DB. 
func (m * MongoDbConn) UpdateTechnique(t *core.Technique) error {
    return m.adminTechnique(DBCmdUpdate, t)
}

// Create a new technique in DB. 
func (m * MongoDbConn) CreateTechnique(t *core.Technique) error {
    // check the ID of the item to be inserted into DB
    if t.Id == "" {
        t.Id = NewId()
    }
    return m.adminTechnique(DBCmdCreate, t)
}

// Delete a new technique in DB. 
func (m * MongoDbConn) DeleteTechnique(t *core.Technique) error {
    return m.adminTechnique(DBCmdDelete, t)
}

// Aux method that administers the technique records in DB
func (m *MongoDbConn) adminTechnique(cmd DbCommand, t *core.Technique) error {

    dblock.Lock()
    defer dblock.Unlock()

    coll := m.Sess.DB(m.name).C("techniques")
    if coll == nil {
        return  fmt.Errorf("Handling a technique: MongoDB descriptor empty.")
    }

    if t == nil {
       return fmt.Errorf("Handling a technique: cannot create empty technique.")
    }

    var err error
    switch cmd {

    case DBCmdUpdate:
        err = coll.UpdateId(t.Id, t)

    case DBCmdCreate:
        err = coll.Insert(t)

    case DBCmdDelete:
         err = coll.RemoveId(t.Id)

    default:
        err = fmt.Errorf("Handling styles: Unknown command.")
    }

    return err
}


////////////////////// Experimental /////////////////
// NOTE: how to abstract away the DB ID (for differents DBs)? By implementing 
// a type that satisfies the interface...
//
// implementing the DbIdentifier interface for MongoDb
type MongoDbId bson.ObjectId

func (m MongoDbId) IdToString(id interface{}) string {
    s := id.(bson.ObjectId)
    return s.String()
}

func (m MongoDbId) StringToId(s string) interface{} {
    return interface{}(bson.ObjectIdHex(s))
}

func (m MongoDbId) New() MongoDbId {
    return MongoDbId(bson.NewObjectId())
}

