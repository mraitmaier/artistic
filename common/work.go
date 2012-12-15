package artistic

import (
    "fmt"
    "encoding/json"
)

type Work struct {
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

func (w *Work) String() string {
    return w.Title
}

func (w *Work) Json() (string, error) {
    s, err := json.Marshal(w)
    return string(s[:]), err
}

const DefaultWorkCapacity = 10
func NewWork() *Work {
    s := make([]string, 0, DefaultWorkCapacity) // sources
    e := make([]string, 0, DefaultWorkCapacity) // exhibitions
    n := make([]Note, 0, DefaultWorkCapacity) // remarks
    w := &Dating{Dating_Unknown, "Default Description"}
    return &Work{"", CreatePerson(), &Technique{"", ""}, "", w, "", "", "",
                 "", "", "", "", "", "", e, s, n, ""}
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


