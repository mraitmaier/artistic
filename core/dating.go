package core

import (
	"fmt"
)

var AllowedDatings = []string{ "L", "S", "A",
                               "a.q.", "a.q.n.", "p.q.", "p.q.n.",
                               "unknown" }

/**
 * Dating - a structure representing a dating
 */
type Dating struct {

    /* a dating value is defined (as enum) above */
	Dating string

    /* this is description of a dating */
	Description string
}

func (d *Dating) String() string {
	return d.Dating
}

func (d *Dating) Display() string {
	s := fmt.Sprintf("%q\n%s\n", d.Dating, d.Description)
	return s
}

