package pointers

import "fmt"

func ChangeWithAddress(val *int) {
	*val = 55
}

func ChangeWithValue(val int) {
	val = 50
}

func hello() *int {
	i := 5
	return &i
}

func modifyArray(arr *[3]int) {
	(*arr)[0] = 90
}

func modifySlice(sls []int) {
	sls[0] = 90
}

func RunPointers()  {
	fmt.Printf("\nBeginning of introduction to pointers\n")
	/*
	A pointer is a variable which stores the memory address of another variable.
	*T is the type of the pointer variable which points to a value of type T.
	*/
	b := 255
	var a *int = &b
	c := &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	fmt.Println("address of b is", c)
	/*
	The & operator is used to get the address of a variable. In line no. 25 of the above program we are assigning the
	address of b to a whose type is *int. Now a is said to point to b.
	 */

	fmt.Printf("\nBeginning of zero value of a pointer\n")
	/*
	The zero value of a pointer is nil
	 */
	d := 25
	var e *int
	if e == nil {
		fmt.Println("e is", e)
		e = &d
		fmt.Println("e after initialization is", e)
	}

	fmt.Printf("\nBeginning of creating pointers using the new function\n")
	/*
	Go also provides a handy function new to create pointers. The new function takes a type as argument and returns a
	pointer to a newly allocated zero value of the type passed as argument.
	 */
	size := new(int) // returns a pointer to a newly allocated zero value of the type passed
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Printf("New size value is %d, type is %T, address is %v\n", *size, size, size)

	fmt.Printf("\nBeginning of dereferencing a pointer\n")
	/*
	Dereferencing a pointer means accessing the value of the variable which the pointer points to. *a is the syntax to
	deference a.
	 */
	ab := 255
	ba := &ab
	fmt.Println("address of ab is", ba)
	fmt.Println("value of ab is", *ba) // here we are dereferencing a pointer to get the value which pointer points to
	/*
	Lets write one more program where we change the value in b using the pointer.
	 */
	ac := 255
	ca := &ac // declares a pointer which points to address of ac
	fmt.Println("address of ac is", ca)
	fmt.Println("value of ac is", *ca)
	*ca++
	fmt.Println("new value of ac is", *ca)
	*ca = 300
	fmt.Println("new value of ac is", *ca)

	fmt.Printf("\nBeginning of passing pointer to a function\n")
	ad := 58
	fmt.Println("value of ad before function call is", ad)
	da := &ad // declares a pointer which points to address of ad
	ChangeWithAddress(da)
	fmt.Println("value of ad after function call is", ad)
	/*
	In the above program, we are passing the pointer variable da which holds the address of ad to the function
	ChangeWithAddress(). Inside ChangeWithAddress() function, value of ad is changed using dereferencing.
	*/

	fmt.Printf("\nBeginning of returning pointer from a function\n")
	/*
	It is perfectly legal for a function to return a pointer of a local variable. The Go compiler is intelligent enough
	and it will allocate this variable on the heap.
	 */
	ae := hello()
	fmt.Println("value of ae", *ae, ", address of ae", ae)
	/*
	In the function hello(), we return the address of the local variable i. The behavior of this code is undefined in
	programming languages such as C and C++ as the variable i goes out of scope once the function hello returns. But in
	the case of Go, the compiler does a escape analysis and allocates i on the heap as the address escapes the local
	scope.
	 */

	fmt.Printf("\nBeginning of DO NOT PASS a pointer to an array as a argument to a function, using slice instead\n")
	/*
	Lets assume that we want to make some modifications to an array inside the function and the changes made to that array
	inside the function should be visible to the caller. One way of doing this is to pass a pointer to an array as an
	argument to the function.
	*/
	simpleArray := [3]int{89, 90, 91}
	modifyArray(&simpleArray)
	fmt.Println(simpleArray)
	/*
	In the above program, we are passing the address of the array simpleArray to the modify function. In the modify
	function we are dereferencing arr and assigning 90 to the first element of the array. This program outputs [90 90 91]
	 */
	/*
	arr[x] is shorthand for (*arr)[x]. So (*arr)[0] in the above program can be replaced by arr[0].
	 */
	/*
	Although this way of passing a pointer to an array as a argument to a function and making modification to it works,
	it is not the idiomatic way of achieving this in Go. We have slices for this.
	 */
	simpleArray = [3]int{89, 90, 91}
	modifySlice(simpleArray[:])
	fmt.Println(simpleArray)
	/*
	In the program above, we pass a slice to the modify function. The first element of the slice is changed
	to 90 inside the modifySlice function. This program also outputs [90 90 91]. So forget about passing pointers to arrays
	around and use slices instead :). This code is much more clean and is idiomatic Go :).
	 */

	fmt.Printf("\nBeginning of Go does not support pointer arithmetic\n")
	/*
	Go does not support pointer arithmetic which is present in other languages like C and C++.
	 */
	//af := [...]int{109, 110, 111}
	//fa := &af
	//fa++
	/*
	The above program will throw compilation error main.go:6: invalid operation: p++ (non-numeric type *[3]int) if we
	uncomment 3 lines.
	 */
}