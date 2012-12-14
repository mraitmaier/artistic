
package artistic

import "fmt"

type Note struct {
    Note string
    Created string
}

func (n *Note) String() string {
    return fmt.Sprintf("[%s] %s\n", n.Created, n.Note)
}
