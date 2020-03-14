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

func Run()  {
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

	fmt.Printf("\nBeginning of more on methods...\n")
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle is %d\n", r.Area())
	c := Circle{radius: 12}
	fmt.Printf("Area of circle is %f\n", c.Area())

	fmt.Printf("\nBeginning of pointer receivers vs value receivers...\n")
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
	fmt.Println("calling the method with pointer receiver with & is not needed.")
	emp3.changeSalary(15000)
	fmt.Printf("Employee salary after change: %d\n", emp3.salary)

	fmt.Printf("\nBeginning of methods of anonymous struct fields...\n")
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
	fmt.Println("When a function has a value argument, it will accept only a value argument.")
	fmt.Println("When a method has a value receiver, it will accept both pointer and value receivers.")
	rec := Rectangle{
		length: 10,
		width:  5,
	}
	area(rec)
	rec.area()
	recPointer := &rec
	recPointer.area()

	fmt.Printf("\nBeginning of pointer receivers in methods vs pointer arguments in functions...\n")
	rec2 := Rectangle{
		length: 10,
		width:  5,
	}
	recPointer2 := &rec2
	perimeter(recPointer2)
	recPointer2.perimeter()
	rec2.perimeter()

	fmt.Printf("\nBeginning of methods with non-struct receivers...\n")
	fmt.Println("To define a method on a type, the definition of the receiver type and the definition of the method " +
		"should be present in the same package. So far, all the structs and the methods on structs we defined were all " +
		"located in the same main package and hence they worked.")
	fmt.Println("we are trying to add a method named add on the built-in type int. This is not allowed since the " +
		"definition of the method add and the definition of type int is not in the same package.")
	fmt.Println("The way to get this working is to create a type alias for the built-in type int and then create a " +
		"method with this type alias as the receiver.")
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}