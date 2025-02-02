package oop

import "fmt"

func RunInheritance() {
	fmt.Printf("\nBeginning of introduction...\n")
	/*
	Go does not support inheritance, however it does support composition. The generic definition of composition is "put
	together". One example of composition is a car. A car is composed of wheels, engine and various other parts.
	Composition can be achieved in Go is by embedding one struct type into another.
	A blog post is a perfect example of composition. Each blog post has a title, content and author information. This
	can be perfectly represented using composition.
	*/

	fmt.Printf("\nBeginning of composition example...\n")
	author1 := author{"Naveen", "Ramanathan", "Golang Enthusiast"}
	post1 := post{"Inheritence in Go", "aslkdjalskdjsaldkj", author1}
	post1.details()

	fmt.Printf("\nBeginning of embedding slice of structures...\n")
	/*
	We can take this example one step further and create a website using a slice of blog posts :)
	In website.go, it is not possible to anonymously embed a slice. A field name is required.
	 */
	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structures",
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