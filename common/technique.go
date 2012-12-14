
package artistic

type Technique struct {
    Name string
    Description string
}

func (t *Technique) String() string {
    return t.Name
}
