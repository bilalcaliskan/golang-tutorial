package variables

import "fmt"

func RunVariables() {
	fmt.Printf("\nBeginning of declaring a single variable...\n")
	/*
	If a variable is not assigned any value, go automatically initialises it with the zero value of the variable's type.
	In this case, age is assigned the value 0
	 */
	var age int
	fmt.Println("my age is", age)
	var age2 int // variable declaration
	fmt.Println("my age is ", age2)
	age2 = 29 //assignment
	fmt.Println("my age is", age2)
	age2 = 54 //assignment
	fmt.Println("my new age is", age2)

	fmt.Printf("\nBeginning of declaring a variable with initial value...\n")
	/*
	A variable can also be given a initial value when it is declared.
	 */
	var age3 int = 29 // variable declaration with initial value
	fmt.Println("my age is", age3)

	fmt.Printf("\nBeginning of type inference...\n")
	/*
	- If a variable has an initial value, Go will automatically be able to infer the type of that variable using that
	initial value. Hence if a variable has an initial value, the type in the variable declaration can be omitted.
	- If the variable is declared using the syntax var name = initialvalue, Go will automatically infer the type of that
	variable from the initial value.
	 */
	var age4 = 29 // type will be inferred
	fmt.Println("my age is", age4)

	fmt.Printf("\nBeginning of multiple variable declaration...\n")
	/*
	Multiple variables can be declared in a single statement.
	var name1, name2 type = initialvalue1, initialvalue2 is the syntax for multiple variable declaration.
	 */
	var width, height int = 100, 50 //declaring multiple variables
	fmt.Println("width is", width, "height is", height)
	/*
	The type can be omitted if the variables have initial value. The program below declares multiple variables using
	type inference.
	 */
	var width2, height2 = 100, 50 //"int" is dropped
	fmt.Println("width is", width2, "height is", height2)
	/*
	There might be cases where we would want to declare variables belonging to different types in a single statement.
	The syntax for doing that is
	*/
	var (
		name   = "naveen"
		age5    = 29
		height3 int // value will be zero value of type
	)
	fmt.Println("my name is", name, ", age is", age5, "and height is", height3)

	fmt.Printf("\nBeginning of short hand declaration...\n")
	/*
	- Go also provides another concise way for declaring variables. This is known as short hand declaration and it
	uses := operator. Short hand declaration requires initial values for all variables in the left hand side of
	the assignment.
	- Short hand syntax can only be used when at least one of the variables in the left side of := is newly declared. You
	can not re initialize the same variable in most of programming languages like go. If you do that, go will throw that
	error: "no new variables on left side of :="
	- Since Go is strongly typed, variables declared as belonging to one type cannot be assigned a value of another type.
	You cannot declare a variable with int value and then assign string value to it.
	*/
}