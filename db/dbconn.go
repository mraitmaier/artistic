//
// dbconn.go
//
package db

import (
    "fmt"
    "time"
    "sync"
    "bitbucket.org/miranr/artistic/utils"
    "bitbucket.org/miranr/artistic/core"
)


// DB access is guarded by simple mutex, it's easier than goroutines...
var dblock sync.Mutex

// let's define DB type enum
type DbType int
const (
    UnknownDB DbType = iota
    MongoDB
)

// Interface that defines the database connection.
type DbConnector interface {

    Connect(url string, timeout time.Duration) error
    Close()
}

// Interface that defines the data provider
type DataProvider interface {

    GetAllUsers() ([]utils.User, error)
    GetUser(string) (*utils.User, error)
    GetUserByUsername(string) (*utils.User, error)
    InsertUser(*utils.User) error
    UpdateUser(*utils.User) error
    DeleteUser(*utils.User) error

    GetAllArtists(core.ArtistType) ([]core.Artist, error)
    GetArtist(string) (*core.Artist, error)
    InsertArtist(*core.Artist) error
    UpdateArtist(*core.Artist) error
    DeleteArtist(*core.Artist) error

    GetAllTechniques() ([]core.Technique, error)
    GetTechnique(string) (*core.Technique, error)
    InsertTechnique(*core.Technique) error
    UpdateTechnique(*core.Technique) error
    DeleteTechnique(*core.Technique) error

    GetAllStyles() ([]core.Style, error)
    GetStyle(string) (*core.Style, error)
    InsertStyle(*core.Style) error
    UpdateStyle(*core.Style) error
    DeleteStyle(*core.Style) error

    GetAllDatings() ([]core.Dating, error)
    GetDating(string) (*core.Dating, error)
    UpdateDating(*core.Dating) error
}

// Init database factory: returns appropriate DB connection URL, initialized
// DB connector and initialized DB data provider, according to DB type given.
// Note that DB connector and data provider are, in general, implemented 
// by the same type. 
func InitDb(dbtype DbType, host string, port int,
            username, password, dbname string) (url string, db DbConnector,
            data DataProvider, e error) {

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
func initMongo(host string, port int, username, passwd,
                dbname string) (string, *MongoDbConn, *MongoDbConn) {

    // create DB connection URL for MongoDB
    s := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", username, passwd,
                                                host, port, dbname)

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
    IdToString(interface{}) string
    StringToId(string) interface{}
}
