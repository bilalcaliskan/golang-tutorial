package reflection

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name string
	id int
	address string
	salary int
	country string
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
}

func createQueryKind(q interface{}) {
	typeOf := reflect.TypeOf(q)
	valueOf := reflect.ValueOf(q)
	tKind := typeOf.Kind()
	vKind := valueOf.Kind()
	fmt.Println("Type ", typeOf)
	fmt.Println("Value ", valueOf)
	fmt.Println("typeOf kind ", tKind)
	fmt.Println("valueOf kind ", vKind)
}

func createQueryNum(q interface{}) {
	//  We first check whether the Kind of q is a struct because the NumField method works only on struct.
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func createQueryComplete(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s(ordId) values(", t)
		if t == "order" {
			query = fmt.Sprintf("insert into %s(ordId, customerId) values(", t)
		} else if t == "employee" {
			query = fmt.Sprintf("insert into %s(empName, empId, empTown, empSalary, empCountry) values(", t)
		}
		v := reflect.ValueOf(q)
		fmt.Println("valueOf = ", reflect.ValueOf(q))
		fmt.Println("typeOf = ", reflect.TypeOf(q))
		fmt.Println("numField = ", reflect.ValueOf(q).NumField())
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

func RunReflection() {
	fmt.Printf("\nBeginning of introduction to reflection...\n")
	/*
	Reflection is the ability of a program to inspect its variables and values at run time and find their type.
	The first question anyone gets when learning about reflection is why do we even need to inspect a variable and
	find its type at runtime when each and every variable in our program is defined by us and we know its type at
	compile time itself. The data that your application process will be different at run time, that is the quickest
	answer.
	*/
	i := 10
	fmt.Printf("%d %T", i, i)
	// In the above program the type of i is known at compile time but we need to understand it at the runtime.

	fmt.Printf("\nBeginning of why reflection...\n")
	/*
	Assume that you have a struct type order, and another type employee. You want a single function createQuery() depending
	on the different struct types and generating different queries. As final, your createQuery() function should work with
	any struct. The only way to write this function is to examine the type of the struct argument passed to it at run time,
	find its fields and then create the query. This is where reflection is useful.
	 */

	fmt.Printf("\nBeginning of reflect package...\n")
	/*
	The reflect package implements run-time reflection in Go. The reflect package helps to identify the underlying
	concrete type and the value of a interface{} variable. This is exactly what we need. The createQuery function
	takes a interface{} argument and the query needs to be created based on the concrete type and value of the
	interface{} argument. This is exactly what the reflect package helps in doing.
	There are a few types and methods in the reflect package which we need to know first:

	1- reflect.Type and reflect.Value:
	The concrete type of interface{} is represented by reflect.Type and the underlying value is represented by
	reflect.Value. There are two functions reflect.TypeOf() and reflect.ValueOf() which return the reflect.Type and
	reflect.Value respectively. These two types are the base to create our query generator.
	*/
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
	/*
	2- reflect.Kind:
	There is one more important type in the re	reflection package called Kind. The types Kind and Type in the reflection
	package might seem similar but they have a difference which will be clear from the program below.
	*/
	o = order{
		ordId:      456,
		customerId: 56,
	}
	createQueryKind(o)
	/*
	3- NumField() and Field() methods:
	The NumField() method returns the number of fields in a struct and the Field(i int) method returns the reflect.Value
	of the ith field. You should first check the interface kind is struct because NumField and Field only works on structures.
	*/
	o = order{
		ordId:      456,
		customerId: 56,
	}
	createQueryNum(o)
	/*
	4- Int() and String() methods:
	The methods Int and String help extract the reflect.Value as an int64 and string respectively.
	 */
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)

	fmt.Printf("\nBeginning of complete program with reflection...\n")
	f := order{
		ordId:      456,
		customerId: 56,
	}
	createQueryComplete(f)
	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQueryComplete(e)
	i = 90
	createQueryComplete(i)

	fmt.Printf("\nBeginning of conclusion...\n")
	/*
	Should reflection be used?
	Rob Pike says "Clear is better than clever. Reflection is never clear."
	Reflection is a very powerful and advanced concept in Go and it should be used with care. It is very difficult to
	write clear and maintainable code using reflection. It should be avoided wherever possible and should be used only
	when absolutely necessary.
	*/
}