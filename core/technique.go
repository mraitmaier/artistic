package core

import (
	"encoding/json"
	"fmt"
)

//
type TechniqueType int

const (
	UnknownTechnique TechniqueType = iota
	PaintTechnique
	SculptTechnique
	GraphicTechnique
)

//GetTechniqueTypes retuens a slice of all possible TechniqueType-s.
func GetTechniqueTypes() []TechniqueType {

	var tt = make([]TechniqueType, 4)

	tt[0] = UnknownTechnique
	tt[1] = PaintTechnique
	tt[2] = SculptTechnique
	tt[3] = GraphicTechnique
	return tt
}

//
func (tt TechniqueType) String() string {

	switch tt {
	case PaintTechnique:
		return "Paint Technique"
	case SculptTechnique:
		return "Sculpting Technique"
	case GraphicTechnique:
		return "Graphic Technique"
	default:
		return "Unknown Technique"
	}
}

// Technique - a type representing an art technique
type Technique struct {

	// name of the technique
	Name string

	// description of the technique
	Description string

	// TechniqueType narrows down the technique: paint, sculpting, graphic etc. tecnique
	Type TechniqueType
}

func NewTechnique(name, description string) *Technique {
	return &Technique{
		Name:        name,
		Description: description,
		Type:        UnknownTechnique}
}

func (t *Technique) String() string { return t.Name }

func (t *Technique) Display() string {
	return fmt.Sprintf("%s (%s)\n%s\n", t.Name, t.Type, t.Description)
}

// serialize a list of techniques into JSON
func techniquesToJson(items []Technique) (data string, err error) {

	var b []byte
	if b, err = json.Marshal(items); err != nil {
		return
	}
	data = string(b[:])
	return
}
