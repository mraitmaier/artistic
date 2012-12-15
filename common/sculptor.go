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

func CreateSculptor() *Sculptor {
    n := make([]Note, 0, DefaultPersonCapacity)
    s := make([]string, 0, DefaultPersonCapacity)
    return &Sculptor{&Name{"", "", ""}, &Name{"", "", ""}, 
                   "", "", "", s, "", n, ""}
}

/*
 * Sculptor.String - a string representation of the Sculptor
 */
func (p *Sculptor) String() string {
    return p.Name.String()
}

/*
 * Sculptor.Json- a JSON-encoded representation of the Painter
 */
func (p *Sculptor) Json() (string, error) {
    s, err := json.Marshal(p)
    return string(s[:]), err
}
