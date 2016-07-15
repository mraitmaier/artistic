package db

//
// mongo.go -
//

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

	if m.Sess, e = mgo.DialWithTimeout(url, timeout); e != nil {
        return
    }
	return ensureIndexes(m)
}

// Close the mongoDB connection.
func (m *MongoDbConn) Close() {
	if m.Sess != nil {
		m.Sess.Close()
	}
}

// The ensureIndexes function ensures that MongoDB indexes are created, when app is started.
func ensureIndexes(m *MongoDbConn) error {

	// wildcard text index is common to all collection
	//wcix := mgo.Index{Key: []string{"$text:$**"}, Background: true, Sparse: true}

	var err error
	// the artists collection indexes
	c := m.Sess.DB(m.name).C("artists")
	err = c.EnsureIndex(mgo.Index{Key: []string{"nationality"}, Background: true, Sparse: true})
	//err = c.EnsureIndex(wcix)

	// the techniques collection indexes
	c = m.Sess.DB(m.name).C("techniques")
	err = c.EnsureIndex(mgo.Index{Key: []string{"techniquetype"}, Background: true, Sparse: true})
    //err = c.EnsureIndex(wcix)

	// the articles collection indexes
	c = m.Sess.DB(m.name).C("articles")
	err = c.EnsureIndex(mgo.Index{Key: []string{"year"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"publisher"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"authors"}, Background: true, Sparse: true})
	//err = c.EnsureIndex(wcix)

	// the books collection indexes
	c = m.Sess.DB(m.name).C("books")
	err = c.EnsureIndex(mgo.Index{Key: []string{"year"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"publication"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"authors"}, Background: true, Sparse: true})
	//err = c.EnsureIndex(wcix)

	// the paintings collection indexes
	c = m.Sess.DB(m.name).C("paintings")
	err = c.EnsureIndex(mgo.Index{Key: []string{"artist"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"technique"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"style"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"dating"}, Background: true, Sparse: true})
	//err = c.EnsureIndex(wcix)

	// the sculptures collection indexes
	c = m.Sess.DB(m.name).C("sculptures")
	err = c.EnsureIndex(mgo.Index{Key: []string{"artist"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"technique"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"style"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"dating"}, Background: true, Sparse: true})
	//err = c.EnsureIndex(wcix)

	// the prints collection indexes
	c = m.Sess.DB(m.name).C("prints")
	err = c.EnsureIndex(mgo.Index{Key: []string{"artist"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"technique"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"style"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"dating"}, Background: true, Sparse: true})
	//err = c.EnsureIndex(wcix)

	// the buildings collection indexes
	c = m.Sess.DB(m.name).C("buildings")
	err = c.EnsureIndex(mgo.Index{Key: []string{"artist"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"technique"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"style"}, Background: true, Sparse: true})
	err = c.EnsureIndex(mgo.Index{Key: []string{"dating"}, Background: true, Sparse: true})
	//err = c.EnsureIndex(wcix)

	return err
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

// Retrieves users from database.
func (m *MongoDbConn) GetUsers(qry string) ([]*User, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Mongo descriptor empty.")
	}

	ch := make(chan []*User)
	go func(ch chan []*User, qry string) {

		if ch == nil {
			return
		}

		users := make([]*User, 0)
		if qry == "" {
			_ = db.C("users").Find(bson.D{}).All(&users) // if qry is empty, retrieve all...
		} else {
			//qry := m.decodeQuery(srch)
			// XXX currently hard-coded, instead of decodeQuery()...
			//qry := bson.M{ "$text": bson.M{ "$search": srch } }
			_ = db.C("users").Find(bson.M{"$text": bson.M{"$search": qry}}).Sort("username").All(&users)
		}

		ch <- users
	}(ch, qry)

	users := <-ch
	return users, nil
}

// Get a single user from the DB: we need a username.
func (m *MongoDbConn) GetUserByUsername(username string) (*User, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("MongoDB descriptor empty.")
	}

	ch := make(chan *User)
	go func(username string, ch chan *User) {

		u := NewUser()
		err := db.C("users").Find(bson.M{"username": username}).One(&u)
		if err != nil {
			return
		}
		ch <- u

	}(username, ch)

	user := <-ch
	return user, nil
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

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("MongoDB descriptor empty.")
	}

	// get a user from DB
	u := NewUser()
	if err := db.C("users").Find(bson.M{"_id": MongoStringToId(id)}).One(&u); err != nil {
		return nil, err
	}

	return u, nil // all OK
}

