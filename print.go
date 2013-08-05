package artistic

import (
	//    "fmt"
	"encoding/json"
)

type Print struct {
	*Work
}

func (p *Print) String() string { return p.Work.Title }

func (p *Print) Json() (string, error) {
	s, err := json.Marshal(p.Work)
	return string(s[:]), err
}

func NewPrint() *Print { return &Print{CreateNewWork()} }

func (p *Print) Display() string { return p.Work.Display() }

func (p *Print) DisplaySources() string { return p.Work.DisplaySources() }

func (p *Print) DisplayNotes() string { return p.Work.DisplayNotes() }

func (p *Print) DisplayExhibitions() string {
	return p.Work.DisplayExhibitions()
}
