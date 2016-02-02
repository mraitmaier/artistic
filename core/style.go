package core

import (
	"encoding/json"
	"fmt"
)

/*
 * Style - a type representing an art style
 */
type Style struct {

	// name of the style
	Name string

	// description of the style
	Description string
}

// Create new instance of Style type
func NewStyle(name, description string) *Style {
	return &Style{name, description}
}

func (s *Style) String() string { return s.Name }

func (s *Style) Display() string {
	return fmt.Sprintf("%s\n%s\n", s.Name, s.Description)
}

// serialize a list of styles into JSON
func stylesToJson(items []Style) (data string, err error) {

	var b []byte
	if b, err = json.Marshal(items); err != nil {
		return
	}
	data = string(b[:])
	return
}
