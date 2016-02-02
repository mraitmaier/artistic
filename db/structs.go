//
//
package db

import (
	// "fmt"
	"github.com/mraitmaier/artistic/core"
	"github.com/mraitmaier/artistic/utils"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Timestamp is the helper type representing timestamp according to RFC 822.
type Timestamp string

// NewTimestamp creates a new Timestamp using time.Now().
func NewTimestamp() Timestamp { return Timestamp(time.Now().Format(time.RFC822)) }

// String returns the standard string represenatation of the Timestamp.
func (t Timestamp) String() string { return string(t) }

// Dating represents the Database version of the core Dating type.
type Dating struct {

	// database ID
	Id bson.ObjectId `bson:"_id"`
	//  embed the instance of the DB Identifier interface, this is generalized DB ID
	//Id DbIdentifier `bson:"_id"` // XXX: need more knowledge about that...

	// original core Dating struct is embedded
	core.Dating `bson:",inline"`

	// created timestamp; SHOULD be read-only,
	Created Timestamp

	//  modified
	Modified Timestamp
}

// NewDating creates new Dating instance to be used for web page.
func NewDating(d *core.Dating) *Dating {
	t := NewTimestamp()
	return &Dating{bson.NewObjectId(), *d, t, t}
}

///
type Technique struct {

	// database ID
	Id bson.ObjectId `bson:"_id"`
	//  embed the instance of the DB Identifier interface, this is generalized DB ID
	//Id DbIdentifier

	// original core Technique struct is embedded
	core.Technique `bson:",inline"`

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

func NewTechnique() *Technique {
	return &Technique{bson.NewObjectId(), core.Technique{}, NewTimestamp(), NewTimestamp()}
}

func CreateTechnique(name, descr string) *Technique {
	return &Technique{bson.NewObjectId(), core.Technique{name, descr}, NewTimestamp(), NewTimestamp()}
}

////
type Style struct {

	// database ID
	Id bson.ObjectId `bson:"_id"`
	//  embed the instance of the DB Identifier interface, this is generalized DB ID
	//Id DbIdentifier `bson:"_id,inline"`

	// original core Sytle struct is embedded
	core.Style `bson:",inline"`

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

func NewStyle() *Style {
	return &Style{bson.NewObjectId(), *core.NewStyle("", ""), NewTimestamp(), NewTimestamp()}
}

////
type User struct {

	// database ID
	Id bson.ObjectId `bson:"_id"`
	//  embed the instance of the DB Identifier interface, this is generalized DB ID
	//Id DbIdentifier `bson:"_id,inline"`

	// original User struct is embedded
	utils.User `bson:",inline"`

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

// NewUser creates new empty DB user instance.
func NewUser() *User {
	return &User{
		bson.NewObjectId(),
		*utils.CreateUser("", "", "guest", "", "", "", true, checkRole),
		NewTimestamp(),
		NewTimestamp()}
}

// CreateUser creates new DB user instance. We need basic stuff: username, password, user role; 'create' is a flag denoting
// if this user's password must be hashed (insert new user into DB) or no (existing user).
func CreateUser(user, pwd, role string, create bool) *User {
	return &User{
		bson.NewObjectId(),
		*utils.CreateUser(user, pwd, role, "Change Myname", "email@blah.org", "", create, checkRole),
		NewTimestamp(),
		NewTimestamp()}
}

// This is the the list of valid user roles.
var AllRoles = []string{"admin", "user", "guest"}

// The checkRole() is a helper function that checks if user roles is valid.
func checkRole(val string) bool {

	for _, v := range AllRoles {
		if v == val {
			return true
		}
	}
	return false
}

////
type ArtistType int

const (
	ArtistTypeUnknown ArtistType = iota
	ArtistTypeArtist             // this is used as "whatever" or "all-types" value
	ArtistTypePainter
	ArtistTypeSculptor
	ArtistTypeArchitect
	ArtistTypePrintmaker
	ArtistTypeCeramicist
	ArtistTypeWriter
	ArtistTypePoet
	ArtistTypePlaywriter
)

func (t ArtistType) String() string {
	var s string
	switch t {
	case ArtistTypeArtist:
		s = "Artist"
	case ArtistTypePainter:
		s = "Painter"
	case ArtistTypeSculptor:
		s = "Sculptor"
	case ArtistTypeArchitect:
		s = "Architect"
	case ArtistTypePrintmaker:
		s = "Printmaker"
	case ArtistTypeCeramicist:
		s = "Ceramicist"
	case ArtistTypeWriter:
		s = "Writer"
	case ArtistTypePoet:
		s = "Poet"
	case ArtistTypePlaywriter:
		s = "Playwriter"
	default:
		s = "Unknown Artist Type"
	}
	return s
}

type Artist struct {

	// database ID
	Id bson.ObjectId `bson:"_id"`

	// original core Artist struct is embedded
	core.Artist `bson:",inline"`

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

func NewArtist() *Artist {
	return &Artist{bson.NewObjectId(), *core.CreateArtist(), NewTimestamp(), NewTimestamp()}
}

///
type Painting struct {

	// database ID
	Id bson.ObjectId `bson:"_id"`

	// original core Artwork interface is embedded
	core.Painting `bson:",inline"`

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

func NewPainting() *Painting {
	return &Painting{bson.NewObjectId(), *core.NewPainting(), NewTimestamp(), NewTimestamp()}
}
