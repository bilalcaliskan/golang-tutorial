package structs

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type Person struct {
	string
	int
}

type Address struct {
	city, state string
}

type Student struct {
	name string
	age int
	address Address
}

type PromotedAddress struct {
	city, state string
}

type Worker struct {
	name string
	age int
	PromotedAddress
}

type Spec struct { //exported struct
	Maker string //exported field
	model string //unexported field
	Price int //exported field
}

type name struct { //unexported struct
	firstName string
	lastName string
}

type image struct {
	data map[int]int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

type Vertex struct {
	X int
	Y int
}

func Run() {
	emp1 := Employee {
		firstName: "asdfasf",
		lastName:  "asdf",
		age:       0,
		salary:    0,
	}
	emp2 := Employee {"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	var emp4 Employee //zero valued structure
	fmt.Println("Employee 4", emp4)

	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)

	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Jones"
	fmt.Println("Employee 7:", emp7)

	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)

	per1 := Person{"Naveen", 50}
	fmt.Println(per1)

	var per2 Person
	per2.string = "naveen"
	per2.int = 50
	fmt.Println(per2)

	var stu1 Student
	stu1.name = "Naveen"
	stu1.age = 50
	stu1.address = Address {
		city: "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", stu1.name)
	fmt.Println("Age:", stu1.age)
	fmt.Println("City:", stu1.address.city)
	fmt.Println("State:", stu1.address.state)

	var p Worker
	p.name = "Naveen"
	p.age = 50
	p.PromotedAddress = PromotedAddress {
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city) //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field

	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}

	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName:"Steve", lastName:"Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	fmt.Println(v1, p, v2, v3)
}