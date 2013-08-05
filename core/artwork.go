package core

type Artwork interface {
	String() string
	Json() (string, error)
}
