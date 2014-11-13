//
// mongo.go -
//
package db

import (
    "fmt"
    "time"
    "errors"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
  //  "bitbucket.org/miranr/artistic/core"
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
        return (bson.ObjectIdHex(id))
    } else {
        return bson.ObjectId(id)
    }
}

// Create new ObjectId
func NewMongoId() bson.ObjectId {
    return bson.NewObjectId()
}

///////////////////////////// Users

// Retrieves all users from database.
func (m *MongoDbConn) GetAllUsers() ([]User, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil, errors.New("Mongo descriptor empty.") }

    // create channel
    ch := make(chan []User)

    // start a new goroutine to get users from DB
    go func(ch chan []User) {

        // check channel
        if ch == nil { return }

        // prepare the empty slice for users
        users := make([]User, 0)

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
func (m * MongoDbConn) GetUserByUsername(username string) (*User, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil, errors.New("MongoDB descriptor empty.") }

    // prepare channel
    ch := make(chan *User)

    // start goroutine to get a user
    go func(username string, ch chan *User) {

        //u := utils.CreateUser("", "") // create empty user
        u := NewUser()

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
/*
// Get a single user from the DB: we need an ID . 
func (m * MongoDbConn) GetUser(id string) (*utils.User, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil, errors.New("MongoDB descriptor empty.") }

    // prepare channel
    ch := make(chan *utils.User)

    // start goroutine to get a user
    go func(id string, ch chan *utils.User) {

        u := utils.CreateUser("", "") // create empty user

        // get a user from DB
        err := db.C("users").Find(bson.M{ "_id": MongoStringToId(id) }).One(&u)
        if err != nil {
            return
        }

        // write a user to channel
        ch <- u
    }(id, ch)

    // read user from channel
    user := <-ch

    return user, nil // all OK
}
*/

func (m * MongoDbConn) GetUser(id string) (*User, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil {
        return nil, errors.New("MongoDB descriptor empty.")
    }

    // create empty user
    u := NewUser()

    // get a user from DB
    if err := db.C("users").Find(bson.M{ "_id": MongoStringToId(id) }).One(&u); err != nil {
        return nil, err
    }

    return u, nil // all OK
}

// Update a single user in DB. 
func (m *MongoDbConn) UpdateUser(u *User) error {
    return m.adminUser(DBCmdUpdate, u)
}

// Create a new user in DB. 
func (m *MongoDbConn) InsertUser(u *User) error {
    // check the ID of the item to be inserted into DB
    if u.Id == "" {
        u.Id = NewMongoId()
    }
    return m.adminUser(DBCmdInsert, u)
}

// Delete a single user in DB. 
func (m *MongoDbConn) DeleteUser(u *User) error {
    return m.adminUser(DBCmdDelete, u)
}

// Aux method that administers the user records in DB
func (m *MongoDbConn) adminUser(cmd DbCommand, u *User) error {

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
        err = coll.UpdateId(u.Id, u)

    case DBCmdInsert:
        err = coll.Insert(u)

    case DBCmdDelete:
         err = coll.RemoveId(u.Id)

    default:
        err = fmt.Errorf("Handling users: Unknown command.")
    }
    return err
}

///////////////////////////// Datings

// Retrieves all datings from database.
func (m *MongoDbConn) GetAllDatings() ([]Dating, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil {
        return nil, errors.New("Get all datings: MongoDB descriptor empty.")
    }

    // get all users from DB
    d := make([]Dating, 0)
    if err := db.C("datings").Find(bson.D{}).All(&d); err != nil {
        return nil, err
    }
    return d, nil
}

// Retrieve a single Dating record by ID.
func (m *MongoDbConn) GetDating(id string) (*Dating, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil {
        return nil, errors.New("Get a dating: MongoDb descriptor empty.")
    }

    t := new(Dating)
    err := db.C("datings").Find(bson.M{ "_id": MongoStringToId(id) }).One(&t)
    if err != nil { return nil, err }

    return t, nil
}

// Update a single dating in DB. 
func (m * MongoDbConn) UpdateDating(d *Dating) error {

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
func (m *MongoDbConn) GetAllStyles() ([]Style, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil,  errors.New("MongoDB descriptor empty.") }

    // prepare the empty slice for users
    s := make([]Style, 0)

    // get all users from DB
    if err := db.C("styles").Find(bson.D{}).All(&s); err != nil {
        return nil, err
    }
    return s, nil
}

// Retrieve a single Style record.
func (m *MongoDbConn) GetStyle(id string) (*Style, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil {
        return nil, errors.New("Update a style: MongoDb descriptor empty.")
    }

    s := NewStyle()
    err := db.C("styles").Find(bson.M{ "_id" : MongoStringToId(id) }).One(&s)
    if err != nil {
        return nil, err
    }

    return s, nil
}

// Update a single style in DB. 
func (m * MongoDbConn) UpdateStyle(s *Style) error {
    return m.adminStyle(DBCmdUpdate, s)
}

// Create a new style in DB. 
func (m * MongoDbConn) InsertStyle(s *Style) error {
    // check the ID of the item to be inserted into DB
    if s.Id == "" {
        s.Id = NewMongoId()
    }
    return m.adminStyle(DBCmdInsert, s)
}

// Delete a single style in DB
func (m *MongoDbConn) DeleteStyle(s *Style) error {
    return m.adminStyle(DBCmdDelete, s)
}

// Aux method that administers the style records in DB
func (m *MongoDbConn) adminStyle(cmd DbCommand, s *Style) error {

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

        case DBCmdInsert:
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
func (m *MongoDbConn) GetAllTechniques() ([]Technique, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil {
        return nil, errors.New("Getting techniques: MongoDB descriptor empty.")
    }

    // get all users from DB
    t := make([]Technique, 0)
    if err := db.C("techniques").Find(bson.D{}).All(&t); err != nil {
        return nil, err
    }
    return t, nil
}

// Retrieve a single Technique record.
func (m *MongoDbConn) GetTechnique(id string) (*Technique, error) {

    dblock.Lock()
    defer dblock.Unlock()

    db := m.Sess.DB(m.name)
    if db == nil {
        return nil, errors.New("Getting a technique: MongoDb descriptor empty.")
    }

    t := new(Technique)
    err := db.C("techniques").Find(bson.M{ "_id": MongoStringToId(id) }).One(&t)
    if err != nil {
        return nil, err
    }

    return t, nil
}

// Update a single technique in DB. 
func (m * MongoDbConn) UpdateTechnique(t *Technique) error {
    return m.adminTechnique(DBCmdUpdate, t)
}

// Create a new technique in DB. 
func (m * MongoDbConn) InsertTechnique(t *Technique) error {
    // check the ID of the item to be inserted into DB
    if t.Id == "" {
        t.Id = NewMongoId()
    }
    return m.adminTechnique(DBCmdInsert, t)
}

// Delete a new technique in DB. 
func (m * MongoDbConn) DeleteTechnique(t *Technique) error {
    return m.adminTechnique(DBCmdDelete, t)
}

// Aux method that administers the technique records in DB
func (m *MongoDbConn) adminTechnique(cmd DbCommand, t *Technique) error {

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

    case DBCmdInsert:
        err = coll.Insert(t)

    case DBCmdDelete:
         err = coll.RemoveId(t.Id)

    default:
        err = fmt.Errorf("Handling techniques: Unknown DB command.")
    }

    return err
}

///////////////////////////// Painters
func (m *MongoDbConn) GetAllArtists(t ArtistType) ([]Artist, error) {

    // acquire DB lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil, errors.New("Mongo descriptor empty.") }

    // create channel
    ch := make(chan []Artist)

    // start a new goroutine to get users from DB
    go func(ch chan []Artist) {

        // check channel
        if ch == nil { return }

        // prepare the empty slice for users
        artists := make([]Artist, 0)
        var err error

        // get all artists from DB
        switch t {

        case ArtistTypePainter:
            if err = db.C("artists").Find( bson.M{ "is_painter" : true } ).All(&artists); err != nil {
                return
            }

        case ArtistTypeSculptor:
            if err = db.C("artists").Find( bson.M{ "is_sculptor" : true } ).All(&artists); err != nil {
                return
            }

        case ArtistTypeArchitect:
            if err = db.C("artists").Find( bson.M{ "is_architect" : true } ).All(&artists); err != nil {
                return
            }

        case ArtistTypePrintmaker:
            if err = db.C("artists").Find( bson.M{ "is_printmaker" : true } ).All(&artists); err != nil {
                return
            }

        case ArtistTypeCeramicist:
            if err = db.C("artists").Find( bson.M{ "is_ceramicist" : true } ).All(&artists); err != nil {
                return
            }
        }

        // write the users to the channel
        ch <- artists

    }(ch)

    // read the answer from channel
    artists := <-ch

    return artists, nil // OK
}

// Get a single artist from the DB: we need an ID . 
func (m * MongoDbConn) GetArtist(id string) (*Artist, error) {

    // acquire lock
    dblock.Lock()
    defer dblock.Unlock()

    // get *mgo.Database instance
    db := m.Sess.DB(m.name)
    if db == nil { return nil, errors.New("MongoDB descriptor empty.") }

    // prepare channel
    ch := make(chan *Artist)

    // start goroutine to get a user
    go func(id string, ch chan *Artist) {

        u := NewArtist() // create empty user

        // get a user from DB
        //var dbid DbIdentifier
        //dbid.StringToId(id)
        err := db.C("artists").Find(bson.M{ "_id": MongoStringToId(id) }).One(&u)
        //err := db.C("artists").Find(bson.M{ "_id": id }).One(&u)
        if err != nil {
            return
        }

        // write a user to channel
        ch <- u
    }(id, ch)

    // read user from channel
    user := <-ch

    return user, nil // all OK
}
// Update a single artist in DB. 
func (m * MongoDbConn) UpdateArtist(a *Artist) error {
    return m.adminArtist(DBCmdUpdate, a)
}

// Create a new artist in DB. 
func (m * MongoDbConn) InsertArtist(a *Artist) error {
    // check the ID of the item to be inserted into DB
    if a.Id == "" {
        a.Id = NewMongoId()
    }
    return m.adminArtist(DBCmdInsert, a)
}

// Delete a new artist in DB. 
func (m * MongoDbConn) DeleteArtist(a *Artist) error {
    return m.adminArtist(DBCmdDelete, a)
}

// Aux method that administers the artist records in DB
func (m *MongoDbConn) adminArtist(cmd DbCommand, a *Artist) error {

    dblock.Lock()
    defer dblock.Unlock()

    coll := m.Sess.DB(m.name).C("artists")
    if coll == nil {
        return  fmt.Errorf("Handling an artist: MongoDB descriptor empty.")
    }

    if a == nil {
       return fmt.Errorf("Handling an artist: cannot create empty artist.")
    }

    var err error
    switch cmd {

    case DBCmdUpdate:
        err = coll.UpdateId(a.Id, a)

    case DBCmdInsert:
        err = coll.Insert(a)

    case DBCmdDelete:
         err = coll.RemoveId(a.Id)

    default:
        err = fmt.Errorf("Handling artists: Unknown DB command.")
    }

    return err
}

