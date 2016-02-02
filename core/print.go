package core

import (
	//    "fmt"
	"encoding/json"
)

type Print struct {
	Work
}

func (p *Print) String() string { return p.Work.Title }

func (p *Print) Json() (string, error) {
	s, err := json.Marshal(p.Work)
	return string(s[:]), err
}

func NewPrint() *Print { return &Print{*CreateNewWork()} }

func (p *Print) Display() string { return p.Work.Display() }

func (p *Print) DisplaySources() string { return p.Work.DisplaySources() }

func (p *Print) DisplayNotes() string { return p.Work.DisplayNotes() }

func (p *Print) DisplayExhibitions() string {
	return p.Work.DisplayExhibitions()
}

//func (p *Print) Created() string { return p.Work.Created() }

// serialize a list of prints into JSON
func printsToJson(items []Print) (data string, err error) {

	var b []byte
	if b, err = json.Marshal(items); err != nil {
		return
	}
	data = string(b[:])
	return
}
