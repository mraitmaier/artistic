package artistic

import (
//    "fmt"
    "encoding/json"
)

/*****************************************************************************
 * type Sculptor - struct representing the person
 *
 * This type is a basis for all artists: painters, sculptors, writers etc.
 *
 * Implements the Artist interface.
 */
type Sculptor struct {
    *Person
}

func CreateSculptor() *Sculptor {
    return &Sculptor{CreatePerson()}
}

/*
 * Sculptor.String - a string representation of the Sculptor
 */
func (sc *Sculptor) String() string { return sc.Person.String() }

/*
 * Sculptor.Json- a JSON-encoded representation of the Painter
 */
func (sc *Sculptor) Json() (string, error) {
    s, err := json.Marshal(sc.Person)
    return string(s[:]), err
}
