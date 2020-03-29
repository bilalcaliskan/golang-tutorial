package slices

import "fmt"

func substractTwo(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func RunSlices() {
	fmt.Printf("\nBeginning of description of slices...\n")
	/*
	Although arrays seem to be flexible enough, they come with the restriction that they are of fixed length. It is not
	possible to increase the length of an array. This is were slices come into picture. In fact in Go, slices are more
	common than conventional arrays.
	A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own.
	They are the just references to existing arrays.
	 */
	fmt.Printf("\nBeginning of creating a slice...\n")
	/*
	A slice with elements of type T is represented by []T
	The syntax a[start:end] creates a slice from array a starting from index start to index end - 1.
	 */
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)
	/*
	Another example. c := []int{6, 7, 8} creates an array with 3 integers and returns a slice reference which is stored
	in c.
	 */
	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(c)

	fmt.Printf("\nBeginning of modifying a slice...\n")
	/*
	A slice does not own any data of its own. It is just a representation of the underlying array. Any modifications
	done to the slice will be reflected in the underlying array. It is just a reference type to arrays.
	 */
	darr := []int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
	/*
	When a number of slices share the same underlying array, the changes that each one makes will be reflected in
	the array.
	 */
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
	fmt.Println("slices after all modifications", nums1, nums2)

	fmt.Printf("\nBeginning of length and capacity of a slice...\n")
	/*
	The length of the slice is the number of elements in the slice. The capacity of the slice is the number of elements
	in the underlying array starting from the index from which the slice is created.
	 */
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Println(fruitslice)
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	/*
	A slice can be re-sliced upto its capacity. Anything beyond that will cause the program to throw a run time error.
	 */
	fruitslice = fruitslice[0:cap(fruitslice)] //re-slicing furitslice till its capacity
	fmt.Println(fruitslice)
	fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))

	fmt.Printf("\nBeginning of creating a slice using make...\n")
	/*
	func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity. The capacity
	parameter is optional and defaults to the length. The make function creates an array and returns a slice reference
	to it.
	 */
	makeSlice := make([]int, 5, 5)
	fmt.Println(makeSlice)
	/*
	The values are zeroed by default when a slice is created using make. The above program will output [0 0 0 0 0].
	 */

	fmt.Printf("\nBeginning of appending to a slice...\n")
	/*
	As we already know arrays are restricted to fixed length and their length cannot be increased. Slices are dynamic
	and new elements can be appended to the slice using append function. The definition of append function is func
	append(s []T, x ...T) []T.
	x ...T in the function definition means that the function accepts variable number of arguments for the parameter x.
	These type of functions are called variadic functions.
	One question might be bothering you though. If slices are backed by arrays and arrays themselves are of fixed length
	then how come a slice is of dynamic length. Well what happens under the hood is, when new elements are appended to
	the slice, a new array is created. The elements of the existing array are copied to this new array and a new slice
	reference for this new array is returned. The capacity of the new slice is now twice that of the old slice. Pretty
	cool right :). The following program will make things clear.
	*/
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota") // you can append multiple elements like Toyota
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	/*
	The zero value of a slice type is nil. A nil slice has length and capacity 0. It is possible to append values to a
	nil slice using the append function.
	 */
	var names []string // zero value of a slice is nil
	if names == nil {
		fmt.Println("slice has old length", len(names), "and capacity", cap(names)) //capacity of cars is 3
		fmt.Println("slice is nil going to append")
		names = append(names, "Josh", "Sebastian", "Vinay", "asddsaf")
		fmt.Println("slice has new length", len(names), "and capacity", cap(names))
	}
	/*
	It is also possible to append one slice to another using the ... operator. You can learn more about this operator in
	the variadic functions tutorial.
	 */
	veggies := []string{"potatoes","tomatoes","brinjal"}
	fruits := []string{"oranges","apples"}
	food := append(veggies, fruits...) // we have appended slice fruits into slice veggies
	fmt.Println("food:",food)

	fmt.Printf("\nBeginning of passing a slice to a function...\n")
	/*
	Slices can be thought of as being represented internally by a structure type. This is how it looks:
		type slice struct {
			Length        int
			Capacity      int
			ZerothElement *byte
		}
	A slice contains the length, capacity and a pointer to the zeroth element of the array. When a slice is passed to a
	function, even though it's passed by value, the pointer variable will refer to the same underlying array. Hence when
	a slice is passed to a function as parameter, changes made inside the function are visible outside the function too.
	*/
	nos := []int{8, 7, 6} // creates an array and returns the slice reference
	fmt.Println("slice before function call", nos)
	substractTwo(nos) //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside
	/*
	The function call in line number 17 of the above program decrements each element of the slice by 2. When the slice
	is printed after the function call, these changes are visible. If you can recall, this is different from an array
	where the changes made to an array inside a function are not visible outside the function.
	 */

	fmt.Printf("\nBeginning of multidimensional slices...\n")
	/*
	Similar to arrays, slices can have multiple dimensions.
	*/
	pls := [][]string {
		{"C", "C++"},
		{"Javascript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\nBeginning of multidimensional slices...\n")
	/**
	Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage
	collected. This might be of concern when it comes to memory management.
	Lets assume that we have a very large array and we are interested in processing only a small part of it. Henceforth
	we create a slice from that array and start processing the slice. The important thing to be noted here is that the
	array will still be in memory since the slice references it.

	One way to solve this problem is to use the copy function func copy(dst, src []T) int to make a copy of that slice.
	This way we can use the new slice and the original array can be garbage collected.
	*/
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	fmt.Println(countriesCpy)
}
