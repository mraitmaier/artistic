/*
    mongo.go -
 */
package db

import (
    "fmt"
    "time"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
    "bitbucket.org/miranr/artistic/utils"
    "bitbucket.org/miranr/artistic/core"
)

func CreateUrl(host string, port int,
               username, passwd string, dbname string) string {
    s := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", username, passwd,
                                                 host, port, dbname)
    return s
}


func Connect(url string, timeout time.Duration) (*mgo.Session, error) {
   return mgo.DialWithTimeout(url, timeout)
}

func Close(dbsess *mgo.Session) {
    if dbsess != nil {
        dbsess.Close()
    }
}

func MongoGetUser(db *mgo.Database, username string) (*utils.User, error) {

    u := utils.CreateUser("", "") // create empty user

    // get all users from DB
    err := db.C("users").Find(bson.M{"username": username }).One(&u)
    if err != nil { return nil, err }

    return u, nil
}

// retrieves all users from DB
func MongoGetAllUsers(db *mgo.Database) ([]utils.User, error) {

    // prepare the empty slice for users
    u := make([]utils.User, 0)

    // get all users from DB
    if err := db.C("users").Find(bson.D{}).All(&u); err != nil {
        return nil, err
    }

    return u, nil
}

// retrieves all datings from DB
func MongoGetAllDatings(db *mgo.Database) ([]core.Dating, error) {

    // prepare the empty slice for users
    d := make([]core.Dating, 0)

    // get all users from DB
    if err := db.C("datings").Find(bson.D{}).All(&d); err != nil {
        return nil, err
    }

    return d, nil
}

// retrieves all styles from DB
func MongoGetAllStyles(db *mgo.Database) ([]core.Style, error) {

    // prepare the empty slice for users
    s := make([]core.Style, 0)

    // get all users from DB
    if err := db.C("styles").Find(bson.D{}).All(&s); err != nil {
        return nil, err
    }

    return s, nil
}

// retrieves all styles from DB
func MongoGetAllTechniques(db *mgo.Database) ([]core.Technique, error) {

    // prepare the empty slice for users
    t := make([]core.Technique, 0)

    // get all users from DB
    if err := db.C("styles").Find(bson.D{}).All(&t); err != nil {
        return nil, err
    }

    return t, nil
}

