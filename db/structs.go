package db

//

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
	ID bson.ObjectId `bson:"_id"`
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
	ID bson.ObjectId `bson:"_id"`
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
	return &Technique{bson.NewObjectId(), core.Technique{name, descr, core.UnknownTechnique}, NewTimestamp(), NewTimestamp()}
}

////
type Style struct {

	// database ID
	ID bson.ObjectId `bson:"_id"`
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
	ID bson.ObjectId `bson:"_id"`
	//  embed the instance of the DB Identifier interface, this is generalized DB ID
	//Id DbIdentifier `bson:"_id,inline"`

	// original User struct is embedded
	utils.User `bson:",inline"`

	// Image represents the image of the user, the JPEG or PNG.
	Image []byte

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

// NewUser creates new empty DB user instance.
func NewUser() *User {
	t := NewTimestamp()
	return &User{
		ID:       bson.NewObjectId(),
		User:     *utils.CreateUser("", "", "guest", "", "", "", true, CheckRole),
		Image:    make([]byte, 0),
		Created:  t,
		Modified: t,
	}
}

// CreateUser creates new DB user instance. We need basic stuff: username, password, user role; 'create' is a flag denoting
// if this user's password must be hashed (insert new user into DB) or no (existing user).
func CreateUser(user, pwd, role string, create bool) *User {
	t := NewTimestamp()
	return &User{
		ID:       bson.NewObjectId(),
		User:     *utils.CreateUser(user, pwd, role, "Change Myname", "email@blah.org", "", create, CheckRole),
		Image:    make([]byte, 0),
		Created:  t,
		Modified: t,
	}
}

// This is the the list of valid user roles.
var AllRoles = []string{"admin", "user", "guest"}

// CheckRole() is a helper function that checks if user roles is valid.
func CheckRole(val string) bool {

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
	ID bson.ObjectId `bson:"_id"`

	// original core Artist struct is embedded
	core.Artist `bson:",inline"`

	// image (PNG or JPEG) of the artist
	Image []byte

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

// NewArtist creates a new instance of Artist
func NewArtist() *Artist {
	t := NewTimestamp()
	return &Artist{
		ID:       bson.NewObjectId(),
		Artist:   *core.CreateArtist(),
		Image:    make([]byte, 0),
		Created:  t,
		Modified: t,
	}
}

///
type Painting struct {

	// database ID
	ID bson.ObjectId `bson:"_id"`

	// original core Artwork interface is embedded
	core.Painting `bson:",inline"`

	// image of the (JPEG or PNG) of the painting
	Image []byte

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

// NewPainting creates a new instance of Painting
func NewPainting() *Painting {
	t := NewTimestamp()
	return &Painting{
		ID:       bson.NewObjectId(),
		Painting: *core.NewPainting(),
		Image:    make([]byte, 0),
		Created:  t,
		Modified: t,
	}
}

///
type Sculpture struct {

	// database ID
	ID bson.ObjectId `bson:"_id"`

	// original core Artwork interface is embedded
	core.Sculpture `bson:",inline"`

	// image of the (JPEG or PNG) of the sculpture
	Image []byte

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

//
func NewSculpture() *Sculpture {
	t := NewTimestamp()
	return &Sculpture{
		ID:        bson.NewObjectId(),
		Sculpture: *core.NewSculpture(),
		Image:     make([]byte, 0),
		Created:   t,
		Modified:  t,
	}
}

///
type Print struct {

	// database ID
	ID bson.ObjectId `bson:"_id"`

	// original core Artwork interface is embedded
	core.Print `bson:",inline"`

	// image of the (JPEG or PNG) of the print
	Image []byte

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

//
func NewPrint() *Print {
	t := NewTimestamp()
	return &Print{
		ID:       bson.NewObjectId(),
		Print:    *core.NewPrint(),
		Image:    make([]byte, 0),
		Created:  t,
		Modified: t,
	}
}

/// buildings

// Building is a DB wrapper around core Building type and represent a building as an artwork.
type Building struct {

	// database ID
	ID bson.ObjectId `bson:"_id"`

	// original core Artwork interface is embedded
	core.Building `bson:",inline"`

	// image of the (JPEG or PNG) of the  building
	Image []byte

	// created timestamp
	Created Timestamp

	// modified timestamp
	Modified Timestamp
}

func NewBuilding() *Building {
	t := NewTimestamp()
	return &Building{
		ID:       bson.NewObjectId(),
		Building: *core.NewBuilding(),
		Image:    make([]byte, 0),
		Created:  t,
		Modified: t,
	}
}

// Book type is a MongoDB wrapper for core Book type.
type Book struct {

	// ID represents the MongoDB Object ID
	ID bson.ObjectId `bson:"_id"`

	// Book is embedded core Book struct
	core.Book

	// Created and Modified represent ordinary database timestamps
	Created, Modified Timestamp
}

// NewBook creates a new instance of Book.
func NewBook() *Book {
	t := NewTimestamp()
	return &Book{
		ID:       bson.NewObjectId(),
		Book:     *core.NewBook(),
		Created:  t,
		Modified: t}
}

// Article type is a MongoDB wrapper for core Article type.
type Article struct {

	// ID represents the MongoDB Object ID
	ID bson.ObjectId `bson:"_id"`

	// Article is embedded core Artcile struct
	core.Article

	// Created and Modified represent ordinary database timestamps
	Created, Modified Timestamp
}

// NewArticle creates a new instance of Article.
func NewArticle() *Article {
	t := NewTimestamp()
	return &Article{
		ID:       bson.NewObjectId(),
		Article:  *core.NewArticle(),
		Created:  t,
		Modified: t}
}
