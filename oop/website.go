package oop

import "fmt"

type website struct {
	// It is not possible to anonymously embed a slice([]posts). A field name is required.
	posts []post
}

func (w website) contents() {
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}