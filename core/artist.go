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

	/* Date of death */
	Died string

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

	/* artist flags, a person can be many things at once */
	IsPainter    bool `bson:"is_painter"`
	IsSculptor   bool `bson:"is_sculptor"`
	IsPrintmaker bool `bson:"is_printmaker" `
	IsCeramicist bool `bson:"is_ceramicist"`

	IsArchitect bool `bson:"is_architect"`

	IsWriter     bool `bson:"is_writer"`
	IsPoet       bool `bson:"is_poet"`
	IsPlayWriter bool `bson:"is_playwriter"`
}

const DefaultArtistCapacity = 10

func CreateArtist() *Artist {
	n := make([]Note, 0, DefaultArtistCapacity)
	s := make([]string, 0, DefaultArtistCapacity)
	return &Artist{&Name{"", "", ""}, // Name
		&Name{"", "", ""}, // Real name
		"",                // Born
		"",                // Died
		"",                // Nationality
		s,                 // Sources
		"",                // Biography
		n,                 // Notes
		"",                // Picture
		false,             // IsPainter
		false,             // IsSculptor
		false,             // IsPrintmaker
		false,             // IsCeramicist
		false,             // IsArchitect
		false,             // IsWriter
		false,             // IsPoet
		false}             // IsPlayWriter
}

/*
 * Artist.String - a string representation of the Artist
 */
func (p *Artist) String() string {
	return p.Name.String()
}

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

/**
  ArtistList - this is type representing a list of persons
*/
type ArtistList []Artist
