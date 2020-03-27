package types

import (
	"fmt"
	"unsafe"
)

func RunTypes() {
	fmt.Printf("\nBeginning of data types in Go...\n")
	/*
	The following are the basic types available in go:
	- bool
	- Numeric Types
		int8, int16, int32, int64, int
		uint8, uint16, uint32, uint64, uint
		float32, float64
		complex64, complex128
		byte
		rune
	- string
	 */

	fmt.Printf("\nBeginning of bool...\n")
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	fmt.Printf("\nBeginning of signed integers...\n")
	/*
	int8: represents 8 bit signed integers
	size: 8 bits
	range: -128 to 127

	int16: represents 16 bit signed integers
	size: 16 bits
	range: -32768 to 32767

	int32: represents 32 bit signed integers
	size: 32 bits
	range: -2147483648 to 2147483647

	int64: represents 64 bit signed integers
	size: 64 bits
	range: -9223372036854775808 to 9223372036854775807

	int: represents 32 or 64 bit integers depending on the underlying platform. You should generally be using int to
	represent integers unless there is a need to use a specific sized integer.
	size: 32 bits in 32 bit systems and 64 bit in 64 bit systems.
	range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems
	 */
	var e int = 89
	f := 95
	fmt.Println("value of a is", e, "and b is", f)
	/*
	The type of a variable can be printed using %T format specifier in Printf function. Go has a package unsafe which
	has a Sizeof function which returns in bytes the size of the variable passed to it. unsafe package should be used
	with care as the code using it might have portability issues, but for the purposes of this tutorial we can use it.
	 */
	var ab int = 89
	ba := 95
	fmt.Println("value of a is", ab, "and b is", ba)
	fmt.Printf("type of a is %T, size of a is %d", ab, unsafe.Sizeof(ab)) //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", ba, unsafe.Sizeof(ba)) //type and size of b

	fmt.Printf("\nBeginning of unsigned integers...\n")
	/*
	uint8: represents 8 bit unsigned integers
	size: 8 bits
	range: 0 to 255

	uint16: represents 16 bit unsigned integers
	size: 16 bits
	range: 0 to 65535

	uint32: represents 32 bit unsigned integers
	size: 32 bits
	range: 0 to 4294967295

	uint64: represents 64 bit unsigned integers
	size: 64 bits
	range: 0 to 18446744073709551615

	uint : represents 32 or 64 bit unsigned integers depending on the underlying platform.
	size : 32 bits in 32 bit systems and 64 bits in 64 bit systems.
	range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems
	 */

	fmt.Printf("\nBeginning of floating point types...\n")
	/*
	float32: 32 bit floating point numbers
	float64: 64 bit floating point numbers(if you define a float variable without specifying a type, it will be float64)
	 */
	// The type of a and b is inferred from the value assigned to them. In this case a and b are of type float64.(float64
	//is the default type for floating point values)
	ca, da := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", ca, da)
	sum := ca + da
	diff := ca - da
	fmt.Println("sum", sum, "diff", diff)
	no1, no2 := 56, 89
	fmt.Println("sum", no1 + no2, "diff", no1 - no2)

	fmt.Printf("\nBeginning of complex types...\n")
	/*
	complex64: complex numbers which have float32 real and imaginary parts
	complex128: complex numbers with float64 real and imaginary parts

	The builtin function complex is used to construct a complex number with real and imaginary parts. The complex
	function has the following definition:
		func complex(r, i FloatType) ComplexType
		It takes a real and imaginary part as parameter and returns a complex type. Both the real and imaginary parts
		must be of the same type. ie either float32 or float64. If both the real and imaginary parts are float32, this
		function returns a complex value of type complex64. If both the real and imaginary parts are of type float64, this
		function returns a complex value of type complex128
	*/
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	fmt.Printf("\nBeginning of other numeric types...\n")
	/*
	byte is an alias of uint8
	rune is an alias of int32
	 */

	fmt.Printf("\nBeginning of string type...\n")
	/*
	Strings are a collection of bytes in golang. It's alright if this definition doesn't make any sense. For now we can
	assume a string to be a collection of characters.
	*/
	first := "Naveen"
	last := "Ramanathan"
	name := first +" "+ last
	fmt.Println("My name is",name)

	fmt.Printf("\nBeginning of type conversion...\n")
	/*
	Go is very strict about explicit typing. There is no automatic type promotion or conversion.
	 */
	i := 55      //int
	j := 67.8    //float64
	sum2 := i + int(j) //int + float64 not allowed, so we must convert one of them to other
	fmt.Println(sum2)
	/*
	The same is the case with assignment. Explicit type conversion is required to assign a variable of one type to
	another.
	 */
	ij := 10
	var ji float64 = float64(ij) //this statement will not work without explicit conversion
	fmt.Println("j", ji)
}