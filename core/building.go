package core

import (
	//    "fmt"
	"encoding/json"
)

//
type Building struct {
	//
	Work
}

func (p *Building) String() string { return p.Work.Title }

func (p *Building) Json() (string, error) {
	s, err := json.Marshal(p.Work)
	return string(s[:]), err
}

func NewBuilding() *Building { return &Building{*CreateNewWork()} }

func (p *Building) Display() string { return p.Work.Display() }

func (p *Building) DisplaySources() string { return p.Work.DisplaySources() }

func (p *Building) DisplayNotes() string { return p.Work.DisplayNotes() }

func (p *Building) DisplayExhibitions() string {
	return p.Work.DisplayExhibitions()
}

// serialize a list of buildings into JSON
func buildingsToJson(items []Building) (data string, err error) {

	var b []byte
	if b, err = json.Marshal(items); err != nil {
		return
	}
	data = string(b[:])
	return
}
