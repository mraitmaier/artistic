package artistic

import (
    "fmt"
    "encoding/json"
)

type Print struct {
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

func (p *Print) String() string {
    return p.Title
}

func (p *Print) Json() (string, error) {
    s, err := json.Marshal(p)
    return string(s[:]), err
}

func NewPrint() *Print {
    s := make([]string, 0, DefaultWorkCapacity) // sources
    e := make([]string, 0, DefaultWorkCapacity) // exhibitions
    n := make([]Note, 0, DefaultWorkCapacity)   // remarks
    w := &Dating{Dating_Unknown, "Default Description"}
    return &Print{"", CreatePerson(), &Technique{"", ""}, "", w, "", "", "",
                 "", "", "", "", "", "", e, s, n, ""}
}

func (p *Print) Display() string {
    s := fmt.Sprintf("%s, %s (%s)\n", p.Title,
                                     p.Artist.String(), p.TimeOfCreation)
    s = s + fmt.Sprintf("%s %s\n", p.Technique.String(), p.Size)
    s = s + fmt.Sprintf("%s %s\n", p.Motive, p.Signature)
    s = s + fmt.Sprintln(p.Place)
    s = s + fmt.Sprintln(p.Location)
    s = s + fmt.Sprintln(p.Description)
    s = s + fmt.Sprintln(p.Provenance)
    s = s + fmt.Sprintln(p.Condition)
    s = s + fmt.Sprintln(p.ConditionDescription)
    s = s + p.DisplaySources()
    s = s + p.DisplayExhibitions()
    s = s + p.DisplayNotes()
    return s
}

func (p *Print) DisplaySources() string {
    s := "Sources:\n"
    for _, src := range p.Sources {
        s = s + fmt.Sprintln(src)
    }
    return s
}

func (p *Print) DisplayNotes() string {
    s := "Notes:\n"
    for _, n := range p.Notes {
        s = s + n.String()
    }
    return s
}

func (p *Print) DisplayExhibitions() string {
    s := "Exhibitions:\n"
    for _, src := range p.Exhibitions {
        s = s + fmt.Sprintln(src)
    }
    return s
}


