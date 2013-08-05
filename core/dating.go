package core

import (
	"fmt"
)

type DatingValue int

const (
	Dating_Unknown DatingValue = iota 
	Dating_L
	Dating_S
	Dating_A
	Dating_aq
	Dating_aqn
	Dating_pq
	Dating_pqn
)

func (d DatingValue) String() string {
	s := ""
	switch d {
	case Dating_L:
		s = "L"
	case Dating_S:
		s = "S"
	case Dating_A:
		s = "A"
	case Dating_aq:
		s = "a.q."
	case Dating_aqn:
		s = "a.q.n."
	case Dating_pq:
		s = "p.q."
	case Dating_pqn:
		s = "p.q.n."
	default:
		s = "Unknown dating"
	}
	return s
}

/**
 * Dating - a structure representing a dating
 */
type Dating struct {

    /* a dating value is defined (as enum) above */
	DatingValue

    /* this is description of a dating */
	Description string
}

func (d *Dating) String() string {
	return d.DatingValue.String()
}

func (d *Dating) Display() string {
	s := fmt.Sprintf("%q\n%s\n", d.DatingValue.String(), d.Description)
	return s
}

/* Convert a dating value into string value */
func ConvDating2str(val DatingValue) string {
	return val.String()
}

/* convert a dating value represented as string into DatingValue */
func ConvDating2val(s string) DatingValue {
	val := Dating_Unknown
	switch s {
	case "L":
		val = Dating_L
	case "S":
		val = Dating_S
	case "A":
		val = Dating_A
	case "a.q.", "aq", "a q":
		val = Dating_aq
	case "a.q.n.", "aqn", "a q n":
		val = Dating_aqn
	case "p.q.", "pq", "p q":
		val = Dating_pq
	case "p.q.n.", "pqn", "p q n":
		val = Dating_pqn
	}
	return val
}
