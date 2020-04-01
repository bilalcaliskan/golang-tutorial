package first_class

import "fmt"

// user defined function type
type add func(a int, b int) int

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

// this function returns another function(func(a, b int) int)
func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

type student struct {
	firstName string
	lastName string
	grade string
	country string
}

func filterStudents(allStudents []student, filterFunction func(student) bool) []student {
	var filteredStudents []student
	for _, v := range allStudents {
		if filterFunction(v) {
			filteredStudents = append(filteredStudents, v)
		}
	}
	return filteredStudents
}

func iMap(integerSlice []int, multiplierFunction func(int) int) []int {
	var returnedSlice []int
	for _, v := range integerSlice {
		returnedSlice = append(returnedSlice, multiplierFunction(v))
	}
	return returnedSlice
}

func RunFirstClassFunctions() {
	fmt.Printf("\nBeginning of introduction to first class functions...\n")
	/*
	A language which supports first class functions allows functions to be assigned to variables, passed as arguments
	to other functions and returned from other functions. Go has support for first class functions.
	*/
	fmt.Printf("\nBeginning of anonymous functions...\n")
	/*
	Let's start with a simple example which assigns a function to a variable.
	These kind of functions like below are called anonymous functions since they do not have a name.
	*/
	a := func() {
		fmt.Println("hello world first first class function")
	}
	a()
	fmt.Printf("%T\n", a)
	/*
	The only way to call this function is using the variable a.
	It is also possible to call a anonymous function without assigning it to a variable.
	*/
	func() {
		fmt.Println("hello world second first class function")
	}()
	// It is also possible to pass arguments to anonymous functions just like any other function.
	func(n string){
		fmt.Println("Welcome", n)
	}("Gophers")

	fmt.Printf("\nBeginning of user defined function types...\n")
	/*
	Just like we define our own struct types, it is possible to define our own function types:
		type add func(a int, b int) int
	The code snippet above creates a new function type add which accepts two integer arguments and returns a integer.
	Now we can define variables of type add.
	 */
	var b add = func(a int, b int) int {
		return a + b
	}
	s := b(5, 6)
	fmt.Println("Sum", s)

	fmt.Printf("\nBeginning of higher-order functions...\n")
	/*
	The definition of Higher-order function from wiki is a function which does at least one of the following:
		1- takes one or more functions as arguments
		2- returns a function as its result
	 */
	/*
	1- Passing functions as arguments to other functions
	In the below example, in line no. 8 we define a function simple which takes a function which accepts two int
	arguments and returns a int as a parameter. Inside the main function in line no. 105 we create a anonymous function
	f whose signature matches the parameter of the function simple. We call simple and pass f as an argument to it in
	the next line. This program prints 67 as output.
	*/
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
	/*
	2- Returning functions from other functions
	Now let's rewrite the program above and return a function from the simple2 function.
	*/
	d := simple2()
	fmt.Println(d(60, 7))

	fmt.Printf("\nBeginning of Closures...\n")
	/*
	Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables
	defined outside the body of the function.
	*/
	e := 5
	// closures are special case of anonymous functions.
	func() {
		fmt.Println("e = ", e)
	}()

	fmt.Printf("\nBeginning of Closure example...\n")
	/*
	Every closure is bound to its own surrounding variable. Let's understand what this means by using a simple example.
	*/
	j := appendStr()
	k := appendStr()
	fmt.Println(j("World"))
	fmt.Println(k("Everyone"))
	fmt.Println(j("Gopher"))
	fmt.Println(k("!"))
	/*
	In the program above, the function appendStr returns a closure. This closure is bound to the variable t. Let's
	understand what this means.
	The variables j and k declared in line nos. 134, 135 are closures and they are bound to their own value of t.
	We first call j with the parameter World. Now the value of j's version of t becomes Hello World.
	In line no. 137 we call k with the parameter Everyone. Since k is bound to its own variable t, k's version of t has
	a initial value of Hello again. Hence after this function call, the value of k's version of t becomes Hello Everyone
	Key note about closures is that Closures are anonymous functions which access the variables defined outside the
	body of the function.
	 */

	fmt.Printf("\nBeginning of practical use of first class functions example 1...\n")
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	allStudents := []student{s1, s2}
	filteredStudents := filterStudents(allStudents, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(filteredStudents)

	fmt.Printf("\nBeginning of practical use of first class functions example 2...\n")
	/*
	This program will perform the same operations on each element of a slice and return the result. For example if
	we want to multiply all integers in a slice by 5 and return the output, it can be easily done using first class
	functions. These kind of functions which operate on every element of a collection are called map functions. I
	have provided the program below.
	 */
	// These kind of functions which operate on every element of a collection are called map functions
	integerSlice := []int{10, 20, 30, 40, 50}
	returnedSlice := iMap(integerSlice, func(i int) int {
		return i * 5
	})
	fmt.Println(returnedSlice)
}