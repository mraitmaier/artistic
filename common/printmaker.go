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

func CreatePrintmaker() *Printmaker {
    n := make([]Note, 0, DefaultPersonCapacity)
    s := make([]string, 0, DefaultPersonCapacity)
    return &Printmaker{&Name{"", "", ""}, &Name{"", "", ""}, 
                   "", "", "", s, "", n, ""}
}

/*
 * Printmaker.String - a string representation of the Printmaker
 */
func (p *Printmaker) String() string {
    return p.Name.String()
}

/*
 * Printmaker.Json- a JSON-encoded representation of the Printmaker
 */
func (p *Printmaker) Json() (string, error) {
    s, err := json.Marshal(p)
    return string(s[:]), err
}
