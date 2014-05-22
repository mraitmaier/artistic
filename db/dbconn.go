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

// interface to rule all DBs types 
type DbConnector interface {
    // connect & close to DB interface methods
    Connect(url string, timeout time.Duration) error
    Close()

    //  get data from DB interface methods
    GetAllUsers() ([]utils.User, error)
    GetUser(username string) (*utils.User, error)
    GetAllTechniques() ([]core.Technique, error)
    GetAllStyles() ([]core.Style, error)
    GetAllDatings() ([]core.Dating, error)
}

// Init database factory: returns appropriate DB connection URL and initialized
// DB connector, according to DB type given.
func InitDb(dbtype DbType, host string, port int,
    username, password, dbname string) (url string, db DbConnector, e error) {

    // initialize
    url = ""
    db = nil

    switch dbtype {

    case MongoDB:
        url, db = initMongo(host, port, username, password, dbname)
        return url, db, e

    }
    return url, db, fmt.Errorf("Unknown database: cannot connect.\n")
}

// Initialization function for MongoDB.
func initMongo(host string, port int,
               username, passwd, dbname string) (string, *MongoDbConn) {

    // create DB connection URL for MongoDB
    s := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", username, passwd,
                                                host, port, dbname)

    // create new instance of MongoDB Session
    db := new(MongoDbConn)
    db.name = dbname

    return s, db
}

