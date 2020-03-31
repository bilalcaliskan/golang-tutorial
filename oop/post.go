package oop

import "fmt"

type post struct {
	title string
	content string
	author // this field can also be declared as anonymous field
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Whenever one struct field is embedded in another, Go gives us the option to access the embedded " +
		"fields as if they were part of the outer struct. This means that p.author.fullName() in below line can be " +
		"replaced with p.fullName(). Hence the details() method can be rewritten as below,")
	fmt.Println("Author: ", p.author.fullName())
	fmt.Println("Like this")
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.author.bio)
}