package oop

import "fmt"

type author struct {
	firstName string
	lastName string
	bio string
}

func (a author) fullName() string {
	// Sprintf returns the required string, printf prints the string to stdout
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}