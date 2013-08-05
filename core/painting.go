package core

import (
	//    "fmt"
	"encoding/json"
)

/*
 * Painting - type representing a single painting
 *
 * Actually just a plain wrapper around Work type - the easiest way to emulate
 * the inheritance in Go
 */
type Painting struct {
	*Work
}

func (p *Painting) String() string { return p.Work.Title }

func (p *Painting) Json() (string, error) {
	s, err := json.Marshal(p.Work)
	return string(s[:]), err
}

func NewPainting() *Painting { return &Painting{CreateNewWork()} }

func (p *Painting) Display() string { return p.Work.Display() }

func (p *Painting) DisplaySources() string { return p.Work.DisplaySources() }

func (p *Painting) DisplayNotes() string { return p.Work.DisplayNotes() }

func (p *Painting) DisplayExhibitions() string {
	return p.Work.DisplayExhibitions()
}
