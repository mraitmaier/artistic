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

// retrieves all users from DoB
func MongoGetAllUsers(db *mgo.Database) ([]utils.User, error) {

    //db := aa.DbSess.DB("artistic")

    // prepare the empty slice for users
    u := make([]utils.User, 0)

    // get all users from DB
    if err := db.C("users").Find(bson.D{}).All(&u); err != nil {
        return nil, err
    }

    return u, nil
}

