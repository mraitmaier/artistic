package db

//
// dbconn.go
//

import (
	"fmt"
	"sync"
	"time"
)

// DB access is guarded by simple mutex, it's easier than goroutines...
var dblock sync.Mutex

// DBType defines list (enumeration) of supported databases. 
type DBType int

const (
	// UnknownDB represent unknown database server
	UnknownDB DBType = 1 << iota
	// MongoDB represents the MongoDB server environment
	MongoDB
	// SQLite represents the SQLite environment
	SQLite
)

const (
	// DefAppUsername contains the default application username
	DefAppUsername = "admin"
	// DefAppPasswd contains the default application password
	DefAppPasswd = "admin123!"
)

/*
// A factory function for the right DB IDs: returns the proper DB ID for the DB type given.
func NewDbId(dbtype DBType) DbIdentifier {

    var id DbIdentifier = nil

    switch dbtype {

    case MongoDB: // if DB is MongoDB
        var id MongoDbId
        id.NewId()

    case SQLite:
        // TODO

    default:     // other cases...
    }

    return id
}
*/

// DBConnector represents the open-close database interface (database connection).
type DBConnector interface {
	Connect(url string, timeout time.Duration) error
	Close()
}

// DataType is an enumeration of possible data types in database. This is basically a helper type
// to easier create new concrete instances of above defined interafaces.
type DataType int

const (
	ErrorDataType = iota
	ArtistDataType
	DatingDataType
	StyleDataType
	TechniqueDataType
	PaintingDataType
	ScultptureDataType
	PrintDataType
	BuildingDataType
	CollectionDataType
	InstitutionDataType
	ExhibitionDataType
	BookDataType
	ArticleDataType
	UserDataType
)

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
func InitDb(dbtype DBType, host string, port int,
	username, password, dbname string) (url string, db DBConnector, data DataProvider, e error) {

	// initialize
	url = ""
	db = nil
	data = nil

	switch dbtype {

	case MongoDB:
		url, db, data = initMongo(host, port, username, password, dbname)
		return url, db, data, e

	case SQLite:
		// TODO

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
type DBCommand int

const (
	DBCmdGetOne DBCommand = iota // get a single record from DB
	DBCmdUpdate                  // update a single record in DB
	DBCmdInsert                  // insert a new record in DB
	DBCmdDelete                  // delete a single record in DB
)
