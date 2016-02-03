package core

/*
 * exhibition.go
 */

import (
	"fmt"
)

// Exhibition is ...
type Exhibition struct {

	// Title..
	Title string

	// Location ...
	Location string

	// City ...
	City string

	//  Time of exhibition...
	Time string

	// Description...
	Description string
}

// Newddress is...
func NewExhibition() *Exhibition { return &Exhibition{"New Exhibition", "", "", "", ""} }

//String - a string representation of the Name
func (e *Exhibition) String() string {
	var s string
	s = fmt.Sprintf("%s [%s]\n%s, %s\n%s", e.Title, e.Time, e.Location, e.City, e.Description)
	return s
}
