package core

import (
	"encoding/json"
	"fmt"
    "time"
)

type Artwork interface {
    Json() (string, error)
    Display() string
    Created() string
}

type Work struct {
	/* a title of the artwork */
	Title string

	/* an artist  */
	*Artist

	/* This artwork is made in which technique? */
	*Technique

	/* Size: dimensions, measures of the artwork */
	Size string

	/* dating of the artwork */
	*Dating

	/* When was this artwork created (painted, sculpted etc.)? */
	TimeOfCreation string

	/* a description of the artwork's motive */
	Motive string

	/* a desicription of the artwork's singnature, if applicable */
	Signature string

	/* where is artwork currently located (gallery, privately owned etc.)? */
	Place string

	/* where is artwork currently located (borrowed to any exhibition)? */
	Location string

	/* an artwork's provenance; a history of the artworks ownership */
	Provenance string

	/* a short, one-word note about artwork's condition: good, bad, lost etc. */
	Condition string

	/* a long description of the artwork's condition */
	ConditionDescription string

	/* a description of the artwork */
	Description string

	/* a list of known exhibitions for this artwork (if applicable) */
	Exhibitions []string

	/* personal sources (either web URIs or anything else) about this artwork */
	Sources []string

	/* personal notes on the artwork */
	Notes []Note

	/* URI (either local or web) to artwork's reproduction */
	Picture string

    /* timestamp when this instance was created */
    created string

    /* a timestamp when this instance was last updated */
    updated string
}

/* work's string representation */
func (w *Work) String() string {
	return w.Title
}

/* work's JSON representation */
func (w *Work) Json() (string, error) {
	s, err := json.Marshal(w)
	return string(s[:]), err
}

func (w *Work) Created() string {
    return w.created
}

/* create a new, default instance of the Work */
const DefaultWorkCapacity = 10

func CreateNewWork() *Work {
    /* generate a 'created' timestamp and apply it to the 'updated', too */
    creat := time.Now().Format("2012-12-15 15:04:05")
	s := make([]string, 0, DefaultWorkCapacity) // sources
	e := make([]string, 0, DefaultWorkCapacity) // exhibitions
	n := make([]Note, 0, DefaultWorkCapacity)   // remarks
	w := &Dating{Dating_Unknown, "Default Description"}
	return &Work{"", CreateArtist(), &Technique{"", ""}, "", w, "", "", "",
		"", "", "", "", "", "", e, s, n, "", creat, creat}
}

func (w *Work) Display() string {
	s := fmt.Sprintf("%s, %s (%s)\n", w.Title,
		w.Artist.String(), w.TimeOfCreation)
	s = s + fmt.Sprintf("%s %s\n", w.Technique.String(), w.Size)
	s = s + fmt.Sprintf("%s %s\n", w.Motive, w.Signature)
	s = s + fmt.Sprintln(w.Place)
	s = s + fmt.Sprintln(w.Location)
	s = s + fmt.Sprintln(w.Description)
	s = s + fmt.Sprintln(w.Provenance)
	s = s + fmt.Sprintln(w.Condition)
	s = s + fmt.Sprintln(w.ConditionDescription)
	s = s + w.DisplaySources()
	s = s + w.DisplayExhibitions()
	s = s + w.DisplayNotes()
	return s
}

func (w *Work) DisplaySources() string {
	s := "Sources:\n"
	for _, src := range w.Sources {
		s = s + fmt.Sprintln(src)
	}
	return s
}

func (w *Work) DisplayNotes() string {
	s := "Notes:\n"
	for _, n := range w.Notes {
		s = s + n.String()
	}
	return s
}

func (w *Work) DisplayExhibitions() string {
	s := "Exhibitions:\n"
	for _, src := range w.Exhibitions {
		s = s + fmt.Sprintln(src)
	}
	return s
}

/* WorkList - a list of Works */
type WorkList []Work

/*
const defaultWorkListLen = 10
func CreateDefaultWorkList() []WorkList {
    l := make([]WorkList, defaultWorkListLen)
    return l
}

func CreateWorkList(length, capacity int) []WorkList {
    l := make([]WorkList, length, capacity)
    return l
}

*/