// Update a single user in DB.
func (m *MongoDbConn) UpdateUser(u *User) error { return m.adminUser(DBCmdUpdate, u) }

// Create a new user in DB.
func (m *MongoDbConn) InsertUser(u *User) error { return m.adminUser(DBCmdInsert, u) }

// Delete a single user in DB.
func (m *MongoDbConn) DeleteUser(u *User) error { return m.adminUser(DBCmdDelete, u) }

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
		err = coll.UpdateId(u.ID, u)

	case DBCmdInsert:
		u.Created = NewTimestamp()
		err = coll.Insert(u)

	case DBCmdDelete:
		err = coll.RemoveId(u.ID)

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

// Retrieves datings from database.
func (m *MongoDbConn) GetDatings(srch string) ([]*Dating, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting datings: MongoDb descriptor empty.")
	}

	var err error
	d := make([]*Dating, 0)
	if srch == "" {
		err = db.C("datings").Find(bson.D{}).All(&d) // if qry is empty, retrieve all...
	} else {
		//qry := m.decodeQuery(srch)
		// XXX currently hard-coded, instead of decodeQuery()...
		//qry := bson.M{ "$text": bson.M{ "$search": srch } }
		err = db.C("datings").Find(bson.M{"$text": bson.M{"$search": srch}}).All(&d)
	}
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Retrieve a single Dating record by ID.
func (m *MongoDbConn) GetDating(id string) (*Dating, error) {

	dblock.Lock()
	defer dblock.Unlock()

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
	if err := db.C("datings").Update(bson.M{"_id": d.ID}, d); err != nil {
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

// Retrieve a single Style record.
func (m *MongoDbConn) GetStyles(srch string) ([]*Style, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Update a style: MongoDb descriptor empty.")
	}

	var err error
	s := make([]*Style, 0)
	if srch == "" {
		err = db.C("styles").Find(bson.D{}).All(&s)
	} else {
		//qry := m.decodeQuery(srch)
		// XXX currently hard-coded, instead of decodeQuery()...
		//qry := bson.M{ "$text": bson.M{ "$search": srch } }
		err = db.C("styles").Find(bson.M{"$text": bson.M{"$search": srch}}).Sort("name").All(&s)
	}
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Retrieve a single Style record.
func (m *MongoDbConn) GetStyle(id string) (*Style, error) {

	dblock.Lock()
	defer dblock.Unlock()

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
func (m *MongoDbConn) UpdateStyle(s *Style) error { return m.adminStyle(DBCmdUpdate, s) }

// Create a new style in DB.
func (m *MongoDbConn) InsertStyle(s *Style) error { return m.adminStyle(DBCmdInsert, s) }

// Delete a single style in DB
func (m *MongoDbConn) DeleteStyle(s *Style) error { return m.adminStyle(DBCmdDelete, s) }

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
		err = coll.UpdateId(s.ID, s)

	case DBCmdInsert:
		s.Created = NewTimestamp()
		err = coll.Insert(s)

	case DBCmdDelete:
		err = coll.RemoveId(s.ID)

	default:
		err = fmt.Errorf("Handling styles: Unknown command.")
	}

	return err
}

///////////////////////////// Techniques

// Retrieve a single Style record.
func (m *MongoDbConn) GetTechniques(srch string) ([]*Technique, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Update a technique: MongoDb descriptor empty.")
	}

	var err error
	t := make([]*Technique, 0)
	if srch == "" {
		err = db.C("techniques").Find(bson.D{}).All(&t)
	} else {
		err = db.C("techniques").Find(bson.M{"$text": bson.M{"$search": srch}}).Sort("name").All(&t)
	}
	if err != nil {
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
func (m *MongoDbConn) UpdateTechnique(t *Technique) error { return m.adminTechnique(DBCmdUpdate, t) }

// Create a new technique in DB.
func (m *MongoDbConn) InsertTechnique(t *Technique) error { return m.adminTechnique(DBCmdInsert, t) }

// Delete a new technique in DB.
func (m *MongoDbConn) DeleteTechnique(t *Technique) error { return m.adminTechnique(DBCmdDelete, t) }

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
		err = coll.UpdateId(t.ID, t)

	case DBCmdInsert:
		t.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(t)

	case DBCmdDelete:
		err = coll.RemoveId(t.ID)

	default:
		err = fmt.Errorf("Handling techniques: Unknown DB command.")
	}
	return err
}

