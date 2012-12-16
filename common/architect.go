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

func CreateArchitect() *Architect {
    return &Architect{CreatePerson()}
}

/*
 * Architect.String - a string representation of the Architect
 */
func (p *Architect) String() string { return p.Person.String() }

/*
 * Architect.Json- a JSON-encoded representation of the Painter
 */
func (p *Architect) Json() (string, error) {
    s, err := json.Marshal(p.Person)
    return string(s[:]), err
}
