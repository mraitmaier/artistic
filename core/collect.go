package core

type ArtCollection struct {
	//
	Name string

	//
	Description string

	//
	Owner string

	//
	Location string
}

//
func NewArtCollection() *ArtCollection {
	return &ArtCollection{
		Name:        "",
		Description: "",
		Owner:       "",
		Location:    "",
	}
}

//
func (a *ArtCollection) String() string { return a.Name }