///////////////////////////// Artists
func (m *MongoDbConn) GetArtists(t ArtistType, srch string) ([]*Artist, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Mongo descriptor empty.")
	}

	ch := make(chan []*Artist)
	go func(ch chan []*Artist, qry string) {

		// check channel
		if ch == nil {
			return
		}

		artists := make([]*Artist, 0)

		switch t {

		case ArtistTypeArtist: // get all case
			if qry == "" {
				_ = db.C("artists").Find(bson.D{}).All(&artists)
			} else {
				_ = db.C("artists").Find(bson.M{"$text": bson.M{"$search": qry}}).Sort("name").All(&artists)
			}

		case ArtistTypePainter:
			if qry == "" {
				_ = db.C("artists").Find(bson.M{"is_painter": true}).Sort("name").All(&artists)
			} else {
				_ = db.C("artists").Find(
					bson.M{"is_painter": true, "$text": bson.M{"$search": qry}}).Sort("name").All(&artists)
			}

		case ArtistTypeSculptor:
			if qry == "" {
				_ = db.C("artists").Find(bson.M{"is_sculptor": true}).All(&artists)
			} else {
				_ = db.C("artists").Find(
					bson.M{"is_sculptor": true, "$text": bson.M{"$search": qry}}).Sort("name").All(&artists)
			}

		case ArtistTypeArchitect:
			if qry == "" {
				_ = db.C("artists").Find(bson.M{"is_architect": true}).All(&artists)
			} else {
				_ = db.C("artists").Find(
					bson.M{"is_architect": true, "$text": bson.M{"$search": qry}}).Sort("name").All(&artists)
			}

		case ArtistTypePrintmaker:
			if qry == "" {
				_ = db.C("artists").Find(bson.M{"is_printmaker": true}).All(&artists)
			} else {
				_ = db.C("artists").Find(
					bson.M{"is_printmaker": true, "$text": bson.M{"$search": qry}}).Sort("name").All(&artists)
			}
		}

		ch <- artists
	}(ch, srch)

	artists := <-ch
	return artists, nil
}

// Get a single artist from the DB: we need an ID .
func (m *MongoDbConn) GetArtist(id string) (*Artist, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("MongoDB descriptor empty.")
	}

	// start goroutine to get a user
	ch := make(chan *Artist)
	go func(id string, ch chan *Artist) {

		u := NewArtist()
		err := db.C("artists").Find(bson.M{"_id": MongoStringToId(id)}).One(&u)
		if err != nil {
			return
		}
		ch <- u

	}(id, ch)

	user := <-ch
	return user, nil
}

// Update a single artist in DB.
func (m *MongoDbConn) UpdateArtist(a *Artist) error { return m.adminArtist(DBCmdUpdate, a) }

// Create a new artist in DB.
func (m *MongoDbConn) InsertArtist(a *Artist) error { return m.adminArtist(DBCmdInsert, a) }

// Delete a new artist in DB.
func (m *MongoDbConn) DeleteArtist(a *Artist) error { return m.adminArtist(DBCmdDelete, a) }

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
		err = coll.UpdateId(a.ID, a)

	case DBCmdInsert:
		a.Created = NewTimestamp()
		err = coll.Insert(a)

	case DBCmdDelete:
		err = coll.RemoveId(a.ID)

	default:
		err = fmt.Errorf("Handling artists: Unknown DB command.")
	}
	return err
}

////////////////////// Paintings

// GetAllPaintings retrieves all paintings from DB with given DB descriptor.
func (m *MongoDbConn) GetPaintings(srch string) ([]*Painting, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting paintings: DB descriptor empty.")
	}

	// start a new goroutine to get paintings from DB
	ch := make(chan []*Painting)
	go func(ch chan []*Painting, qry string) {

		// check channel
		if ch == nil {
			return
		}

		// get all users from DB
		p := make([]*Painting, 0)
		if qry == "" {
			_ = db.C("paintings").Find(bson.D{}).All(&p)
		} else {
			_ = db.C("paintings").Find(bson.M{"$text": bson.M{"$search": qry}}).Sort("work.title").All(&p)
		}
		ch <- p
	}(ch, srch)

	p := <-ch
	return p, nil // OK
}

