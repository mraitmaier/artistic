package core

import (
	"fmt"
)

// Article is type representing a single newspaper/magazine/web article.
type Article struct {

	//
	Title string

	//
	Authors string

	//
	Publication string

	//
	Volume string

	//
	Issue string

	//
	Year string

	//
	Location string

	//
	Publisher string

	//
	ISSN string

	// Link is a link to the article online (if possible)
	Link string

	//
	Notes string

	//
	Keywords string
}

// NewArticle creates a new instance of Article.
func NewArticle() *Article {
	return &Article{
		Title:       "",
		Authors:     "",
		Publication: "",
		Volume:      "",
		Issue:       "",
		Year:        "",
		Location:    "",
		Publisher:   "",
		ISSN:        "",
		Link:        "",
		Notes:       "",
		Keywords:    "",
	}
}

//
func (b *Article) String() string {
	return fmt.Sprintf("%s: %s; %s Vol. %s Issue %s, %s, %s",
		b.Authors, b.Title, b.Publication, b.Volume, b.Issue, b.Year, b.Location)
}
