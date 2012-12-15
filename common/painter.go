package artistic

import (
//    "fmt"
    "encoding/json"
)

/*****************************************************************************
 * type Painter - struct representing the person
 *
 * This type is a basis for all artists: painters, sculptors, writers etc.
 *
 * Implements the Artist interface.
 */
type Painter struct {
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

func CreatePainter() *Painter {
    n := make([]Note, 0, DefaultPersonCapacity)
    s := make([]string, 0, DefaultPersonCapacity)
    return &Painter{&Name{"", "", ""}, &Name{"", "", ""}, 
                   "", "", "", s, "", n, ""}
}

/*
 * Painter.String - a string representation of the Painter
 */
func (p *Painter) String() string {
    return p.Name.String()
}

/*
 * Painter.Json- a JSON-encoded representation of the Painter
 */
func (p *Painter) Json() (string, error) {
    s, err := json.Marshal(p)
    return string(s[:]), err
}
