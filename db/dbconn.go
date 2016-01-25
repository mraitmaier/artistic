//
// dbconn.go
//
package db

import (
    "fmt"
    "time"
    "sync"
)


// DB access is guarded by simple mutex, it's easier than goroutines...
var dblock sync.Mutex

// let's define DB type enum
type DbType int
const (
    UnknownDB DbType = iota
    MongoDB
)


const (
    //
    DefAppUsername = "admin"
    //
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

// Interface that defines the database connection.
type DbConnector interface {

    Connect(url string, timeout time.Duration) error
    Close()
}

// Interface that defines the data provider
type DataProvider interface {

    GetAllUsers() ([]User, error)
    GetUser(string) (*User, error)
    GetUserByUsername(string) (*User, error)
    InsertUser(*User) error
    UpdateUser(*User) error
    DeleteUser(*User) error
    CountUsers() (int, error)

    GetAllArtists(ArtistType) ([]Artist, error)
    GetArtist(string) (*Artist, error)
    InsertArtist(*Artist) error
    UpdateArtist(*Artist) error
    DeleteArtist(*Artist) error

    GetAllTechniques() ([]Technique, error)
    GetTechnique(string) (*Technique, error)
    InsertTechnique(*Technique) error
    UpdateTechnique(*Technique) error
    DeleteTechnique(*Technique) error

    GetAllStyles() ([]Style, error)
    GetStyle(string) (*Style, error)
    InsertStyle(*Style) error
    UpdateStyle(*Style) error
    DeleteStyle(*Style) error

    GetAllDatings() ([]*Dating, error)
    GetDating(string) (*Dating, error)
    UpdateDating(*Dating) error
    CountDatings() (int, error)
    InsertDatings([]*Dating) error
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

///////////////////////////// EXPERIMENTAL ///////////////////
// NOTE: how to abstract away the DB ID (for differents DBs)? By implementing
// a type that satisfies the interface...
//
type DbIdentifier interface {

    IdToString() string

    StringToId(string)

    NewId()
}
