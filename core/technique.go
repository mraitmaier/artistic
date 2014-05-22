package core

import (
	"fmt"
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

/*
func (t *Technique) Create(dbconn *DbConn) error {

    // check DB connection; if emtpy, return error
    if dbconn == nil {
        return fmt.Errorf("DB connection is empty.")
    }

    
    return nil
}
*/
