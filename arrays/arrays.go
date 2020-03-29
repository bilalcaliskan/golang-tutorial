package arrays

import "fmt"

func changeLocal(numbers [5]int) {
	numbers[0] = 55
	fmt.Println("inside function", numbers)
}

func printArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func RunArrays()  {
	fmt.Printf("\nBeginning of declaring arrays...\n")
	/*
	An array is a collection of elements that belong to the same type. For example the collection of integers 5, 8, 9,
	79, 76 form an array. Mixing values of different types, for example an array that contains both strings and integers
	is not allowed in Go.
	An array belongs to type [n]T. n denotes the number of elements in an array and T represents the type of each element.
	The number of elements n is also a part of the type
	 */
	/*
	var a [3]int declares a integer array of length 3. All elements in an array are automatically assigned the zero value
	of the array type. In this case a is an integer array and hence all elements of a are assigned to 0, the zero value
	of int. Running the above program will output [0 0 0].
	 */
	var a [3] int // int array with length 3
	fmt.Println(a)
	/*
	The index of an array starts from 0 and ends at length - 1. Lets assign some values to the above array.
	 */
	var b [3]int // int array with length 3
	b[0] = 12
	b[1] = 14
	b[2] = 16
	fmt.Println(b)
	/*
	Lets create the same array using the short hand declaration.
	 */
	c := [3]int{12, 78, 50}
	fmt.Println(c)
	/*
	It is not necessary that all elements in an array have to be assigned a value during short hand declaration.
	 */
	d := [3]int{12}
	fmt.Println(d)
	/*
	In the above program in line no. 8 a := [3]int{12} declares an array of length 3 but is provided with only one
	value 12. The remaining 2 elements are assigned 0 automatically. The program will output [12 0 0]
	 */
	/*
	You can even ignore the length of the array in the declaration and replace it with ... and let the compiler find the
	length for you. This is done in the following program.
	 */
	e := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(e)
	/* !!!
	The size of the array is a part of the type. Hence [5]int and [25]int are distinct types. Because of this, arrays
	cannot be resized. Don't worry about this restriction since slices exist to overcome this.
	 */
	f := [3]int{5, 78, 8}
	var g [5]int
	// g = f // it is not possible since [3]int and [5]int are distinct types
	fmt.Println(f)
	fmt.Println(g)

	fmt.Printf("\nBeginning of arrays as value types...\n")
	/*
	Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable, a
	copy of the original array is assigned to the new variable. If changes are made to the new variable, it will not be
	reflected in the original array.
	 */
	h := [...]string{"USA", "China", "India", "Germany", "France"}
	j := h // a copy of a is assigned to b
	j[0] = "Singapore"
	fmt.Println("h is ", h)
	fmt.Println("j is ", j)
	/*
	Similarly when arrays are passed to functions as parameters, they are passed by value and the original array in
	unchanged.
	 */
	numbers := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function", numbers)
	changeLocal(numbers) // num is passed by value
	fmt.Println("after passing to function", numbers)


	fmt.Printf("\nBeginning of iterating arrays using range...\n")
	/*
	The for loop can be used to iterate over elements of an array.
	*/
	floatArray := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of floatArray is", len(floatArray))
	for i := 0; i < len(floatArray); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of floatArray array is %.2f\n", i, floatArray[i])
	}
	fmt.Println()
	/*
	Go provides a better and concise way to iterate over an array by using the range form of the for loop. range returns
	both the index and the value at that index.
	 */
	floatNumbers := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range floatNumbers { //range returns both the index and value
		fmt.Printf("%d the element of floatNumbers array is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of floatNumbers array",sum)
	/*
	In case you want only the value and want to ignore the index, you can do this by replacing the index with
	the _ blank identifier.
		for _, v := range a { //ignores index
		}
	*/

	fmt.Printf("\nBeginning of multidimensional arrays...\n")
	/*
	The arrays we created so far are all single dimension. It is possible to create multidimensional arrays.
	 */
	multiDimensionalArray1 := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printArray(multiDimensionalArray1)
	var multiDimensionalArray2 [3][2]string
	multiDimensionalArray2[0][0] = "apple"
	multiDimensionalArray2[0][1] = "samsung"
	multiDimensionalArray2[1][0] = "microsoft"
	multiDimensionalArray2[1][1] = "google"
	multiDimensionalArray2[2][0] = "AT&T"
	multiDimensionalArray2[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printArray(multiDimensionalArray2)
}