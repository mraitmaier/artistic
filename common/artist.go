/*
 * person.go
 */
package artistic

import ( "fmt" )

type Artist interface {
    String() string
    Json() (string, error)
}

/*****************************************************************************
 * type Name - struct representing a person's name
 */
type Name struct {
    First string
    Middle string
    Last string
}

/*
 * Name.String - a string representation of the Name
 */
func (n *Name) String() string {
    return fmt.Sprint("%s %s %s", n.First, n.Middle, n.Last)
}

/*
 * CreateName - creates a new empty instance of the Name
 */
func CreateName (first string, middle string, last string) *Name {
    return &Name{first, middle, last}
}

