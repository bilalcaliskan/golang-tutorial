package constants

import(
	"fmt"
	"math"
)

func RunConstants() {
	fmt.Printf("\nBeginning of constants...\n")
	/*
	The term constant is used in Go to denote fixed values such 5, -89, "I love Go", 67.89 and so on.
	Constants as the name indicate cannot be reassigned again to any other value.
	*/
	var a int = 50
	var b string = "I love Go"
	fmt.Println("a=", a, " b=", b)
	/*
	In the above code a and b are assigned to constants 50 and I love Go respectively. The keyword const is used to
	denote constants such as 50 and I love Go. Even though we do not explicitly use the keyword const anywhere in the
	above code, internally they are constants in Go.
	 */
	const c = 55 //allowed
	// c = 30 // can not re assign constants
	/*
	Constants as the name indicate cannot be reassigned again to any other value and hence the above program will not
	work and it will fail with compilation error cannot assign to a.
	 */
	var d = math.Sqrt(4) //allowed
	// const e = math.Sqrt(4) //not allowed
	fmt.Println("d=", d)
	/*
	In the above program, a is a variable and hence it can be assigned to the result of the function math.Sqrt(4).
	b is a constant and the value of b needs to be know at compile time. The function math.Sqrt(4) will be evaluated
	only during run time and hence const b = math.Sqrt(4) throws error main.go:11: const initializer math.Sqrt(4) is
	not a constant if we uncomment constant decleration.
	 */

	fmt.Printf("\nBeginning of string constants...\n")
	/*
	Any value enclosed between double quotes is a string constant in Go. What type does a string constant belong to?
	The answer is they are untyped.
	 */
	// string constants like "Hello World" does not have a type
	const hello  = "Hello World"
	const hello1 = 12
	fmt.Printf("type %T value %v", hello, hello)
	fmt.Printf("type %T value %v", hello1, hello1)
	/*
	untyped constants have a default type associated with them and they supply it if and only if a line of code demands
	it. In the statement var name = "Sam", name needs a type and it gets it from the default type of the string constant
	"Sam" which is a string.
	 */
	const typedHello string = "Hello World"
	fmt.Printf("type %T value %v", typedHello, typedHello)
	/*
	Go is a strongly typed language. Mixing types during assignment is not allowed. Let's see what this means by the
	help of a program.
	 */
	var defaultName = "Sam" // untyped string
	type myString string
	var customName myString = "Sam" // typed myString
	//customName = defaultName // not allowed since customName and defaultName have different types
	/*
	Even though we know that myString is an alias of string, Go's strong typing policy disallows variables of one type
	to be assigned to another.
	*/
	fmt.Println(defaultName, customName)

	fmt.Printf("\nBeginning of boolean constants...\n")
	/*
	Boolean constants are no different from string constants. They are two untyped constants true and false.
	*/
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst //allowed
	var customBool myBool = trueConst //allowed
	// defaultBool = customBool // not allowed since defaultBool and customBool have different types
	/*
	Even though we know that myString is an alias of string, Go's strong typing policy disallows variables of one type
	to be assigned to another.
	*/
	fmt.Println(defaultBool, customBool)

	fmt.Printf("\nBeginning of numeric constants...\n")
	/*
	Numeric constants include integers, floats and complex constants. 
	 */
	const ab = 5
	var intVar int = ab
	var int32Var int32 = ab
	var float64Var float64 = ab
	var complex64Var complex64 = ab
	fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)
	/*
	In this program, the value of a is 5 and the syntax of a is generic (it can represent a float, integer or even a
	complex number with no imaginary part) and hence it is possible to be assigned to any compatible type. The default
	type of these kind of constants can be thought of as being generated on the fly depending on the context.
	var intVar int = a requires a to be int so it becomes an int constant. var complex64Var complex64 = a requires a to
	be a complex number and hence it becomes a complex constant.
	*/

	fmt.Printf("\nBeginning of numeric expressions...\n")
	/*
	Numeric constants are free to be mixed and matched in expressions and a type is needed only when they are assigned
	to variables or used in any place in code which demands a type.
	 */
	var abc = 5.9/8
	abc = 5.8/7
	fmt.Printf("a's type %T value %v",abc, abc)
}
