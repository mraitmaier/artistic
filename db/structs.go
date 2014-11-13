//
//
package db

import (
   // "fmt"
   // "time"
   "gopkg.in/mgo.v2/bson"
   "bitbucket.org/miranr/artistic/core"
   "bitbucket.org/miranr/artistic/utils"
)

//
type Timestamp struct {

    // we keep the actual data read-only
    timestamp string
}

func NewTimestamp() Timestamp {
    //creat := time.Now().Format("2012-12-15 15:0405")
    return Timestamp {""}
}

func (t Timestamp) String() string {
    return t.timestamp
}

func (t Timestamp) Update(stamp string) {
    t.timestamp = stamp
}

///
type Dating  struct {

    Id bson.ObjectId `bson:"_id"`
    //  embed the instance of the DB Identifier interface, this is generalized DB ID
    //Id DbIdentifier `bson:"_id"` // XXX: need more knowledge about that...

    // original core Dating struct is embedded
    core.Dating `bson:",inline"`

    // created timestamp; SHOULD be read-only, 
    //Created string
    Created Timestamp

    //  modified 
    //Modified string
    Modified Timestamp
}

// create new Dating instance to be used for web page
func NewDating() *Dating {
    return &Dating{ bson.NewObjectId(), core.Dating{}, NewTimestamp(), NewTimestamp() }
}

///
type Technique struct {

    Id bson.ObjectId `bson:"_id"`
    //  embed the instance of the DB Identifier interface, this is generalized DB ID
    //Id DbIdentifier

    // original core Technique struct is embedded
    core.Technique `bson:",inline"`


    //
    Created Timestamp

    //  modified 
    Modified Timestamp
}

func NewTechnique() *Technique {
    return &Technique{ bson.NewObjectId(), core.Technique{}, NewTimestamp(), NewTimestamp() }
}

func CreateTechnique(name, descr string) *Technique {
    return &Technique{ bson.NewObjectId(), core.Technique{ name, descr }, NewTimestamp(), NewTimestamp() }
}

////
type Style struct {

    Id bson.ObjectId `bson:"_id"`
    //  embed the instance of the DB Identifier interface, this is generalized DB ID
    //Id DbIdentifier `bson:"_id,inline"`

    // original core Sytle struct is embedded
    core.Style `bson:",inline"`

    //
    Created Timestamp

    //  modified 
    Modified Timestamp
}

func NewStyle() *Style {
    return &Style{ bson.NewObjectId(), *core.NewStyle("",""), NewTimestamp(), NewTimestamp() }
}

////
type User struct {

    Id bson.ObjectId `bson:"_id"`
    //  embed the instance of the DB Identifier interface, this is generalized DB ID
    //Id DbIdentifier `bson:"_id,inline"`

    // original User struct is embedded
    utils.User `bson:",inline"`

    //
    Created Timestamp

    //  modified 
    Modified Timestamp
}

func NewUser() *User {
    return &User{ bson.NewObjectId(), *utils.CreateUser("",""), NewTimestamp(), NewTimestamp() }
}

func CreateUser(user, pwd string) *User {
    return &User{ bson.NewObjectId(), *utils.CreateUser(user, pwd), NewTimestamp(), NewTimestamp() }
}

////
type Artist struct {

    Id bson.ObjectId `bson:"_id"`
    //  embed the instance of the DB Identifier interface, this is generalized DB ID
    //Id DbIdentifier

    // original core Artist struct is embedded
    core.Artist `bson:",inline"`

    //
    Created Timestamp

    //  modified 
    Modified Timestamp
}

func NewArtist() *Artist {
    return &Artist{ bson.NewObjectId(), *core.CreateArtist(), NewTimestamp(), NewTimestamp() }
}

///
type Artwork struct {

    Id bson.ObjectId `bson:"_id"`
    //  embed the instance of the DB Identifier interface, this is generalized DB ID
    //Id DbIdentifier

    // original core Artwork interface is embedded
    core.Artwork `bson:",inline"`

    //
    Created Timestamp

    //  modified 
    Modified Timestamp

}

func NewArtwork(w *core.Artwork) *Artwork {
    return &Artwork{ bson.NewObjectId(), *w, NewTimestamp(), NewTimestamp() }
}

