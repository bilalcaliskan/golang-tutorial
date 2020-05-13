package chapter11

import (
	"fmt"
	"golang-tutorial/exercises/golangbook/chapter11/math"
)

func RunChapter11() {
	/*
	Chapter 11 is about Packages
	https://www.golang-book.com/books/intro/11
	*/
	xs := []float64{1,2,3,4}
	avg := math.Average(xs)
	fmt.Println(avg)
}