// GetPainting retrieves a single painting from the DB: we need an ID.
func (m *MongoDbConn) GetPainting(id string) (*Painting, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting single painting: MongoDB descriptor empty.")
	}

	// start goroutine to get a painting
	ch := make(chan *Painting)
	go func(id string, ch chan *Painting) {

		p := NewPainting() // create empty user
		err := db.C("paintings").Find(bson.M{"_id": MongoStringToId(id)}).One(&p)
		if err != nil {
			return
		}

		ch <- p
	}(id, ch)

	p := <-ch
	return p, nil // all OK
}

// UpdatePainting modifies a single painting in DB.
func (m *MongoDbConn) UpdatePainting(p *Painting) error { return m.adminPainting(DBCmdUpdate, p) }

// InsertPainting creates a new painting in DB.
func (m *MongoDbConn) InsertPainting(p *Painting) error { return m.adminPainting(DBCmdInsert, p) }

// DeletePainting removes a single painting from DB.
func (m *MongoDbConn) DeletePainting(p *Painting) error { return m.adminPainting(DBCmdDelete, p) }

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
		err = coll.UpdateId(p.ID, p)

	case DBCmdInsert:
		p.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(p)

	case DBCmdDelete:
		err = coll.RemoveId(p.ID)

	default:
		err = fmt.Errorf("Handling paintings: Unknown DB command.")
	}
	return err
}

////////////////////// sculpture

// GetSculptures retrieves all paintings from DB with given DB descriptor.
func (m *MongoDbConn) GetSculptures(srch string) ([]*Sculpture, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting paintings: DB descriptor empty.")
	}

	// start a new goroutine to get sculptures from DB
	ch := make(chan []*Sculpture)
	go func(ch chan []*Sculpture, qry string) {

		if ch == nil {
			return
		}

		p := make([]*Sculpture, 0)
		if qry == "" {
			_ = db.C("sculptures").Find(bson.D{}).All(&p)
		} else {
			_ = db.C("sculptures").Find(bson.M{"$text": bson.M{"$search": qry}}).Sort("work.title").All(&p)
		}
		ch <- p

	}(ch, srch)

	p := <-ch
	return p, nil // OK
}

// GetSculpture retrieves a single sculpture from the DB: we need an ID.
func (m *MongoDbConn) GetSculpture(id string) (*Sculpture, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting single painting: MongoDB descriptor empty.")
	}

	// start goroutine to get a user
	ch := make(chan *Sculpture)
	go func(id string, ch chan *Sculpture) {

		p := NewSculpture() // create empty user

		err := db.C("sculptures").Find(bson.M{"_id": MongoStringToId(id)}).One(&p)
		if err != nil {
			return
		}

		ch <- p
	}(id, ch)

	p := <-ch
	return p, nil // all OK
}

// UpdateSculpture modifies a single sculpture in DB.
func (m *MongoDbConn) UpdateSculpture(p *Sculpture) error { return m.adminSculpture(DBCmdUpdate, p) }

// InsertSculpture creates a new painting in DB.
func (m *MongoDbConn) InsertSculpture(p *Sculpture) error { return m.adminSculpture(DBCmdInsert, p) }

// DeleteSculpture removes a single sculpture from DB.
func (m *MongoDbConn) DeleteSculpture(p *Sculpture) error { return m.adminSculpture(DBCmdDelete, p) }

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
		err = coll.UpdateId(p.ID, p)

	case DBCmdInsert:
		p.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(p)

	case DBCmdDelete:
		err = coll.RemoveId(p.ID)

	default:
		err = fmt.Errorf("Handling sculptures: Unknown DB command.")
	}
	return err
}

////////////////////// Graphic prints

// GetAllPrints retrieves all prints from DB with given DB descriptor.
func (m *MongoDbConn) GetPrints(srch string) ([]*Print, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting graphic prints: DB descriptor empty.")
	}

	// start a new goroutine to get prints from DB
	ch := make(chan []*Print)
	go func(ch chan []*Print, qry string) {

		// check channel
		if ch == nil {
			return
		}

		p := make([]*Print, 0)
		if qry == "" {
			_ = db.C("prints").Find(bson.D{}).All(&p)
		} else {
			_ = db.C("prints").Find(bson.M{"$text": bson.M{"$search": qry}}).Sort("work.title").All(&p)
		}
		ch <- p

	}(ch, srch)

	p := <-ch
	return p, nil // OK
}

