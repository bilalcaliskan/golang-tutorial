package interfaces

import "fmt"

type salaryCalculator interface {
	displaySalary()
}

type leaveCalculator interface {
	calculateLeavesLeft() int
}

type employeeOperations interface {
	salaryCalculator
	leaveCalculator
}

type employee struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesTaken int
}

func (e employee) displaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, e.basicPay + e.pf)
}

func (e employee) calculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

type describer interface {
	describe()
}

type person struct {
	name string
	age int
}

func (p person) describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type address struct {
	state string
	country string
}

func (a *address) describe() {
	fmt.Printf("State %s Country %s\n", a.state, a.country)
}

func RunInterfaces2() {
	fmt.Printf("\nBeginning of implementing interfaces using pointer receivers vs value receivers...\n")
	var d1 describer
	p1 := person{"Sam", 25}
	d1 = p1
	d1.describe()
	p2 := person{"James", 32}
	d1 = &p2
	d1.describe()

	var d2 describer
	a := address{"Washington", "USA"}
	d2 = &a
	d2.describe()
	if d2 == nil {
		fmt.Printf("d2 is nil and has type %T value %v\n", d2, d2)
	} else {
		fmt.Printf("d2 is not nil and has type %T value %v\n", d2, d2)
	}

	fmt.Printf("\nBeginning of implementing multiple interfaces...\n")
	e := employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s salaryCalculator = e
	s.displaySalary()
	var l leaveCalculator = e
	fmt.Printf("%d leaves left", l.calculateLeavesLeft())

	fmt.Printf("\nBeginning of embedded interfaces...\n")
	fmt.Println("Although go does not offer inheritance, it is possible to create a new interfaces by embedding " +
		"other interfaces.")
	e2 := employee{"Naveen", "Ramanathan", 5000, 200, 30, 5}
	var empOp employeeOperations = e2
	empOp.displaySalary()
	fmt.Printf("%d leaves left", empOp.calculateLeavesLeft())

	fmt.Printf("\nBeginning of zero value of interface...\n")
	var nilInterface describer
	if nilInterface == nil {
		fmt.Printf("nilInterface is nil and has type %T value %v\n", nilInterface, nilInterface)
	}
	nilInterface.describe()
}