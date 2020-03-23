package oop

import "fmt"

func RunInheritance() {
	fmt.Printf("\nBeginning of introduction...\n")
	fmt.Println("Go does not support inheritance, however it does support composition. The generic definition of " +
		"composition is 'put together'. One example of composition is a car. A car is composed of wheels, engine and " +
		"various other parts. ")
	fmt.Println("Composition can be achieved in Go is by embedding one struct type into another. ")

	fmt.Printf("\nBeginning of composition example...\n")
	author1 := author{"Naveen", "Ramanathan", "Golang Enthusiast"}
	post1 := post{"Inheritence in Go", "aslkdjalskdjsaldkj", author1}
	post1.details()

	fmt.Printf("\nBeginning of embedding slice of structs...\n")
	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{[]post{post1, post2, post3}}
	w.contents()
}