package artistic

import (
	"fmt"
	"time"
)

/*
 * Note - a type representing a single note:
 * - a string representing a note itself and
 * - a formatted timestamp (format: "2012-12-15 15:05:05")
 */
type Note struct {

	/* a string representing a note */
	note string

	/* a string representing a formatted timestamp */
	created string
}

/* display a note */
func (n *Note) String() string {
	return fmt.Sprintf("[%s] %s\n", n.created, n.note)
}

/* create a new note with given string; timestamp is created (and formatted)
 * automatically */
 /*
func CreateNote(n string) *Note {
	t := time.Now()
	note := &Note{n, t.Format("2012-12-15 15:04:05")}
	return note
}
*/

/* append a new note to an existing slice of notes */
/*
func AppendNote(notes []Note, n *Note) []Note {
	if n != nil && notes != nil {
		notes = append(notes, *n)
	}
	return notes
}
*/

func AppendNote(notes []Note, s string) []Note {
    if notes != nil {
        t := time.Now()
        note := &Note{s, t.Format("2012-12-15 15:04:05")}
        notes = append(notes, *note)
    }
    return notes
}

/* extend existing slice of notes with new slice of notes */
/*
func ExtendNotes(existing []Note, toadd []Note) []Note {
	return append(existing, toadd...)
}
*/

