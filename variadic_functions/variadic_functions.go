package variadic_functions

import "fmt"

func findWithVariadicArguments(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func findWithSliceArguments(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func change(s ...string) {
	s[0] = "Go"
}

func changeAndAdd(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func RunVariadicFunctions() {
	fmt.Printf("\nBeginning of description of variadic functions...\n")
	/*
	Functions in general accept only a fixed number of arguments. A variadic function is a function that accepts a variable
	number of arguments. If the last parameter of a function definition is prefixed by ellipsis ..., then the function
	can accept any number of arguments for that parameter.
	Only the last parameter of a function can be variadic.
		func hello(a int, b ...int) {
		}
	In the above function, the parameter b is variadic since it's prefixed by ellipsis and it can accept any number of
	arguments. This function can be called by using the syntax:
		hello(1, 2) //passing one argument "2" to b
		hello(5, 6, 7, 8, 9) //passing arguments "6, 7, 8 and 9" to b
	It is also possible to pass zero arguments to a variadic function.
		hello(1)
	In the above code, we call hello with zero arguments for b. This is perfectly fine.
	If you make first parameter of the hello function variadic, compiler will complain like that: syntax error: cannot
	use ... with non-final parameter b
	*/

	fmt.Printf("\nBeginning of examples...\n")
	findWithVariadicArguments(89, 89, 90, 95)
	findWithVariadicArguments(45, 56, 67, 45, 90, 109)
	findWithVariadicArguments(78, 38, 56, 98)
	findWithVariadicArguments(87)
	/*
	The way variadic functions work is by converting the variable number of arguments to a slice of the type of the
	variadic parameter. For instance, 78 of the program above, the variable number of arguments to the find function
	are 89, 90, 95. The find function expects a variadic int argument. Hence these three arguments will be converted by
	the compiler to a slice of type int []int{89, 90, 95} and then it will be passed to the find function.
	 */

	fmt.Printf("\nBeginning of slice arguments vs variadic arguments...\n")
	/*
	We should definitely have a question lingering in your mind now. In the above section, we learned that the variadic
	arguments to a function are in fact converted a slice. Then why do we even need variadic functions when we can
	achieve the same functionality using slices? I have rewritten the program above using slices below.
	 */
	findWithSliceArguments(89, []int{89, 90, 95})
	findWithSliceArguments(45, []int{56, 67, 45, 90, 109})
	findWithSliceArguments(78, []int{38, 56, 98})
	findWithSliceArguments(87, []int{})
	/*
	The following of the advantages of using variadic arguments instead of slices:
		- There is no need to create a slice during each function call. If you look at the program above, we have created
		new slices during each function call.
		- In line no.98 of the program above, we are creating an empty slice just to satisfy the signature of the find
		function. This is totally not needed in the case of variadic functions. This line can just be find(87) when
		variadic function is used.
		- I personally feel that the program with variadic functions is more readable than the once with slices :)
	*/

	fmt.Printf("\nBeginning of append is a variadic function...\n")
	/*
	Have you ever wondered how the append function in the standard library used to append values to a slice accepts any
	number of arguments. It's because it's a variadic function.
		func append(slice []Type, elems ...Type) []Type
	The above is the definition of append function. In this definition elems is a variadic parameter. Hence append can
	accept a variable number of arguments.
	*/

	fmt.Printf("\nBeginning of passing a slice to a variadic function...\n")
	/*
	Below commented lines will not work, because these variadic arguments will be converted to a slice of type int since
	findWithVariadicArguments expects variadic int arguments. In this case, nums is already a []int slice and the compiler tries to create a
	new []int i.e the compiler tries to do.
	*/
	// nums := []int{89, 90, 95}
	// findWithVariadicArguments(89, nums)
	/*
	So is there a way to pass a slice to a variadic function? The answer is yes.
	There is a syntactic sugar which can be used to pass a slice to a variadic function. You have to suffix the slice
	with ellipsis ... If that is done, the slice is directly passed to the function without a new slice being created.
	In the above program if you replace findWithVariadicArguments(89, nums) in line no. 23 with findWithVariadicArguments(89, nums...),
	the program will compiled.
	*/
	nums := []int{89, 90, 95}
	findWithVariadicArguments(89, nums...)
	/*
	On above lines of codes, the slice is directly passed to the function without a new slice being created.
	 */

	fmt.Printf("\nBeginning of gotcha...\n")
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
	/*
	What happens on the above lines of codes?
	We are using the syntactic sugar ... and passing the slice as a variadic argument to the change function.
	As we have already discussed, if ... is used, the welcome slice itself will be passed as an argument without a
	new slice being created. Hence welcome will be passed to the change function as argument.
	 */
	welcome2 := []string{"hello", "world"}
	changeAndAdd(welcome2...)
	fmt.Println(welcome2)
}