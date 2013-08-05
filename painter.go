package artistic

import (
	//    "fmt"
	"encoding/json"
)

/*****************************************************************************
 * type Painter - struct representing the person
 *
 * Implements the Artist interface.
 * Just a wrapper around a Person type - a way of implement inheritance in Go
 */
type Painter struct {
	/* Painter is a just a person */
	*Person
}

func CreatePainter() *Painter {
	return &Painter{CreatePerson()}
}

/*
 * Painter.String - a string representation of the Painter
 */
func (p *Painter) String() string {
	return p.Person.String()
}

/*
 * Painter.Json- a JSON-encoded representation of the Painter
 */
func (p *Painter) Json() (string, error) {
	s, err := json.Marshal(p.Person)
	return string(s[:]), err
}
