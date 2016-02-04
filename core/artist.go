package core

import (
	//"fmt"
	//	"time"
	"encoding/json"
)

// Artist is type representing a single artist
// This type is a basis for all artists: painters, sculptors, writers etc.
type Artist struct {

	//Name of the artist
	*Name

	//RealName: if the person used pseudonim; with artists this is common
	RealName *Name

	// Born reperesents the date of birth
	Born string

	// Birthplace represents the place of birth
	Birthplace string

	// Died represents the date of death
	Died string

	// Deathplace represents the place of death
	Deathplace string

	// Nationality: his/her nationality
	Nationality string

	// Sources is a list of sources about the person
	Sources []string

	// Biography: his/her biography, of course
	Biography string

	// Notes is a list of private notes about the artist
	Notes []Note

	// Pricture holds the path (URL) to artist's picture
	//    Picture url.URL;
	Picture string

	// Set of artist flags, an artist can be many things at once
	IsPainter      bool `bson:"is_painter"`
	IsSculptor     bool `bson:"is_sculptor"`
	IsPrintmaker   bool `bson:"is_printmaker" `
	IsArchitect    bool `bson:"is_architect"`
	IsPhotographer bool
	IsIllustrator  bool
	IsDesigner     bool
	IsWriter       bool
	IsPoet         bool
	IsPlayWriter   bool
}

// DefaultCapacity defines the default number of items for slices
const DefaultCapacity = 10

// CreateArtist creates a new instance of Artist.
func CreateArtist() *Artist {
	return &Artist{
		Name:           &Name{"", "", ""},
		RealName:       &Name{"", "", ""},
		Born:           "",
		Birthplace:     "",
		Died:           "",
		Deathplace:     "",
		Nationality:    "",
		Sources:        make([]string, 0, DefaultCapacity),
		Biography:      "",
		Notes:          make([]Note, 0, DefaultCapacity),
		Picture:        "",
		IsPainter:      false,
		IsSculptor:     false,
		IsPrintmaker:   false,
		IsArchitect:    false,
		IsPhotographer: false,
		IsIllustrator:  false,
		IsDesigner:     false,
		IsWriter:       false,
		IsPoet:         false,
		IsPlayWriter:   false}
}

// String returns a string representation of the Artist
func (p *Artist) String() string { return p.Name.String() }

// JSON returns JSON-encoded representation of the Artist
func (p *Artist) JSON() (string, error) {
	s, err := json.Marshal(p)
	return string(s[:]), err
}

// Serialize a list of artists into JSON
func artistsToJSON(items []Artist) (data string, err error) {

	var b []byte
	if b, err = json.Marshal(items); err != nil {
		return
	}
	data = string(b[:])
	return
}
