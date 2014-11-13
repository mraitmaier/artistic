package core

import (
	//    "fmt"
	"encoding/json"
)

/*
 * Ceramic - type representing a single ceramic piece 
 *
 * Actually just a plain wrapper around Work type - the easiest way to emulate
 * the inheritance in Go
 */
type Ceramic struct {
    // 
	*Work

    // is this ceramic a pottery?
    IsPottery bool

    // is this ceramic a porcelain?
    IsPorcelain bool
}

func (p *Ceramic) String() string { return p.Work.Title }

func (p *Ceramic) Json() (string, error) {
	s, err := json.Marshal(p.Work)
	return string(s[:]), err
}

func NewCeramic() *Ceramic { return &Ceramic{CreateNewWork(), false, false} }

func (p *Ceramic) Display() string { return p.Work.Display() }

func (p *Ceramic) DisplaySources() string { return p.Work.DisplaySources() }

func (p *Ceramic) DisplayNotes() string { return p.Work.DisplayNotes() }

func (p *Ceramic) DisplayExhibitions() string {
	return p.Work.DisplayExhibitions()
}

//func (p *Ceramic) Created() string {
//    return p.Work.Created()
//}

// serialize a list of ceramics into JSON
func ceramicsToJson(items []Ceramic) (data string, err error) {

    var b []byte
    if b, err = json.Marshal(items); err != nil {
        return
    }
    data = string(b[:])
    return
}
