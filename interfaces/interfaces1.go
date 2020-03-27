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

func assert(i interface{}) {
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

type Permanent struct {
	empId int
	basicpay int
	pf int
}

type Contract struct {
	empId int
	basicpay int
}

type Freelancer struct {
	empId int
	ratePerHour int
	totalHours int
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total expense per month $%d", expense)
}

func RunInterfaces1() {
	fmt.Printf("\nBeginning of declaring and implementing an interface...\n")
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name
	fmt.Printf("Vowels are %c\n", v.FindVowels())

	fmt.Printf("\nBeginning of practical use of interface...\n")
	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}

	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}

	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
	}

	freelancer1 := Freelancer{
		empId:       4,
		ratePerHour: 100,
		totalHours:  100,
	}

	freelancer2 := Freelancer{
		empId:       5,
		ratePerHour: 70,
		totalHours:  120,
	}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
	totalExpense(employees)

	fmt.Printf("\nBeginning of interface internal representation...\n")
	p := Person{name: "Naveen"}
	var w Worker = p
	describe(w)
	w.Work()

	fmt.Printf("\nBeginning of empty interface...\n")
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

	fmt.Printf("\nBeginning of type assertion...\n")
	var ss interface{} = 12
	assert(ss)

	fmt.Printf("\nBeginning of type switch_statement...\n")
	findType("Naveen")
	findType(77)
	findType(89.98)
	findDescriberType("Naveen")
	student := Student{
		name: "Naveen R",
		age:  25,
	}
	findDescriberType(student)
}