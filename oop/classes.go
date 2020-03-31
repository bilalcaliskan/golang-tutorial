package oop

import(
	"fmt"
	"golang-tutorial/oop/employee"
)

func RunClasses() {
	fmt.Printf("\nBeginning of is Go object oriented?\n")
	/*
		Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type
		hierarchy. The concept of “interface” in Go provides a different approach that we believe is easy to use and in
		some ways more general. There are also ways to embed types in other types to provide something analogous—but not
		identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for
		any sort of data, even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes).
	*/

	fmt.Printf("\nBeginning of structs instead of classes...\n")
	/*
	Go does not provide classes but it does provide structs. Methods can be added on structs. This provides the behaviour
	of bundling the data and methods that operate on the data together akin to a class.
	 */
	e := employee.Employee {
		FirstName: "Sam",
		LastName: "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()

	fmt.Printf("\nBeginning of New() function instead of constructors...\n")
	var e1 employee.Employee
	e1.LeavesRemaining()
	/*
	As you can see, the variable created with the zero value of Employee is unusable. It doesn't have a valid first name,
	last name and also doesn't have valid leave details.
	In other OOP languages like java, this problem can be solved by using constructors. A valid object can be created
	by using parameterised constructor.

	Go doesn't support constructors. If the zero value of a type is not usable, it is the job of the programmer to
	unexport the type to prevent access from other packages and also to provide a function named NewT(parameters) which
	initialises the type T with the required values. It is a convention in Go to name a function which creates a value
	of type T to NewT(parameters). This will act like a constructor. If the package defines only one type, then it's a
	convention in Go to name this function just New(parameters) instead of NewT(parameters).
	 */
	e2 := employee.NewEmployee2("Sam", "Adolf", 30, 20)
	e2.LeavesRemaining()
	/*
	Thus you can understand that although Go doesn't support classes, structs can effectively be used instead of classes
	and methods of signature New(parameters) can be used in the place of constructors.
	 */
}