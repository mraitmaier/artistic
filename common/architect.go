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

func CreateArchitect() *Architect {
    n := make([]Note, 0, DefaultPersonCapacity)
    s := make([]string, 0, DefaultPersonCapacity)
    return &Architect{&Name{"", "", ""}, &Name{"", "", ""},
                            "", "", "", s, "", n, ""}
}

/*
 * Architect.String - a string representation of the Architect
 */
func (p *Architect) String() string { return p.Name.String() }

/*
 * Architect.Json- a JSON-encoded representation of the Painter
 */
func (p *Architect) Json() (string, error) {
    s, err := json.Marshal(p)
    return string(s[:]), err
}
