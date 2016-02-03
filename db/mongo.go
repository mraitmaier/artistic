//
// mongo.go -
//
package db

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
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

/*
func CheckDefaultUser(m *MongoDbConn) error {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("users")
	if coll == nil {
		return fmt.Errorf("Creating a default user: MongoDB descriptor empty.")
	}

	cnt, err := coll.Count()
	if err != nil {
		return fmt.Errorf("Creating a default user: Error counting collection.")
	}

	// create default user only if the collection is empty...
	if cnt == 0 {

		u := CreateUser(DefAppUsername, DefAppPasswd)
        u.SetRole("admin")
		err = coll.Insert(u)
        u.Name = "Change MyName"
        u.Email = "change_me@somewhere.org"
        coll.In
	}
	return err
}
*/

// CountUsers returns the number of users currently present in database.
func (m *MongoDbConn) CountUsers() (int, error) {

	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("users")
	if coll == nil {
		return -1, fmt.Errorf("Counting users: MongoDB descriptor empty.")
	}

	return coll.Count()
}

// Retrieves all users from database.
func (m *MongoDbConn) GetAllUsers() ([]*User, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Mongo descriptor empty.")
	}

	// create channel
	ch := make(chan []*User)

	// start a new goroutine to get users from DB
	go func(ch chan []*User) {

		// check channel
		if ch == nil {
			return
		}

		// prepare the empty slice for users
		users := make([]*User, 0)

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
func (m *MongoDbConn) GetUserByUsername(username string) (*User, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("MongoDB descriptor empty.")
	}

	// prepare channel
	ch := make(chan *User)

	// start goroutine to get a user
	go func(username string, ch chan *User) {

		//u := utils.CreateUser("", "") // create empty user
		u := NewUser()

		// get a user from DB
		err := db.C("users").Find(bson.M{"username": username}).One(&u)
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

func (m *MongoDbConn) GetUser(id string) (*User, error) {

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
	if err := db.C("users").Find(bson.M{"_id": MongoStringToId(id)}).One(&u); err != nil {
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
		return fmt.Errorf("Handling a user: MongoDB descriptor empty.")
	}

	if u == nil {
		return fmt.Errorf("Handling a user: cannot create empty style.")
	}

	var err error
	switch cmd {

	case DBCmdUpdate:
		u.Modified = NewTimestamp()
		err = coll.UpdateId(u.Id, u)

	case DBCmdInsert:
		u.Created = NewTimestamp()
		err = coll.Insert(u)

	case DBCmdDelete:
		err = coll.RemoveId(u.Id)

	default:
		err = fmt.Errorf("Handling users: Unknown command.")
	}
	return err
}

///////////////////////////// Datings

// CountDatings returns the number of datings currently present in database.
func (m *MongoDbConn) CountDatings() (int, error) {

	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("datings")
	if coll == nil {
		return -1, fmt.Errorf("Couting datings: MongoDB descriptor empty.")
	}

	return coll.Count()
}

// Retrieves all datings from database.
func (m *MongoDbConn) GetAllDatings() ([]*Dating, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Get all datings: MongoDB descriptor empty.")
	}

	// get all users from DB
	d := make([]*Dating, 0)
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
	err := db.C("datings").Find(bson.M{"_id": MongoStringToId(id)}).One(&t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Update a single dating in DB.
func (m *MongoDbConn) UpdateDating(d *Dating) error {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	if d == nil {
		return fmt.Errorf("Update a dating: cannot update empty dating.")
	}

	db := m.Sess.DB(m.name)
	if db == nil {
		return fmt.Errorf("Update a dating: MongoDB descriptor empty.")
	}

	// update the dating in DB
	d.Modified = NewTimestamp() // Update modified timestamp before commiting
	if err := db.C("datings").Update(bson.M{"_id": d.Id}, d); err != nil {
		return err
	}

	return nil
}

// Insert datings in DB.
func (m *MongoDbConn) InsertDatings(datings []*Dating) error {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	if datings == nil {
		return fmt.Errorf("Insert datings: cannot update empty dating.")
	}

	db := m.Sess.DB(m.name)
	if db == nil {
		return fmt.Errorf("Insert datings: MongoDB descriptor empty.")
	}

	// update the dating in DB
	for _, d := range datings {
		if err := db.C("datings").Insert(d); err != nil {
			return err
		}
	}
	return nil
}

///////////////////////////// Styles

// Retrieves all styles from DB with given DB descriptor.
func (m *MongoDbConn) GetAllStyles() ([]*Style, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("MongoDB descriptor empty.")
	}

	// prepare the empty slice for users
	s := make([]*Style, 0)

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
	err := db.C("styles").Find(bson.M{"_id": MongoStringToId(id)}).One(&s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Update a single style in DB.
func (m *MongoDbConn) UpdateStyle(s *Style) error {
	return m.adminStyle(DBCmdUpdate, s)
}

// Create a new style in DB.
func (m *MongoDbConn) InsertStyle(s *Style) error {
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
		return fmt.Errorf("Handling a style: MongoDB descriptor empty.")
	}

	if s == nil {
		return fmt.Errorf("Handling a style: cannot create empty style.")
	}

	var err error
	switch cmd {

	case DBCmdUpdate:
		s.Modified = NewTimestamp()
		err = coll.UpdateId(s.Id, s)

	case DBCmdInsert:
		s.Created = NewTimestamp()
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
func (m *MongoDbConn) GetAllTechniques() ([]*Technique, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting techniques: MongoDB descriptor empty.")
	}

	// get all users from DB
	t := make([]*Technique, 0)
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
	err := db.C("techniques").Find(bson.M{"_id": MongoStringToId(id)}).One(&t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Update a single technique in DB.
func (m *MongoDbConn) UpdateTechnique(t *Technique) error {
	return m.adminTechnique(DBCmdUpdate, t)
}

// Create a new technique in DB.
func (m *MongoDbConn) InsertTechnique(t *Technique) error {
	// check the ID of the item to be inserted into DB
	if t.Id == "" {
		t.Id = NewMongoId()
	}
	return m.adminTechnique(DBCmdInsert, t)
}

// Delete a new technique in DB.
func (m *MongoDbConn) DeleteTechnique(t *Technique) error {
	return m.adminTechnique(DBCmdDelete, t)
}

// Aux method that administers the technique records in DB
func (m *MongoDbConn) adminTechnique(cmd DbCommand, t *Technique) error {

	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("techniques")
	if coll == nil {
		return fmt.Errorf("Handling a technique: MongoDB descriptor empty.")
	}

	if t == nil {
		return fmt.Errorf("Handling a technique: cannot create empty technique.")
	}

	var err error
	switch cmd {

	case DBCmdUpdate:
		t.Modified = NewTimestamp() // update modified timestamp first...
		err = coll.UpdateId(t.Id, t)

	case DBCmdInsert:
		t.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(t)

	case DBCmdDelete:
		err = coll.RemoveId(t.Id)

	default:
		err = fmt.Errorf("Handling techniques: Unknown DB command.")
	}
	return err
}

///////////////////////////// Artists
func (m *MongoDbConn) GetAllArtists(t ArtistType) ([]*Artist, error) {

	// acquire DB lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Mongo descriptor empty.")
	}

	// create channel
	ch := make(chan []*Artist)

	// start a new goroutine to get users from DB
	go func(ch chan []*Artist) {

		// check channel
		if ch == nil {
			return
		}

		// prepare the empty slice for users
		artists := make([]*Artist, 0)
		var err error

		// get all artists from DB
		switch t {

		case ArtistTypeArtist: // get all case
			if err = db.C("artists").Find(bson.M{}).All(&artists); err != nil {
				return
			}

		case ArtistTypePainter:
			if err = db.C("artists").Find(bson.M{"is_painter": true}).All(&artists); err != nil {
				return
			}

		case ArtistTypeSculptor:
			if err = db.C("artists").Find(bson.M{"is_sculptor": true}).All(&artists); err != nil {
				return
			}

		case ArtistTypeArchitect:
			if err = db.C("artists").Find(bson.M{"is_architect": true}).All(&artists); err != nil {
				return
			}

		case ArtistTypePrintmaker:
			if err = db.C("artists").Find(bson.M{"is_printmaker": true}).All(&artists); err != nil {
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
func (m *MongoDbConn) GetArtist(id string) (*Artist, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("MongoDB descriptor empty.")
	}

	// prepare channel
	ch := make(chan *Artist)

	// start goroutine to get a user
	go func(id string, ch chan *Artist) {

		u := NewArtist() // create empty user

		// get a user from DB
		err := db.C("artists").Find(bson.M{"_id": MongoStringToId(id)}).One(&u)
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
func (m *MongoDbConn) UpdateArtist(a *Artist) error {
	return m.adminArtist(DBCmdUpdate, a)
}

// Create a new artist in DB.
func (m *MongoDbConn) InsertArtist(a *Artist) error {

	// check the ID of the item to be inserted into DB
	if a.Id == "" {
		a.Id = NewMongoId()
	}
	return m.adminArtist(DBCmdInsert, a)
}

// Delete a new artist in DB.
func (m *MongoDbConn) DeleteArtist(a *Artist) error {
	return m.adminArtist(DBCmdDelete, a)
}

// Aux method that administers the artist records in DB
func (m *MongoDbConn) adminArtist(cmd DbCommand, a *Artist) error {

	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("artists")
	if coll == nil {
		return fmt.Errorf("Handling an artist: MongoDB descriptor empty.")
	}

	if a == nil {
		return fmt.Errorf("Handling an artist: cannot create empty artist.")
	}

	var err error
	switch cmd {

	case DBCmdUpdate:
		a.Modified = NewTimestamp()
		err = coll.UpdateId(a.Id, a)

	case DBCmdInsert:
		a.Created = NewTimestamp()
		err = coll.Insert(a)

	case DBCmdDelete:
		err = coll.RemoveId(a.Id)

	default:
		err = fmt.Errorf("Handling artists: Unknown DB command.")
	}
	return err
}

////////////////////// Paintings

// GetAllPaintings retrieves all paintings from DB with given DB descriptor.
func (m *MongoDbConn) GetAllPaintings() ([]*Painting, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting paintings: DB descriptor empty.")
	}

	// create channel
	ch := make(chan []*Painting)

	// start a new goroutine to get users from DB
	go func(ch chan []*Painting) {

		// check channel
		if ch == nil {
			return
		}

		// prepare the empty slice for users
		p := make([]*Painting, 0)

		// get all users from DB
		if err := db.C("paintings").Find(bson.D{}).All(&p); err != nil {
			return
		}

		// write the users to the channel
		ch <- p

	}(ch)

	// read the answer from channel
	p := <-ch

	return p, nil // OK
}

// GetPainting retrieves a single painting from the DB: we need an ID.
func (m *MongoDbConn) GetPainting(id string) (*Painting, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting single painting: MongoDB descriptor empty.")
	}

	// prepare channel
	ch := make(chan *Painting)

	// start goroutine to get a user
	go func(id string, ch chan *Painting) {

		p := NewPainting() // create empty user

		// get a user from DB
		err := db.C("paintings").Find(bson.M{"_id": MongoStringToId(id)}).One(&p)
		if err != nil {
			return
		}

		// write a user to channel
		ch <- p
	}(id, ch)

	// read user from channel
	p := <-ch

	return p, nil // all OK
}

// UpdatePainting modifies a single painting in DB.
func (m *MongoDbConn) UpdatePainting(p *Painting) error {
	return m.adminPainting(DBCmdUpdate, p)
}

// InsertPainting creates a new painting in DB.
func (m *MongoDbConn) InsertPainting(p *Painting) error {
	// check the ID of the item to be inserted into DB
	if p.Id == "" {
		p.Id = NewMongoId()
	}
	return m.adminPainting(DBCmdInsert, p)
}

// DeletePainting removes a single painting from DB.
func (m *MongoDbConn) DeletePainting(p *Painting) error {
	return m.adminPainting(DBCmdDelete, p)
}

// Aux method that administers the painting records in DB
func (m *MongoDbConn) adminPainting(cmd DbCommand, p *Painting) error {

	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("paintings")
	if coll == nil {
		return fmt.Errorf("Handling a painting: MongoDB descriptor empty.")
	}

	if p == nil {
		return fmt.Errorf("Handling a painting: cannot create empty painting.")
	}

	var err error
	switch cmd {

	case DBCmdUpdate:
		p.Modified = NewTimestamp() // update modified timestamp first...
		err = coll.UpdateId(p.Id, p)

	case DBCmdInsert:
		p.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(p)

	case DBCmdDelete:
		err = coll.RemoveId(p.Id)

	default:
		err = fmt.Errorf("Handling paintings: Unknown DB command.")
	}
	return err
}

////////////////////// sculpture

// GetAllSculptures retrieves all paintings from DB with given DB descriptor.
func (m *MongoDbConn) GetAllSculptures() ([]*Sculpture, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting paintings: DB descriptor empty.")
	}

	// create channel
	ch := make(chan []*Sculpture)

	// start a new goroutine to get sculptures from DB
	go func(ch chan []*Sculpture) {

		// check channel
		if ch == nil {
			return
		}

		// prepare the empty slice for users
		p := make([]*Sculpture, 0)

		// get all users from DB
		if err := db.C("sculptures").Find(bson.D{}).All(&p); err != nil {
			return
		}

		// write the users to the channel
		ch <- p

	}(ch)

	// read the answer from channel
	p := <-ch

	return p, nil // OK
}

// GetSculpture retrieves a single sculpture from the DB: we need an ID.
func (m *MongoDbConn) GetSculpture(id string) (*Sculpture, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting single painting: MongoDB descriptor empty.")
	}

	// prepare channel
	ch := make(chan *Sculpture)

	// start goroutine to get a user
	go func(id string, ch chan *Sculpture) {

		p := NewSculpture() // create empty user

		// get a user from DB
		err := db.C("sculptures").Find(bson.M{"_id": MongoStringToId(id)}).One(&p)
		if err != nil {
			return
		}

		// write a user to channel
		ch <- p
	}(id, ch)

	// read user from channel
	p := <-ch

	return p, nil // all OK
}

// UpdateSculpture modifies a single sculpture in DB.
func (m *MongoDbConn) UpdateSculpture(p *Sculpture) error {
	return m.adminSculpture(DBCmdUpdate, p)
}

// InsertSculpture creates a new painting in DB.
func (m *MongoDbConn) InsertSculpture(p *Sculpture) error {
	// check the ID of the item to be inserted into DB
	if p.Id == "" {
		p.Id = NewMongoId()
	}
	return m.adminSculpture(DBCmdInsert, p)
}

// DeleteSculpture removes a single sculpture from DB.
func (m *MongoDbConn) DeleteSculpture(p *Sculpture) error {
	return m.adminSculpture(DBCmdDelete, p)
}

// Aux method that administers the painting records in DB
func (m *MongoDbConn) adminSculpture(cmd DbCommand, p *Sculpture) error {

	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("sculptures")
	if coll == nil {
		return fmt.Errorf("Handling a sculpture: MongoDB descriptor empty.")
	}

	if p == nil {
		return fmt.Errorf("Handling a sculpture: cannot create empty painting.")
	}

	var err error
	switch cmd {

	case DBCmdUpdate:
		p.Modified = NewTimestamp() // update modified timestamp first...
		err = coll.UpdateId(p.Id, p)

	case DBCmdInsert:
		p.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(p)

	case DBCmdDelete:
		err = coll.RemoveId(p.Id)

	default:
		err = fmt.Errorf("Handling sculptures: Unknown DB command.")
	}
	return err
}

////////////////////// Graphic prints

// GetAllPrints retrieves all prints from DB with given DB descriptor.
func (m *MongoDbConn) GetAllPrints() ([]*Print, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting graphic prints: DB descriptor empty.")
	}

	// create channel
	ch := make(chan []*Print)

	// start a new goroutine to get prints from DB
	go func(ch chan []*Print) {

		// check channel
		if ch == nil {
			return
		}

		// prepare the empty slice for users
		p := make([]*Print, 0)

		// get all users from DB
		if err := db.C("prints").Find(bson.D{}).All(&p); err != nil {
			return
		}

		// write the users to the channel
		ch <- p

	}(ch)

	// read the answer from channel
	p := <-ch

	return p, nil // OK
}

// GetPrint retrieves a single print from the DB: we need an ID.
func (m *MongoDbConn) GetPrint(id string) (*Print, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting single graphic print: MongoDB descriptor empty.")
	}

	// prepare channel
	ch := make(chan *Print)

	// start goroutine to get a user
	go func(id string, ch chan *Print) {

		p := NewPrint() // create empty print

		// get a user from DB
		err := db.C("prints").Find(bson.M{"_id": MongoStringToId(id)}).One(&p)
		if err != nil {
			return
		}

		// write a user to channel
		ch <- p
	}(id, ch)

	// read user from channel
	p := <-ch

	return p, nil // all OK
}

// UpdatePrint modifies a single print in DB.
func (m *MongoDbConn) UpdatePrint(p *Print) error {
	return m.adminPrint(DBCmdUpdate, p)
}

// Insertprint creates a new print in DB.
func (m *MongoDbConn) InsertPrint(p *Print) error {
	// check the ID of the item to be inserted into DB
	if p.Id == "" {
		p.Id = NewMongoId()
	}
	return m.adminPrint(DBCmdInsert, p)
}

// DeletePrint removes a single graphic print from DB.
func (m *MongoDbConn) DeletePrint(p *Print) error {
	return m.adminPrint(DBCmdDelete, p)
}

// Aux method that administers the graphic print records in DB
func (m *MongoDbConn) adminPrint(cmd DbCommand, p *Print) error {

	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("prints")
	if coll == nil {
		return fmt.Errorf("Handling a graphic print: MongoDB descriptor empty.")
	}

	if p == nil {
		return fmt.Errorf("Handling a graphic print: cannot create empty painting.")
	}

	var err error
	switch cmd {

	case DBCmdUpdate:
		p.Modified = NewTimestamp() // update modified timestamp first...
		err = coll.UpdateId(p.Id, p)

	case DBCmdInsert:
		p.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(p)

	case DBCmdDelete:
		err = coll.RemoveId(p.Id)

	default:
		err = fmt.Errorf("Handling a graphic print: Unknown DB command.")
	}
	return err
}

////////////////////// Buildings

// GetAllBuildings retrieves all buildings from DB with given DB descriptor.
func (m *MongoDbConn) GetAllBuildings() ([]*Building, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting buildings: DB descriptor empty.")
	}

	// create channel
	ch := make(chan []*Building)

	// start a new goroutine to get buildings from DB
	go func(ch chan []*Building) {

		// check channel
		if ch == nil {
			return
		}

		// prepare the empty slice for users
		p := make([]*Building, 0)

		// get all users from DB
		if err := db.C("buildings").Find(bson.D{}).All(&p); err != nil {
			return
		}

		// write the users to the channel
		ch <- p

	}(ch)

	// read the answer from channel
	p := <-ch

	return p, nil // OK
}

// GetBuilding retrieves a single buiding from the DB: we need an ID.
func (m *MongoDbConn) GetBuilding(id string) (*Building, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting a single building: MongoDB descriptor empty.")
	}

	// prepare channel
	ch := make(chan *Building)

	// start goroutine to get a building
	go func(id string, ch chan *Building) {

		p := NewBuilding()

		// get a user from DB
		err := db.C("buildings").Find(bson.M{"_id": MongoStringToId(id)}).One(&p)
		if err != nil {
			return
		}

		// write a user to channel
		ch <- p
	}(id, ch)

	// read user from channel
	p := <-ch

	return p, nil // all OK
}

// UpdateBuilding modifies a single in DB.
func (m *MongoDbConn) UpdateBuilding(b *Building) error {
	return m.adminBuilding(DBCmdUpdate, b)
}

// InsertBuilding creates a new building in DB.
func (m *MongoDbConn) InsertBuilding(b *Building) error {
	// check the ID of the item to be inserted into DB
	if b.Id == "" {
		b.Id = NewMongoId()
	}
	return m.adminBuilding(DBCmdInsert, b)
}

// DeleteBuilding removes a single building from DB.
func (m *MongoDbConn) DeleteBuilding(b *Building) error {
	return m.adminBuilding(DBCmdDelete, b)
}

// Aux method that administers the building records in DB
func (m *MongoDbConn) adminBuilding(cmd DbCommand, p *Building) error {

	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("buildings")
	if coll == nil {
		return fmt.Errorf("Handling a building: MongoDB descriptor empty.")
	}

	if p == nil {
		return fmt.Errorf("Handling a building: cannot create empty painting.")
	}

	var err error
	switch cmd {

	case DBCmdUpdate:
		p.Modified = NewTimestamp() // update modified timestamp first...
		err = coll.UpdateId(p.Id, p)

	case DBCmdInsert:
		p.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(p)

	case DBCmdDelete:
		err = coll.RemoveId(p.Id)

	default:
		err = fmt.Errorf("Handling a building: Unknown DB command.")
	}
	return err
}

//// Additional methods

//
func (m *MongoDbConn) GetDatingNames() ([]string, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	var err error
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, err
	}

	// get all dating names from DB
	d := make([]*Dating, 0)
	//if err := db.C("datings").Find(nil).Select(bson.M{"dating": 1}).All(&d); err != nil {
	if err := db.C("datings").Find(bson.D{}).All(&d); err != nil { // XXX: This is not OK, but I can't get it the other way...
		return nil, err
	}
	datings := make([]string, 0)
	for _, val := range d {
		datings = append(datings, val.Dating.Dating)
	}
	return datings, nil
}

func (m *MongoDbConn) GetStyleNames() ([]string, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	var err error
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, err
	}

	// get all style names from DB
	s := make([]*Style, 0)
	//if err := db.C("styles").Find(bson.D{}).Select(bson.M{ "name": 1, "_id":0 }).All(&s); err != nil {
	if err := db.C("styles").Find(bson.D{}).All(&s); err != nil { // XXX: This is not OK, but I can't get it the other way...
		return nil, err
	}
	styles := make([]string, 0)
	for _, val := range s {
		styles = append(styles, val.Style.Name)
	}
	return styles, nil
}

func (m *MongoDbConn) GetTechniqueNames() ([]string, error) {

	// acquire lock
	dblock.Lock()
	defer dblock.Unlock()

	// get *mgo.Database instance
	var err error
	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, err
	}

	// get all style names from DB
	t := make([]*Technique, 0)
	//if err := db.C("techniques").Find(bson.D{}).Select(bson.M{ "name": 1, "_id":0 }).All(&t); err != nil {
	if err := db.C("techniques").Find(bson.D{}).All(&t); err != nil { // XXX: This is not OK, but I can't get it the other way...
		return nil, err
	}
	techs := make([]string, 0)
	for _, val := range t {
		techs = append(techs, val.Technique.Name)
	}
	return techs, nil
}
