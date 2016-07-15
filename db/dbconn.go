//
// dbconn.go
//
package db

import (
	"fmt"
	"sync"
	"time"
)

// DB access is guarded by simple mutex, it's easier than goroutines...
var dblock sync.Mutex

// let's define DB type enum
type DbType int

const (

	// UnknownDB represent unknown database server
	UnknownDB DbType = 1 << iota

	// MongoDB represents the MongoDB server environment
	MongoDB
)

const (

	// DefAppUsername contains the default application username
	DefAppUsername = "admin"

	// DefAppPasswd contains the default application password
	DefAppPasswd = "admin123!"
)

/*
// A factory function for the right DB IDs: returns the proper DB ID for the DB type given.
func NewDbId(dbtype DbType) DbIdentifier {

    var id DbIdentifier = nil

    switch dbtype {

    case MongoDB: // if DB is MongoDB
        var id MongoDbId
        id.NewId()

    default:     // other cases...
    }

    return id
}
*/

// DbConnector represents the open-close database interface (database connection).
type DbConnector interface {
	Connect(url string, timeout time.Duration) error
	Close()
}

// Counter represents the count-items database interface.
type Counter interface {
	Count() (int, error)
}

// Inserter represents the insert database interface
type Inserter interface {
	Insert() error
}

// Updater represents the update database interface.
type Updater interface {
	Update() error
}

// Remover represents the remove database interface. Several items can be removed at once.
type Remover interface {
	Remove() error
}

// Getter represents the get DB interface.
type Getter interface {
	All() ([]interface{}, error)
	One() (interface{}, error)
}

// UserGetter add an additional behavior to Getter interface: users are usually retrieved by username, not by ID.
type UserGetter interface {
	Getter
	ByUsername(uname string) (interface{}, error)
}

// NOTE: how to abstract away the DB ID (for differents DBs)? By implementing
// a type that satisfies the interface...
//
type DBIdentifier interface {
	IdToString() string
	StringToId(string)
}

// Interface that defines the data provider
type DataProvider interface {
	GetUsers(string) ([]*User, error)
	GetUser(string) (*User, error)
	GetUserByUsername(string) (*User, error)
	InsertUser(*User) error
	UpdateUser(*User) error
	DeleteUser(*User) error
	CountUsers() (int, error)

	GetArtists(ArtistType, string) ([]*Artist, error)
	GetArtist(string) (*Artist, error)
	InsertArtist(*Artist) error
	UpdateArtist(*Artist) error
	DeleteArtist(*Artist) error

	GetTechniques(string) ([]*Technique, error)
	GetTechnique(string) (*Technique, error)
	InsertTechnique(*Technique) error
	UpdateTechnique(*Technique) error
	DeleteTechnique(*Technique) error

	GetStyles(string) ([]*Style, error)
	GetStyle(string) (*Style, error)
	InsertStyle(*Style) error
	UpdateStyle(*Style) error
	DeleteStyle(*Style) error

	GetDatings(string) ([]*Dating, error)
	GetDating(string) (*Dating, error)
	UpdateDating(*Dating) error
	CountDatings() (int, error)
	InsertDatings([]*Dating) error

	GetPaintings(string) ([]*Painting, error)
	GetPainting(string) (*Painting, error)
	InsertPainting(*Painting) error
	UpdatePainting(*Painting) error
	DeletePainting(*Painting) error

	GetSculptures(string) ([]*Sculpture, error)
	GetSculpture(string) (*Sculpture, error)
	InsertSculpture(*Sculpture) error
	UpdateSculpture(*Sculpture) error
	DeleteSculpture(*Sculpture) error

	GetPrints(string) ([]*Print, error)
	GetPrint(string) (*Print, error)
	InsertPrint(*Print) error
	UpdatePrint(*Print) error
	DeletePrint(*Print) error

	GetBuildings(string) ([]*Building, error)
	GetBuilding(string) (*Building, error)
	InsertBuilding(*Building) error
	UpdateBuilding(*Building) error
	DeleteBuilding(*Building) error

	GetBooks(string) ([]*Book, error)
	GetBook(string) (*Book, error)
	InsertBook(*Book) error
	UpdateBook(*Book) error
	DeleteBook(*Book) error

	GetArticles(string) ([]*Article, error)
	GetArticle(string) (*Article, error)
	InsertArticle(*Article) error
	UpdateArticle(*Article) error
	DeleteArticle(*Article) error

	// helper methods
	GetDatingNames() ([]string, error)
	GetStyleNames() ([]string, error)
	GetTechniqueNames() ([]string, error)
}

// Init database factory: returns appropriate DB connection URL, initialized
// DB connector and initialized DB data provider, according to DB type given.
// Note that DB connector and data provider are, in general, implemented
// by the same type.
func InitDb(dbtype DbType, host string, port int,
	username, password, dbname string) (url string, db DbConnector, data DataProvider, e error) {

	// initialize
	url = ""
	db = nil
	data = nil

	switch dbtype {

	case MongoDB:
		url, db, data = initMongo(host, port, username, password, dbname)
		return url, db, data, e

	}
	return url, db, data, fmt.Errorf("Unknown database: cannot connect.\n")
}

// Initialization function for MongoDB.
func initMongo(host string, port int, username, passwd, dbname string) (string, *MongoDbConn, *MongoDbConn) {

	// create DB connection URL for MongoDB
	s := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", username, passwd, host, port, dbname)

	// create new instance of MongoDB Session
	db := new(MongoDbConn)
	db.name = dbname

	return s, db, db
}

///////////////
type DbCommand int

const (
	DBCmdGetOne DbCommand = iota // get a single record from DB
	DBCmdUpdate                  // update a single record in DB
	DBCmdInsert                  // insert a new record in DB
	DBCmdDelete                  // delete a single record in DB
)

