package variadicfunctions

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

func variadicArguments() {
	findWithVariadicArguments(89, 89, 90, 95)
	findWithVariadicArguments(45, 56, 67, 45, 90, 109)
	findWithVariadicArguments(78, 38, 56, 98)
	findWithVariadicArguments(87)
}

func sliceArguments() {
	/**
	Same functionality with the above variadicArguments function can be done with the below methods. But there is a little
	bit disadvantage of it.
	The following of the advantages of using variadic arguments instead of slices.
		- There is no need to create a slice during each function call.
		- We are creating an empty slice just to satisfy the signature of the find function. This is totally not needed
		in the case of variadic functions.
	*/
	findWithSliceArguments(89, []int{89, 90, 95})
	findWithSliceArguments(45, []int{56, 67, 45, 90, 109})
	findWithSliceArguments(78, []int{38, 56, 98})
	findWithSliceArguments(87, []int{})
}

func Run() {
	/**
	Functions in general accept only a fixed number of arguments. A variadic function is a function that accepts a variable
	number of arguments. If the last parameter of a function definition is prefixed by ellipsis ..., then the function
	can accept any number of arguments for that parameter.
	Only the last parameter of a function can be variadic
	The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic
	parameter.
	Remark that append is a variadic function
	func append(slice []Type, elems ...Type) []Type
	*/
	variadicArguments()
	sliceArguments()
}