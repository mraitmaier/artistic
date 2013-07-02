package artistic

import (
	"fmt"
)

/*
 * Style - a type representing an art style
 */
type Style struct {
	/* name of the style */
	Name string
	/* description of the style */
	Description string
}

func (s *Style) String() string { return s.Name }

func (s *Style) Display() string {
	return fmt.Sprintf("%s\n%s\n", s.Name, s.Description)
}
