package core

/*
 * address.go
 */

import (
	"fmt"
)

// AddressType is...
type AddressType int

const (
	//
	AddressAngloSaxon AddressType = iota

	//
	AddressContinental
)

//
func (t AddressType) String() string {

	var s string
	switch t {
	case AddressAngloSaxon:
		s = "Anglosaxon address type"
	case AddressContinental:
		s = "Continental address type"
	}
	return s
}

// Address is ...
type Address struct {
	// Street..
	Street string

	// Number ...
	Number string

	// City ...
	City string

	// ZipCode ...
	ZipCode string

	// State ...
	State string

	// Country ...
	Country string

	// Type is
	Type AddressType
}

// NewAddress is...
func NewAddress() *Address { return &Address{"", "", "", "", "-", "", AddressContinental} }

//String - a string representation of the Name
func (a *Address) String() string {
	var s string
	switch a.Type {
	case AddressAngloSaxon:
		s = fmt.Sprintf("%s, %s\n%s %s\n%s\n%s", a.Number, a.Street, a.City, a.ZipCode, a.State, a.Country)
	case AddressContinental:
		s = fmt.Sprintf("%s, %s\n%s %s\n%s, %s", a.Street, a.Number, a.ZipCode, a.City, a.State, a.Country)
	}
	return s
}
