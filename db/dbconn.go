//
// dbconn.go
//
package db

import (
    "fmt"
    "time"
    "bitbucket.org/miranr/artistic/utils"
    "bitbucket.org/miranr/artistic/core"
)

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
    GetUser(username string) (*utils.User, error)
    //CreateUser(*utils.User) error
    UpdateUser(*utils.User) error

    GetAllTechniques() ([]core.Technique, error)
    GetTechnique(name string) (*core.Technique, error)
    //CreateTechnique(*core.Technique) error
    //UpdateTechnique(*core.Technique) error

    GetAllStyles() ([]core.Style, error)
    GetStyle(name string) (*core.Style, error)
    //CreateStyle(*core.Style) error
    //UpdateStyle(*core.Style) error

    GetAllDatings() ([]core.Dating, error)
    GetDating(name string) (*core.Dating, error)
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
