package artistic

import (
//    "fmt"
    "encoding/json"
)

/*****************************************************************************
 * type Printmaker - struct representing the person
 *
 * This type is a basis for all artists: painters, sculptors, writers etc.
 *
 * Implements the Artist interface.
 */
type Printmaker struct {
    *Person
}

func CreatePrintmaker() *Printmaker {
    return &Printmaker{CreatePerson()}
}

/*
 * Printmaker.String - a string representation of the Printmaker
 */
func (p *Printmaker) String() string { return p.Person.String() }

/*
 * Printmaker.Json- a JSON-encoded representation of the Printmaker
 */
func (p *Printmaker) Json() (string, error) {
    s, err := json.Marshal(p.Person)
    return string(s[:]), err
}
