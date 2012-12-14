package artistic

import (
    "fmt"
//    "time"
//    "net/url"
    "encoding/json"
)

/*
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

/*
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
    Picture string;
}
const DefaultPersonCapacity = 10
func CreatePerson() *Person {
    n := make([]Note, 0, DefaultPersonCapacity)
    s := make([]string, 0, DefaultPersonCapacity)
    return &Person{&Name{"", "", ""}, &Name{"", "", ""}, 
                   "", "", "", s, "", n, ""}
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
