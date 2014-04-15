package core

import (
	//    "fmt"
	//    "time"
	//    "net/url"
	"encoding/json"
)

/*****************************************************************************
 * type Person - struct representing the person
 *
 * This type is a basis for all artists: painters, sculptors, writers etc.
 *
 * Implements the Artist interface.
 */
type Person struct {

	/* name of the person */
	*Name

	/* real name if the person used pseudonim; with artists this is common */
	RealName *Name

	/* Date of birth */
	Born string

	/* Date of death */
	Dead string

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
    Painter bool
    Sculptor bool
    Printmaker bool
    Architect bool
    Ceramicist bool
}

const DefaultPersonCapacity = 10

func CreatePerson() *Person {
	n := make([]Note, 0, DefaultPersonCapacity)
	s := make([]string, 0, DefaultPersonCapacity)
	return &Person{&Name{"", "", ""}, &Name{"", "", ""},
		"", "", "", s, "", n, "", false, false, false, false, false }
}

/*
 * Person.String - a string representation of the Person
 */
func (p *Person) String() string {
	return p.Name.String()
}

/*
 * Person.Json- a JSON-encoded representation of the Person
 */
func (p *Person) Json() (string, error) {
	s, err := json.Marshal(p)
	return string(s[:]), err
}

/**
    PersonList - this is type representing a list of persons
*/
type PersonList []Person
