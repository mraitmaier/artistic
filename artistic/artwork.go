package artistic

type Artwork interface {
	String() string
	Json() (string, error)
}