// GetPrint retrieves a single print from the DB: we need an ID.
func (m *MongoDbConn) GetPrint(id string) (*Print, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting single graphic print: MongoDB descriptor empty.")
	}

	// start goroutine to get a user
	ch := make(chan *Print)
	go func(id string, ch chan *Print) {

		p := NewPrint() // create empty print

		err := db.C("prints").Find(bson.M{"_id": MongoStringToId(id)}).One(&p)
		if err != nil {
			return
		}

		ch <- p
	}(id, ch)

	p := <-ch
	return p, nil // all OK
}

// UpdatePrint modifies a single print in DB.
func (m *MongoDbConn) UpdatePrint(p *Print) error { return m.adminPrint(DBCmdUpdate, p) }

// Insertprint creates a new print in DB.
func (m *MongoDbConn) InsertPrint(p *Print) error { return m.adminPrint(DBCmdInsert, p) }

// DeletePrint removes a single graphic print from DB.
func (m *MongoDbConn) DeletePrint(p *Print) error { return m.adminPrint(DBCmdDelete, p) }

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
		err = coll.UpdateId(p.ID, p)

	case DBCmdInsert:
		p.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(p)

	case DBCmdDelete:
		err = coll.RemoveId(p.ID)

	default:
		err = fmt.Errorf("Handling a graphic print: Unknown DB command.")
	}
	return err
}

////////////////////// Buildings

// GetAllBuildings retrieves all buildings from DB with given DB descriptor.
func (m *MongoDbConn) GetBuildings(srch string) ([]*Building, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting buildings: DB descriptor empty.")
	}

	// start a new goroutine to get buildings from DB
	ch := make(chan []*Building)
	go func(ch chan []*Building, qry string) {

		// check channel
		if ch == nil {
			return
		}

		p := make([]*Building, 0)
		if qry == "" {
			_ = db.C("buildings").Find(bson.D{}).All(&p) // get all
		} else {
			_ = db.C("buildings").Find(bson.M{"$text": bson.M{"$search": qry}}).Sort("work.title").All(&p)
		}
		ch <- p

	}(ch, srch)

	p := <-ch
	return p, nil // OK
}

// GetBuilding retrieves a single buiding from the DB: we need an ID.
func (m *MongoDbConn) GetBuilding(id string) (*Building, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting a single building: MongoDB descriptor empty.")
	}

	// start goroutine to get a building
	ch := make(chan *Building)
	go func(id string, ch chan *Building) {

		p := NewBuilding()

		// get a user from DB
		err := db.C("buildings").Find(bson.M{"_id": MongoStringToId(id)}).One(&p)
		if err != nil {
			return
		}

		ch <- p
	}(id, ch)

	// read user from channel
	p := <-ch
	return p, nil // all OK
}

// UpdateBuilding modifies a single in DB.
func (m *MongoDbConn) UpdateBuilding(b *Building) error { return m.adminBuilding(DBCmdUpdate, b) }

// InsertBuilding creates a new building in DB.
func (m *MongoDbConn) InsertBuilding(b *Building) error { return m.adminBuilding(DBCmdInsert, b) }

// DeleteBuilding removes a single building from DB.
func (m *MongoDbConn) DeleteBuilding(b *Building) error { return m.adminBuilding(DBCmdDelete, b) }

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
		err = coll.UpdateId(p.ID, p)

	case DBCmdInsert:
		p.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(p)

	case DBCmdDelete:
		err = coll.RemoveId(p.ID)

	default:
		err = fmt.Errorf("Handling a building: Unknown DB command.")
	}
	return err
}

////////////////////// Books

// GetBooks retrieves books from DB: either all or filtered (defined by search query).
func (m *MongoDbConn) GetBooks(srch string) ([]*Book, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting books: DB descriptor empty.")
	}

	ch := make(chan []*Book)
	go func(ch chan []*Book, qry string) {

		if ch == nil {
			return
		}

		b := make([]*Book, 0)
		if qry == "" {
			_ = db.C("books").Find(bson.D{}).All(&b) // get all
		} else {
			_ = db.C("books").Find(bson.M{"$text": bson.M{"$search": qry}}).Sort("title").All(&b)
		}
		ch <- b

	}(ch, srch)

	b := <-ch
	return b, nil
}

// GetBook retrieves a single book from the DB.
func (m *MongoDbConn) GetBook(id string) (*Book, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting a single book: MongoDB descriptor empty.")
	}

	ch := make(chan *Book)
	go func(id string, ch chan *Book) {

		b := NewBook()
		err := db.C("books").Find(bson.M{"_id": MongoStringToId(id)}).One(&b)
		if err != nil {
			return
		}
		ch <- b

	}(id, ch)

	b := <-ch
	return b, nil
}

