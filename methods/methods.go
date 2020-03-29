package methods

import (
	"fmt"
	"math"
)

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func area(r Rectangle) {
	fmt.Printf("Area function result: %d\n", r.length * r.width)
}

/*
This method can accept both pointer receiver and value receiver.
 */
func (r Rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func (r *Rectangle) perimeter() {
	fmt.Println("perimeter method output: ", 2 * (r.length + r.width))
}

func perimeter(r *Rectangle) {
	fmt.Println("perimeter function output:", 2 * (r.length + r.width))
}

type address struct {
	city string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

type person struct {
	firstName string
	lastName string
	address // anonymous field
}

type Employee struct {
	name string
	salary int
	currency string
}

type Rectangle struct {
	length int
	width int
}

type Circle struct {
	radius float64
}

// method with the int return type
func (r Rectangle) Area() int {
	return r.length * r.width
}

// method with the float64 return type
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// void method, does not return anything
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

// void function, does not return anything
func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

// method with value receiver
func (e Employee) changeName(newName string) {
	e.name = newName
}

// method with pointer receiver
func (e *Employee) changeSalary(newSalary int) {
	e.salary = newSalary
}

func RunMethods()  {
	fmt.Printf("\nBeginning of introduction to methods...\n")
	/*
	A method is just a function with a special receiver type between the func keyword and the method name. The receiver
	can either be a struct type or non-struct type.
	The syntax of a method declaration is provided below:
		func (t Type) methodName(parameter list) {
		}
	The above snippet creates a method named methodName with receiver type Type. t is called as the receiver and it can
	be accessed within the method.
	*/

	fmt.Printf("\nBeginning of sample methods...\n")
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary()

	fmt.Printf("\nBeginning of sample functions...\n")
	emp2 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	displaySalary(emp2)

	fmt.Printf("\nBeginning of methods vs functions...\n")
	/*
	So why do we have methods when we can write the same program using functions. There are a couple of reasons for this.
	Let's look at them one by one:
		- Go is not a pure object-oriented programming language and it does not support classes. Hence methods on types
		are a way to achieve behavior similar to classes. Methods allow a logical grouping of behavior related to a type
		similar to classes. In the above sample program, all behaviors related to the Employee type can be grouped by
		creating methods using Employee receiver type. For example, we can add methods like calculatePension, calculateLeaves
		and so on.
		- Methods with the same name can be defined on different types whereas functions with the same names are not allowed.
		Let's assume that we have a Square and Circle structure. It's possible to define a method named Area on both Square
		and Circle.
	 */

	fmt.Printf("\nBeginning of more on methods...\n")
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle is %d\n", r.Area())
	c := Circle{radius: 12}
	fmt.Printf("Area of circle is %f\n", c.Area())

	fmt.Printf("\nBeginning of pointer receivers vs value receivers...\n")
	/*
	So far we have seen methods only with value receivers. It is possible to create methods with pointer receivers. The
	difference between value and pointer receiver is, changes made inside a method with a pointer receiver is visible to
	the caller whereas this is not the case in value receiver.
	*/
	emp3 := Employee{
		name:     "Mark Andrew",
		salary:   5000,
		currency: "$",
	}
	fmt.Printf("Employee name before change: %s\n", emp3.name)
	emp3.changeName("Michael Andrew")
	fmt.Printf("Employee name after change: %s\n", emp3.name)

	fmt.Printf("Employee salary before change: %d\n", emp3.salary)
	(&emp3).changeSalary(10000)
	fmt.Printf("Employee salary after change: %d\n", emp3.salary)
	/*
	calling the method with pointer receiver with & is not needed. Same functionality can be done with below line.
	 */
	emp3.changeSalary(15000)
	fmt.Printf("Employee salary after change: %d\n", emp3.salary)

	fmt.Printf("\nBeginning of when to use pointer receiver and when to use value receiver...\n")
	/*
	Generally, pointer receivers can be used when changes made to the receiver inside the method should be visible to
	the caller.
	Pointers receivers can also be used in places where it's expensive to copy a data structure. Consider a struct that
	has many fields. Using this struct as a value receiver in a method will need the entire struct to be copied which
	will be expensive. In this case, if a pointer receiver is used, the struct will not be copied and only a pointer to
	it will be used in the method.
	In all other situations, value receivers can be used.
	*/

	fmt.Printf("\nBeginning of methods of anonymous struct fields...\n")
	/*
	Methods belonging to anonymous fields of a struct can be called as if they belong to the structure where the
	anonymous field is defined. This behavior looks like promoted fields.
	 */
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address:   address{
			city: "Los Angeles",
			state: "California",
		},
	}
	p.fullAddress()

	fmt.Printf("\nBeginning of value receivers in methods vs value arguments in functions...\n")
	/*
	When a function has a value argument, it will accept only a value argument.
	When a method has a value receiver, it will accept both pointer and value receivers.
	The reason is that the line p.area(), for convenience will be interpreted by Go as (*p).area() since area has a
	value receiver.
	*/
	rec := Rectangle{
		length: 10,
		width:  5,
	}
	area(rec)
	rec.area()
	recPointer := &rec
	recPointer.area() // this line, for convenience will be interpreted by Go as (*recPointer).area() since area has a value receiver.

	fmt.Printf("\nBeginning of pointer receivers in methods vs pointer arguments in functions...\n")
	/*
	Similar to value arguments, functions with pointer arguments will accept only pointers whereas methods with pointer
	receivers will accept both pointer and value receiver.
	 */
	rec2 := Rectangle{
		length: 10,
		width:  5,
	}
	recPointer2 := &rec2
	perimeter(recPointer2)
	recPointer2.perimeter()
	rec2.perimeter() // this line will be interpreted by the language as (&recPointer2).perimeter() for convenience.

	fmt.Printf("\nBeginning of methods with non-struct receivers...\n")
	/*
	So far we have defined methods only on struct types. It is also possible to define methods on non-struct types, but
	there is a catch. To define a method on a type, the definition of the receiver type and the definition of the method
	should be present in the same package. So far, all the structures and the methods on structures we defined were all
	located in the same main package and hence they worked.
	 */
	/*
	To achieve required challenge, we have created a type alias myInt for int. Then we have defined a method add with
	myInt as the receiver.
	*/
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}