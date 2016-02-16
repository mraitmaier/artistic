package core

import (
	"fmt"
)

// Book is type representing a single book.
type Book struct {

	//
	Title string

	//
	Authors string

	//
	Edition string

	//
	Publisher string

	//
	Year string

	//
	Location string

	//
	ISBN string

	// Front is a link to front page picture
	Front string

	//
	Notes string

    //
    Keywords string
}

// NewBook creates a new instance of Book.
func NewBook() *Book {
	return &Book{
		Title:     "",
		Authors:   "",
		Edition:   "",
		Publisher: "",
		Year:      "",
		Location:  "",
		ISBN:      "",
		Front:     "",
		Notes:     "",
        Keywords:  "",
	}
}

//
func (b *Book) String() string {
	return fmt.Sprintf("%s: %s %s; %s, %s, %s", b.Authors, b.Title, b.Edition, b.Publisher, b.Year, b.Location)
}
