package artistic

import (
    "fmt"
    "encoding/json"
)

type Statue struct {
    Title string
    Artist
    *Technique
    Size string
    *Dating
    TimeOfCreation string
    Motive string
    Signature string
    Place string
    Location string
    Provenance string
    Condition string
    ConditionDescription string
    Description string
    Exhibitions []string
    Sources []string
    Notes []Note
    Picture string
}

func (st *Statue) String() string {
    return st.Title
}

func (st *Statue) Json() (string, error) {
    s, err := json.Marshal(st)
    return string(s[:]), err
}

func NewStatue() *Statue {
    s := make([]string, 0, DefaultWorkCapacity) // sources
    e := make([]string, 0, DefaultWorkCapacity) // exhibitions
    n := make([]Note, 0, DefaultWorkCapacity)   // remarks
    w := &Dating{Dating_Unknown, "Default Description"}
    return &Statue{"", CreatePerson(), &Technique{"", ""}, "", w, "", "", "",
                 "", "", "", "", "", "", e, s, n, ""}
}

func (st *Statue) Display() string {
    s := fmt.Sprintf("%s, %s (%s)\n", st.Title,
                                     st.Artist.String(), st.TimeOfCreation)
    s = s + fmt.Sprintf("%s %s\n", st.Technique.String(), st.Size)
    s = s + fmt.Sprintf("%s %s\n", st.Motive, st.Signature)
    s = s + fmt.Sprintln(st.Place)
    s = s + fmt.Sprintln(st.Location)
    s = s + fmt.Sprintln(st.Description)
    s = s + fmt.Sprintln(st.Provenance)
    s = s + fmt.Sprintln(st.Condition)
    s = s + fmt.Sprintln(st.ConditionDescription)
    s = s + st.DisplaySources()
    s = s + st.DisplayExhibitions()
    s = s + st.DisplayNotes()
    return s
}

func (st *Statue) DisplaySources() string {
    s := "Sources:\n"
    for _, src := range st.Sources {
        s = s + fmt.Sprintln(src)
    }
    return s
}

func (st *Statue) DisplayNotes() string {
    s := "Notes:\n"
    for _, n := range st.Notes {
        s = s + n.String()
    }
    return s
}

func (st *Statue) DisplayExhibitions() string {
    s := "Exhibitions:\n"
    for _, src := range st.Exhibitions {
        s = s + fmt.Sprintln(src)
    }
    return s
}


