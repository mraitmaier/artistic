package core

import (
// "fmt"
)

type Institution struct {
	//
	Name string

	//
	Address

	//
	Established string

	//
	Description string

	//
	Picture string

	//
	Books []*Book

	//
	Articles []*Article
}

func NewInstitution() *Institution {
	return &Institution{
		Name:        "",
		Address:     *NewAddress(),
		Established: "",
		Description: "",
		Picture:     "",
		Books:       make([]*Book, 0),
		Articles:    make([]*Article, 0)}
}

func (i *Institution) String() string { return i.Name }
