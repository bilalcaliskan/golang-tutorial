package maps

import "fmt"

func RunMaps() {
	fmt.Printf("\nBeginning of introduction to maps...\n")
	/*
	A map is a builtin type in Go which associates a value to a key. The value can be retrieved using the corresponding
	key.
	A map can be created by passing the type of key and value to the make function. make(map[type of key]type of value)
	is the syntax to create a map.
	 */
	personSalaries := make(map[string]int)
	fmt.Println(personSalaries)
	/*
	The above line of code creates a map named personSalary which has string keys and int values.
	The zero value of a map is nil. If you try to add items to nil map, a run time panic will occur. Hence the map has
	to be initialized using make function.
	 */
	var personSalaries2 map[string]int
	if personSalaries2 == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalaries2 = make(map[string]int)
	}
	/*
	In the above program, personSalary is nil and hence it will be initialised using the make function.
	 */

	fmt.Printf("\nBeginning of adding items to a map...\n")
	/*
	The syntax for adding new items to a map is the same as that of arrays.
	 */
	personSalaries3 := make(map[string]int)
	personSalaries3["steve"] = 12000
	personSalaries3["jamie"] = 15000
	personSalaries3["mike"] = 9000
	fmt.Println("personSalaries3 map contents:", personSalaries3)
	/*
	It is also possible to initialize a map during declaration itself
	 */
	personSalaries4 := map[string]int {
		"steve" : 12000,
		"jamie" : 15000,
	}
	personSalaries4["mike"] = 9000
	fmt.Println("personSalaries4 map contents:", personSalaries4)
	/*
	It's not necessary that only string types should be keys. All comparable types such as boolean, integer, float,
	complex, string, ... can also be keys.
	http://golang.org/ref/spec#Comparison_operators
	 */

	fmt.Printf("\nBeginning of accessing items of a map...\n")
	personSalaries5 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalaries5["mike"] = 9000
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalaries5[employee])
	/*
	What will happen if a element is not present? The map will return the zero value of the type of that element. In the
	case of personSalaries5 map, if we try to access an element which is not present then, the zero value of int which is 0
	will be returned.
	 */
	fmt.Println("Salary of joe is", personSalaries5["joe"])
	/*
	The above program returns the salary of joe as 0. We did not get any runtime error stating that the key joe is not
	present in the personSalary map.
	What if we want to know whether a key is present in a map or not?
		value, ok := map[key]
	The above is the syntax to find out whether a particular key is present in a map. If ok is true, then the key is
	present and its value is present in the variable value, else the key is absent.
	*/
	personSalaries6 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalaries6["mike"] = 9000
	newEmp := "joe"
	value, ok := personSalaries6[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp,"not found")
	}
	/*
	The range form of the for loop is used to iterate over all elements of a map.
	 */
	personSalaries7 := map[string]int {
		"steve": 12000,
		"jamie": 15000,
	}
	personSalaries7["mike"] = 9000
	fmt.Println("All items of a map")
	for key, value := range personSalaries7 {
		fmt.Printf("personSalaries7[%s] = %d\n", key, value)
	}
	/*
	One important fact is that the order of the retrieval of values from a map when using for range is not guaranteed
	to be the same for each execution of the program.
	 */

	fmt.Printf("\nBeginning of deleting items...\n")
	/*
	delete(map, key) is the syntax to delete key from a map. The delete function does no return any value.
	 */
	personSalaries8 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalaries8["mike"] = 9000
	fmt.Println("map before deletion", personSalaries8)
	delete(personSalaries8, "steve")
	fmt.Println("map after deletion", personSalaries8)

	fmt.Printf("\nBeginning of length of the map...\n")
	/*
	Length of the map can be determined using the len function.
	 */
	fmt.Println("length of personSalaries8 map is", len(personSalaries8))

	fmt.Printf("\nBeginning of maps are reference types...\n")
	/*
	Similar to slices, maps are reference types. When a map is assigned to a new variable, they both point to the same
	internal data structure. Hence changes made in one will reflect in the other.
	Similar is the case when maps are passed as parameters to functions. When any change is made to the map inside the
	function, it will be visible to the caller.
	 */
	personSalaries9 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalaries9["mike"] = 9000
	fmt.Println("Original person salary", personSalaries9)
	newPersonSalaries := personSalaries9
	newPersonSalaries["mike"] = 18000
	fmt.Println("Person salary changed", personSalaries9)

	fmt.Printf("\nBeginning of maps equality...\n")
	/*
	Maps can't be compared using the == operator. The == can be only used to check if a map is nil.
	One way to check whether two maps are equal is to compare each one's individual elements one by one.
	 */
}