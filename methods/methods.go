package methods

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

func (e Employee) displaySalaryMethod() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func displaySalaryFunction(e Employee) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func Run()  {
	emp1 := Employee {
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalaryMethod() //Calling displaySalary() method of Employee type
	//displaySalaryFunction(emp1)

	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address {
			city:  "Los Angeles",
			state: "California",
		},
	}
	p.fullAddress() //accessing fullAddress method of address struct
}