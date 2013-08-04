package artistic

import (
	//    "fmt"
	"encoding/json"
)

/*****************************************************************************
 * type Architect - struct representing the architect
 *
 * Implements the Artist interface.
 */
type Architect struct {
	*Person
}

func CreateArchitect() *Architect { return &Architect { CreatePerson() } }

/*
 * Architect.String - a string representation of the Architect
 */
func (a *Architect) String() string { return a.Person.String() }

/*
 * Architect.Json- a JSON-encoded representation of the Architect
 */
func (a *Architect) Json() (string, error) {
	s, err := json.Marshal(a.Person)
	return string(s[:]), err
}
