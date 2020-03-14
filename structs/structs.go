package structs

import "fmt"

// anonymous struct
type Employee struct {
	firstName, lastName string
	age, salary	int
}

type PersonWithAnonymousFields struct {
	string
	int
}

type Address struct {
	city, state string
}

type PromotedPerson struct {
	name string
	age int
	Address
}

// anonymous struct
type Person struct {
	name string
	age int
	address Address
}

func Run() {
	// Basic struct operations
	fmt.Printf("\nBeginning basic struct operations...\n")
	emp1 := Employee{firstName: "Sam", lastName: "Anderson", age: 25, salary: 500}
	emp2 := Employee{
		firstName: "Josh",
		lastName: "Sagredo",
		age: 25,
		salary: 1000,
	}
	emp3 := Employee{"Sam", "Sagredo", 24, 4000}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
	fmt.Println("Employee 3", emp3)

	// Anonymous structures
	fmt.Printf("\nBeginning anonymous structures...\n")
	emp4 := struct {
		firstName, lastName string
		age, salary			int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 4", emp4)

	// Zero value of a structure
	fmt.Printf("\nBeginning of zero value structures...\n")
	var emp5 Employee
	fmt.Println("Employee 5", emp5)
	emp6 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("Employee 6", emp6)
	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7", emp7)

	// Pointers to a struct
	fmt.Printf("\nBeginning pointers to a struct...\n")
	emp8 := &Employee{
		firstName: "Sam",
		lastName:  "Alderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println("Pointer value of Employee 8", (*emp8))
	fmt.Println("Pointer value of Employee 8 name", (*emp8).firstName)

	// Anonymous fields
	fmt.Printf("\nBeginning anonymous fields...\n")
	per1 := PersonWithAnonymousFields{
		string: "Naveen",
		int:    50,
	}
	per2 := PersonWithAnonymousFields{"Hasan", 100}
	fmt.Println("Person 1", per1)
	fmt.Println("Person 2", per2)

	// Nested structs
	fmt.Printf("\nBeginning anonymous fields...\n")
	var per3 Person
	per3.name = "Naveen"
	per3.age = 50
	per3.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Person 3", per3)

	// Promoted fields
	fmt.Printf("\nBeginning of promoted fields...\n")
	var per4 PromotedPerson
	per4.name = "Naveen"
	per4.age = 40
	per4.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Person 4 name", per4.name)
	fmt.Println("Person 4 age", per4.age)
	fmt.Println("Person 4 city", per4.Address.city) // city is promoted field
	fmt.Println("Person 4 state", per4.Address.state) // state is promoted field

	// Exported Structs and Fields
	fmt.Printf("\nBeginning of exported structs...\n")
	

}