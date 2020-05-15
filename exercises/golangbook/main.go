package main

import (
	"fmt"
)

func main() {
	// Check https://www.golang-book.com/books/intro for all
	// chapter10.RunChapter10()
	// chapter11.RunChapter11()

	/*
	A struct's fields usually represent the has-a relationship.
	Suppose we had a person struct. And we wanted to create a new Android struct.
	This would work, but we would rather say an Android is a Person, rather than an Android has a Person. Go supports
	relationships like this by using an embedded type. Also known as anonymous fields.
	*/
	a := new(android)
	a.model = "ahmet"
	a.person.talk()

	// This will work like, android is a person rather than android has a person. The is-a relationship works this way
	//intuitively: People can talk, an android is a person, therefore an android can talk. It looks like extends in Java.
	b := new(android2)
	b.model = "hasan"
	b.talk()
}

type person struct {
	name string
}

func (p *person) talk() {
	fmt.Println("Hi, my name is", p.name)
}

type android struct {
	person person
	model string
}

type android2 struct {
	person
	model string
}