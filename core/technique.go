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

