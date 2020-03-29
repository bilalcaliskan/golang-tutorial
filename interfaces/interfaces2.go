package interfaces

import "fmt"

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

type employee struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesTaken int
}

type address struct {
	state string
	country string
}

//https://de.spankbang.com/3ukvk/video/gata
func (a *address) describe() {
	fmt.Printf("State %s Country %s\n", a.state, a.country)
}

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

func (e employee) displaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, e.basicPay + e.pf)
}

func (e employee) calculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func RunInterfaces2() {
	fmt.Printf("\nBeginning of implementing interfaces using pointer receivers vs value receivers...\n")
	/*
	All the example interfaces we discussed in part 1 were implemented using value receivers. It is also possible to
	implement interfaces using pointer receivers. There is a subtlety to be noted while implementing interfaces using
	pointer receivers.
	 */
	var d1 describer
	p1 := person{"Sam", 25}
	d1 = p1
	d1.describe()
	fmt.Printf("p1 has type %T value %v\n", p1, p1)
	p2 := person{"James", 32}
	d1 = &p2
	d1.describe()
	fmt.Printf("p2 has type %T value %v\n", p2, p2)
	/*
	As we have already learnt during our discussion about methods, methods with value receivers accept both pointer and
	value receivers. It is legal to call a value method on anything which is a value or whose value can be dereferenced.
	p1 is a value of type Person and it is assigned to d1 in line no. 29. Person implements the Describer interface and
	hence line no. 30 will print Sam is 25 years old.
	Similarly d1 is assigned to &p2 in line no. 32 and hence line no. 33 will print James is 32 years old. Awesome :).
	 */
	var d2 describer
	a := address{"Washington", "USA"}
	/*
	The Address struct implements the Describer interface using pointer receiver.
	In the below code, we can not assign d2 to a. Because This is because, the Describer interface is implemented using
	a Address Pointer receiver in line 22 and we are trying to assign a which is a value type and it has not implemented
	the Describer interface. The reason is The concrete value stored in an interface is not addressable and hence it is
	not possible for the compiler to automatically take the address of a.
	*/
	// d2 = a
	d2 = &a
	d2.describe()
	if d2 == nil {
		fmt.Printf("d2 is nil and has type %T value %v\n", d2, d2)
	} else {
		fmt.Printf("d2 is not nil and has type %T value %v\n", d2, d2)
	}

	fmt.Printf("\nBeginning of implementing multiple interfaces...\n")
	/*
	A type can implement more than one interface. employee structure both implements salaryCalculator and leaveCAlculator
	interfaces.
	 */
	e := employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	e.displaySalary()
	/*
	Below declaration is possible since e which of type Employee implements SalaryCalculator interface.
	 */
	var s salaryCalculator = e
	s.displaySalary()
	fmt.Printf("%d leaves left\n", e.calculateLeavesLeft())
	/*
	Below declaration is possible since e which of type Employee implements LeaveCalculator interface.
	*/
	var l leaveCalculator = e
	fmt.Printf("%d leaves left\n", l.calculateLeavesLeft())

	fmt.Printf("\nBeginning of embedded interfaces...\n")
	/*
	Although go does not offer inheritance, it is possible to create a new interfaces by embedding other interfaces.
	 */
	e2 := employee{"Naveen", "Ramanathan", 5000, 200, 30, 5}
	var empOp employeeOperations = e2
	/*
	Below declaration is possible since e2 which of type Employee implements EmployeeOperations interface which includes
	both SalaryCalculator and LeaveCalculator interfaces.
	*/
	empOp.displaySalary()
	fmt.Printf("%d leaves left", empOp.calculateLeavesLeft())
	/*
	EmployeeOperations interface of the program above is created by embedding SalaryCalculator and LeaveCalculator interfaces.
	Any type is said to implement EmployeeOperations interface if it provides method definitions for the methods present
	in both SalaryCalculator and LeaveCalculator interfaces.
	The Employee struct implements EmployeeOperations interface since it provides definition for both DisplaySalary and
	CalculateLeavesLeft methods.
	 */

	fmt.Printf("\nBeginning of zero value of interface...\n")
	/*
	The zero value of a interface is nil. A nil interface has both its underlying value and as well as concrete type as nil.
	If we try to call a method on the nil interface, the program will panic since the nil interface neither has a underlying
	value nor a concrete type.
	*/
	var nilInterface describer
	/*
	If we uncomment the below line, program will panic like that:
	panic: runtime error: invalid memory address or nil pointer dereference
	*/
	// nilInterface.describe()
	if nilInterface == nil {
		fmt.Printf("nilInterface is nil and has type %T value %v\n", nilInterface, nilInterface)
	} else {
		nilInterface.describe()
	}
}