package error_handling

import "fmt"

type person struct {
	firstName string
	lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}