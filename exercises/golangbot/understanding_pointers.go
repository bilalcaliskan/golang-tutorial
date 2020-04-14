package main

import (
	"fmt"
	"reflect"
)

func createPointer() *float64 { // declare that the function returns a float64 pointer
	var myFloat = 98.5
	return &myFloat // return a pointer of the specified type(*float64 in that case)
}

func printPointer(myBoolPointer *bool) { // use a pointer type for this parameter
	fmt.Println(*myBoolPointer) // print the value at the pointer that gets passed in
}

func runUnderstandingPointers() {
	/*
	We can print the types of pointer variables using reflection
	 */
	fmt.Printf("\nBeginning of examples...\n")
	var myInt int
	fmt.Printf("Type of &myInt is %v\n", reflect.TypeOf(&myInt)) // get a pointer to myInt and print the pointer's type
	var myFloat float64
	fmt.Printf("Type of &myFloat is %v\n", reflect.TypeOf(&myFloat)) // get a pointer to myFloat and print the pointer's type
	var myBoolean bool
	fmt.Printf("Type of &myBoolean is %v\n", reflect.TypeOf(&myBoolean)) // get a pointer to myBoolean and print the pointer's type

	/*
	We can declare variables that holds pointers. A pointer value can only hold pointers to one type of value, so a variable
	might only hold *int pointers, only *float64 pointers and so on.
	 */
	var myInt2 int = 32
	var myInt2Pointer *int // declare a variable that holds a pointer to an int
	myInt2Pointer = &myInt2 // assign a pointer of type int to the variable of type *int
	fmt.Printf("\nType of myInt2 is %v\n", reflect.TypeOf(myInt2))
	fmt.Printf("Value of myInt2 is %v\n", reflect.ValueOf(myInt2))
	fmt.Printf("Type of myInt2Pointer is %v\n", reflect.TypeOf(myInt2Pointer))
	fmt.Printf("Value of myInt2Pointer is %v\n", reflect.ValueOf(myInt2Pointer))
	fmt.Println(myInt2Pointer)

	/*
	You can get the value of the variable a pointer refers to by typing the * operator right before the pointer in your
	code. To get the value at myIntPointer, for example, you’d type *myIntPointer. (There’s no official consensus on how
	to read * aloud, but we like to pronounce it as “value at,” so *myIntPointer is “value at myIntPointer.”)
	 */
	myInt3 := 4
	myInt3Pointer := &myInt3
	fmt.Println(myInt3Pointer) // print the pointer itself
	fmt.Println(*myInt3Pointer) // print the value at the pointer

	/*
	The * operator can also be used to update the value at a pointer
	*/
	*myInt3Pointer = 8
	fmt.Println(myInt3Pointer) // print the pointer itself
	fmt.Println(*myInt3Pointer) // print the value at the pointer
	/*
	In the code above, *myInt3Pointer = 8 accesses the variable at myInt3Pointer (that is, the myInt3 variable) and assigns
	a new value to it. So not only is the value of *myInt3Pointer updated, but myInt is as well.
	 */

	/*
	It’s possible to return pointers from functions; just declare that the function’s return type is a pointer type.
	(By the way, unlike in some other languages in Go, it’s okay to return a pointer to a variable that’s local to a
	function. Even though that variable is no longer in scope, as long as you still have the pointer, Go will ensure
	you can still access the value.)
	*/
	var myFloatPointer *float64 = createPointer()
	fmt.Println(myFloatPointer) // print the pointer itself
	fmt.Println(*myFloatPointer) // print the value at the pointer

	/*
	You can also pass pointers to functions as arguments. Just specify that the type of one or more parameters should
	be a pointer.
	 */
	var myBool bool = true
	printPointer(&myBool) // pass a pointer to the function

}