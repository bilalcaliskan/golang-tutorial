package structures

import "fmt"

// struct with anonymous fields
type employeeWithSalary struct {
	firstName, lastName string
	age, salary	int
}

/*
The above employee struct is called a named structure because it creates a new type named employee which can be used to
create structures of type employee.
*/
type employee struct {
	firstName string
	lastName  string
	age       int
}

type personWithAnonymousFields struct {
	string
	int
}

type address struct {
	city, state string
}

type promotedPerson struct {
	name string
	age int
	address
}

type person struct {
	name    string
	age     int
	address address
}

type name struct {
	firstName, lastName string
}

type image struct {
	data map[int] int
}

func RunStructures() {
	fmt.Printf("\nBeginning of introduction to structures...\n")
	/*
	A structure is a user defined type which represents a collection of fields. It can be used in places where it makes
	sense to group the data into a single unit rather than maintaining each of them as separate types.
	For instance a employee has a firstName, lastName and age. It makes sense to group these three properties into a
	single structure employee.
	 */

	fmt.Printf("\nBeginning of creating named structures...\n")
	/*
	We have created 2 different named structs employee and employeeWithSalary. Here we are creating different
	structs of employee and employeeWithSalary struct types.
	 */
	// creating structure with using field names
	emp0 := employee{
		firstName: "Bilal",
		lastName:  "Caliskan",
		age:       26, // comma is needed here
	}
	// creating structure with using field names
	emp1 := employeeWithSalary{firstName: "Sam", lastName: "Anderson", age: 25, salary: 500}
	// creating structure with using field names
	emp2 := employeeWithSalary{
		firstName: "Josh",
		lastName: "Sagredo",
		age: 25,
		salary: 1000, // comma is needed here
	}
	// creating structure without using field names
	emp3 := employeeWithSalary{"Sam", "Sagredo", 24, 4000}
	fmt.Println("employee 0", emp0)
	fmt.Println("employee 1", emp1)
	fmt.Println("employee 2", emp2)
	fmt.Println("employee 3", emp3)

	fmt.Printf("\nBeginning of creating anonymous structures...\n")
	emp4 := struct {
		firstName, lastName string
		age, salary			int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("employee 4", emp4)

	fmt.Printf("\nBeginning of zero value structures...\n")
	/*
	When a struct is defined and it is not explicitly initialised with any value, the fields of the struct are assigned
	their zero values by default.
	 */
	var emp5 employee
	fmt.Println("employee 5", emp5)
	emp6 := employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("employee 6", emp6)
	var emp7 employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("employee 7", emp7)

	fmt.Printf("\nBeginning of accessing individual fields of a struct...\n")
	/*
	The dot . operator is used to access the individual fields of a structure.
	 */
	emp60 := employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
	}
	fmt.Println("First Name:", emp60.firstName)
	fmt.Println("Last Name:", emp60.lastName)
	fmt.Println("Age:", emp60.age)

	fmt.Printf("\nBeginning of pointers to a struct...\n")
	/*
	It is also possible to create pointers to a struct.
	 */
	emp8 := &employeeWithSalary{
		firstName: "Sam",
		lastName:  "Alderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println("Pointer value of employee 8", *emp8) // dereferencing emp8 to get struct values
	fmt.Println("Pointer value of employee 8 name", (*emp8).firstName) // dereferencing emp8 to get firstName
	/*
	The language gives us the option to use emp8.firstName instead of the explicit dereference (*emp8).firstName to
	access the firstName field.
	 */
	fmt.Println("Pointer value of employee 8 name:", emp8.firstName)
	fmt.Println("Pointer value of employee 8 age:", emp8.age)

	fmt.Printf("\nBeginning anonymous fields...\n")
	/*
	It is possible to create structs with fields which contain only a type without the field name. These kind of fields
	are called anonymous fields. Check personWithAnonymousFields struct type above
	 */
	per1 := personWithAnonymousFields{
		string: "Naveen",
		int:    50,
	}
	per2 := personWithAnonymousFields{"Hasan", 100}
	fmt.Println("person 1", per1)
	fmt.Println("person 2", per2)
	/*
	Even though an anonymous fields does not have a name, by default the name of a anonymous field is the name of its type.
	per2.string will give you the name of per2
	*/
	var per0 personWithAnonymousFields
	per0.string = "naveen"
	per0.int = 50
	fmt.Println(per0)

	fmt.Printf("\nBeginning of nested structs...\n")
	/*
	It is possible that a struct contains a field which in turn is a struct. These kind of structs are called as nested
	structs. Check the person struct type above to see what is going on.
	 */
	var p person
	p.name = "Naveen"
	p.age = 50
	p.address = address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:",p.age)
	fmt.Println("City:",p.address.city)
	fmt.Println("State:",p.address.state)
	/*
	The person struct in the above program has a field address which in turn is a struct.
	 */

	fmt.Printf("\nBeginning of promoted fields...\n")
	/*
	Fields that belong to a anonymous struct field in a structure are called promoted fields since they can be accessed
	as if they belong to the structure which holds the anonymous struct field.
	*/
	var per4 promotedPerson
	per4.name = "Naveen"
	per4.age = 40
	per4.address = address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("person 4 name", per4.name)
	fmt.Println("person 4 age", per4.age)
	fmt.Println("person 4 city", per4.city)   // city is promoted field
	fmt.Println("person 4 state", per4.state) // state is promoted field

	fmt.Printf("\nBeginning of exported structs and fields...\n")
	/*
	If a struct type starts with a capital letter, then it is a exported type and it can be accessed from other packages.
	Similarly if the fields of a structure start with caps, they can be accessed from other packages. But of course
	first struct itself must be reachable to reach fields
	 */

	fmt.Printf("\nBeginning of structures equality...\n")
	/*
	Structs are value types and are comparable if each of their fields are comparable. Two struct variables are
	considered equal if their corresponding fields are equal.
	 */
	name1 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}
	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
	/*
	In the above program, name struct type contain two string fields. Since strings are comparable, it is possible to
	compare two struct variables of type name.
	 */
	/*
	Struct variables are not comparable if they contain fields which are not comparable.
	 */
	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}
	if image1 == image2 {
		fmt.Println("image1 and image2 are equal")
	} else {
		fmt.Println("image1 and image2 not equal")
	}
	/*
	In the program above image struct type contains a field data which is of type map. maps are not comparable, hence
	image1 and image2 cannot be compared. If you run this program, compilation will fail with error main.go:18: invalid
	operation: image1 == image2 (struct containing map[int]int cannot be compared).
	 */
}