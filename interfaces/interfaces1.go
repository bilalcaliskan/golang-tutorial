package interfaces

import (
	"fmt"
)

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age int
}

func (s Student) Describe() {
	fmt.Printf("%s is %d years old", s.name, s.age)
}

func findDescriberType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a %T and my value is %s\n", i, i.(string))
	case int:
		fmt.Printf("I am a %T and my value is %d\n", i, i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func assertUnsafe(i interface{}) {
	v := i.(int)
	fmt.Println(v)
}

func assertSafe(i interface{}) {
	/*
	If the concrete type of i is T then v will have the underlying value of i and ok will be true.
	If the concrete type of i is not T then ok will be false and v will have the zero value of type T and the program
	will not panic.
	 */
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func describeEmptyInterface(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

type Worker interface {
	Work()
}

type Person struct {
	name string
}

func (p Person) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type permanent struct {
	empId int
	basicpay int
	pf int
}

type contract struct {
	empId int
	basicpay int
}

type freelancer struct {
	empId int
	ratePerHour int
	totalHours int
}

func (f freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func (p permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total expense per month $%d\n", expense)
}

func RunInterfaces1() {
	fmt.Printf("\nBeginning of introduction to interfaces...\n")
	/*
	In Go, an interface is a set of method signatures. When a type provides definition for all the methods in the
	interface, it is said to implement the interface. It is much similar to the OOP world. Interface specifies what
	methods a type should have and the type decides how to implement these methods.
	 */

	fmt.Printf("\nBeginning of declaring and implementing an interface...\n")
	/*
	Other languages like Java where a class has to explicitly state that it implements an interface using the implements
	keyword. This is not needed in Go and Go interfaces are implemented implicitly if a type contains all the methods
	declared in the interface.
	 */
	name := MyString("Sam Anderson")
	var v VowelsFinder
	/*
	Below line is possible since MyString implements the VowelsFinder interface.
	 */
	v = name
	fmt.Printf("Vowels are %c\n", v.FindVowels())

	fmt.Printf("\nBeginning of practical use of interface...\n")
	pemp1 := permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}

	pemp2 := permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}

	cemp1 := contract{
		empId:    3,
		basicpay: 3000,
	}

	freelancer1 := freelancer{
		empId:       4,
		ratePerHour: 100,
		totalHours:  100,
	}

	freelancer2 := freelancer{
		empId:       5,
		ratePerHour: 70,
		totalHours:  120,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
	totalExpense(employees)

	fmt.Printf("\nBeginning of interface internal representation...\n")
	/*
	An interface can be thought of as being represented internally by a tuple (type, value). type is the underlying
	concrete type of the interface and value holds the value of the concrete type.
	 */
	/*
	Worker interface has one method Work() and Person struct type implements that interface. In line no. 186, we assign
	the variable p of type Person to w which is of type Worker. Now the concrete type of w is Person and it contains a
	Person with name field Naveen. The describe function in line no.188 prints the value and concrete type of the interface.
	*/
	p := Person{name: "Naveen"}
	var w Worker = p
	describe(w)
	w.Work()

	fmt.Printf("\nBeginning of empty interface...\n")
	/*
	An interface that has zero methods is called an empty interface. It is represented as interface{}. Since the empty
	interface has zero methods, all types implement the empty interface.
	*/
	s := "Hello World"
	describeEmptyInterface(s)
	i := 55
	describeEmptyInterface(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describeEmptyInterface(strt)
	/*
	In the program above, the describeEmptyInterface(i interface{}) function takes an empty interface as an argument and
	hence any type can be passed.
	 */

	fmt.Printf("\nBeginning of type assertion...\n")
	/*
	Type assertion is used to extract the underlying value of the interface.
	i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T.
	*/
	var ss interface{} = 12
	assertUnsafe(ss)
	// The concrete type of ss in line no. 216 is int. This program prints 56.
	/*
	What will happen if the concrete type in the above program is not int?
	*/
	var sss interface{} = "Steven Paul"
	assertSafe(sss)
	/*
	In the program above we pass s of concrete type string to the assertSafe function which tries to extract a int value
	from it. If we used the assertUnsafe function, this program will panic with the message panic: interface
	conversion: interface {} is string, not int.
	Above syntax solves the problem in the assertSafe function:
		v, ok := i.(T)
	If the concrete type of i is T then v will have the underlying value of i and ok will be true.
	If the concrete type of i is not T then ok will be false and v will have the zero value of type T and the program
	will not panic.
	*/
	/*
	When Steven Paul is passed to the assertSafe function, ok will be false since the concrete type of i is not int and
	v will have the value 0 which is the zero value of int.
	 */

	fmt.Printf("\nBeginning of type switch...\n")
	/*
	A type switch is used to compare the concrete type of an interface against multiple types specified in various case
	statements. It is similar to switch case. The only difference being the cases specify types and not values as in
	normal switch.
	 */
	findType("Naveen")
	findType(77)
	findType(89.98)
	/*
	It is also possible to compare a type to an interface. If we have a type and if that type implements an interface,
	it is possible to compare this type with the interface it implements.
	 */
	findDescriberType("Naveen")
	student := Student{
		name: "Naveen R",
		age:  25,
	}
	findDescriberType(student)
}