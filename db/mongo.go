/*
    mongo.go -
 */
package db

import (
    "fmt"
    "time"
    "labix.org/v2/mgo"
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