// UpdateBook modifies a single book in DB.
func (m *MongoDbConn) UpdateBook(b *Book) error { return m.adminBook(DBCmdUpdate, b) }

// InsertBook creates a new book in DB.
func (m *MongoDbConn) InsertBook(b *Book) error { return m.adminBook(DBCmdInsert, b) }

// DeleteBook removes a single book from DB.
func (m *MongoDbConn) DeleteBook(b *Book) error { return m.adminBook(DBCmdDelete, b) }

// Aux method that administers the book records in DB
func (m *MongoDbConn) adminBook(cmd DbCommand, p *Book) error {

	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("books")
	if coll == nil {
		return fmt.Errorf("Handling a book: MongoDB descriptor empty")
	}

	if p == nil {
		return fmt.Errorf("Handling a book: cannot create empty book")
	}

	var err error
	switch cmd {

	case DBCmdUpdate:
		p.Modified = NewTimestamp() // update modified timestamp first...
		err = coll.UpdateId(p.ID, p)

	case DBCmdInsert:
		p.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(p)

	case DBCmdDelete:
		err = coll.RemoveId(p.ID)

	default:
		err = fmt.Errorf("Handling a book: Unknown DB command.")
	}
	return err
}

////////////////////// Articles

// GetArticles retrieves articles from DB with given DB descriptor.
func (m *MongoDbConn) GetArticles(srch string) ([]*Article, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting articles: DB descriptor empty.")
	}

	ch := make(chan []*Article)
	go func(ch chan []*Article, qry string) {

		if ch == nil {
			return
		}

		a := make([]*Article, 0)
		if qry == "" {
			_ = db.C("articles").Find(bson.D{}).All(&a) // get all
		} else {
			_ = db.C("articles").Find(bson.M{"$text": bson.M{"$search": qry}}).Sort("title").All(&a)
		}
		ch <- a

	}(ch, srch)

	a := <-ch
	return a, nil
}

// GetArticle retrieves a single item from the DB: we need an ID.
func (m *MongoDbConn) GetArticle(id string) (*Article, error) {

	dblock.Lock()
	defer dblock.Unlock()

	db := m.Sess.DB(m.name)
	if db == nil {
		return nil, errors.New("Getting a single article: MongoDB descriptor empty.")
	}

	ch := make(chan *Article)
	go func(id string, ch chan *Article) {

		a := NewArticle()
		err := db.C("articles").Find(bson.M{"_id": MongoStringToId(id)}).One(&a)
		if err != nil {
			return
		}
		ch <- a

	}(id, ch)

	a := <-ch
	return a, nil
}

// UpdateArticle modifies a single article in DB.
func (m *MongoDbConn) UpdateArticle(b *Article) error { return m.adminArticle(DBCmdUpdate, b) }

// InsertArticle creates a new article in DB.
func (m *MongoDbConn) InsertArticle(b *Article) error { return m.adminArticle(DBCmdInsert, b) }

// DeleteArticle removes a single article from DB.
func (m *MongoDbConn) DeleteArticle(b *Article) error { return m.adminArticle(DBCmdDelete, b) }

// Aux method that administers the article records in DB
func (m *MongoDbConn) adminArticle(cmd DbCommand, p *Article) error {

	dblock.Lock()
	defer dblock.Unlock()

	coll := m.Sess.DB(m.name).C("articles")
	if coll == nil {
		return fmt.Errorf("Handling an article: MongoDB descriptor empty")
	}

	if p == nil {
		return fmt.Errorf("Handling an article: cannot create empty article")
	}

	var err error
	switch cmd {

	case DBCmdUpdate:
		p.Modified = NewTimestamp() // update modified timestamp first...
		err = coll.UpdateId(p.ID, p)

	case DBCmdInsert:
		p.Created = NewTimestamp() // create new timestamp first...
		err = coll.Insert(p)

	case DBCmdDelete:
		err = coll.RemoveId(p.ID)

	default:
		err = fmt.Errorf("Handling an article: Unknown DB command.")
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

// The decodeQuery method is a helper function that takes a search string (as input by user in web browser) and
// returns a MongoDB (JSON-formatted) query string that is to be used directly to retrieve data.
func (m *MongoDbConn) decodeQuery(srch string) string {
	return ""
}
