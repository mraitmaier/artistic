package core

import (
	//    "fmt"
	"encoding/json"
)

type Sculpture struct {

    // embedded struct
	*Work
}

func (s *Sculpture) String() string { return s.Work.Title }

func (s *Sculpture) Json() (string, error) {
	st, err := json.Marshal(s.Work)
	return string(st[:]), err
}

func NewSculpture() *Sculpture {
    return &Sculpture{CreateNewWork()}
}

func (s *Sculpture) Display() string { return s.Work.Display() }

func (s *Sculpture) DisplaySources() string { return s.Work.DisplaySources() }

func (s *Sculpture) DisplayNotes() string { return s.Work.DisplayNotes() }

func (s *Sculpture) DisplayExhibitions() string {
	return s.Work.DisplayExhibitions()
}

func (s *Sculpture) Created() string { return s.Work.Created() }
