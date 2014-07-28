package core

import (
	"fmt"
    "encoding/json"
)

/*
 * Technique - a type representing an art technique
 */
type Technique struct {

	/* name of the technique */
	Name string

	/* description of the technique */
	Description string
}

func (t *Technique) String() string { return t.Name }

func (t *Technique) Display() string {
	return fmt.Sprintf("%s\n%s\n", t.Name, t.Description)
}

// serialize a list of techniques into JSON
func techniquesToJson(items []Technique) (data string, err error) {

    var b []byte
    if b, err = json.Marshal(items); err != nil {
        return
    }
    data = string(b[:])
    return
}
