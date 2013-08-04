package artistic

import (
	//    "fmt"
	"encoding/json"
)

type Statue struct {
	*Work
}

func (st *Statue) String() string { return st.Work.Title }

func (st *Statue) Json() (string, error) {
	s, err := json.Marshal(st.Work)
	return string(s[:]), err
}

func NewStatue() *Statue { return &Statue{CreateNewWork()} }

func (st *Statue) Display() string { return st.Work.Display() }

func (st *Statue) DisplaySources() string { return st.Work.DisplaySources() }

func (st *Statue) DisplayNotes() string { return st.Work.DisplayNotes() }

func (st *Statue) DisplayExhibitions() string {
	return st.Work.DisplayExhibitions()
}
