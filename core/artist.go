package core

import (
	//"fmt"
	//	"time"
	"encoding/json"
)

// type representing the person
// This type is a basis for all artists: painters, sculptors, writers etc.
type Artist struct {

	/* name of the artist */
	*Name

	/* real name if the person used pseudonim; with artists this is common */
	RealName *Name

	/* Date of birth */
	Born string

	//
	Birthplace string

	/* Date of death */
	Died string

	//
	Deathplace string

	/* his/her nationality */
	Nationality string

	/* a list of sources about the person */
	Sources []string

	/* his/her biography, of course */
	Biography string

	/* private notes about the person */
	Notes []Note

	/* this is actually an URI */
	//    Picture url.URL;
	Picture string

	// artist flags, a person can be many things at once
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

const DefaultCapacity = 10

func CreateArtist() *Artist {
	//n := make([]Note, 0, DefaultCapacity)
	//s := make([]string, 0, DefaultCapacity)
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

/*
 * Artist.String - a string representation of the Artist
 */
func (p *Artist) String() string { return p.Name.String() }

/*
 * Artist.Json- a JSON-encoded representation of the Artist
 */
func (p *Artist) Json() (string, error) {
	s, err := json.Marshal(p)
	return string(s[:]), err
}

// serialize a list of artists into JSON
func artistsToJson(items []Artist) (data string, err error) {

	var b []byte
	if b, err = json.Marshal(items); err != nil {
		return
	}
	data = string(b[:])
	return
}
