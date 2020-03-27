package functions

import "fmt"

/*
If consecutive parameters are of the same type, we can avoid writing the type each time and it is enough to be
written once at the end.ie price int, no int can be written as price, no int. The above function can hence be
rewritten as,
 */
func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

/*
If a function returns multiple return values then they must be specified between ( and )
 */
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

/*
It is possible to return named values from a function. If a return value is named, it can be considered as being declared
as a variable in the first line of the function. But with zero value. You can then assign another value.
*/
func rectPropsNamedReturn(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // no explicit return value because the function knows what to return, but you can return with return values if you like
}

func RunFunctions() {
	fmt.Printf("\nBeginning of introduction to functions...\n")
	/*
	A function is a block of code that performs a specific task. A function takes a input, performs some calculations on
	the input and generates a output.
	The parameters and return type are optional in a function. The general syntax for declaring a function in go is:
	func functionname(parametername type) returntype {
		//function body
	}
	*/
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	fmt.Printf("\nBeginning of multiple return values...\n")
	/*
	It is possible to return multiple values from a function.
	 */
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %.2f Perimeter %.2f\n", area, perimeter)

	fmt.Printf("\nBeginning of named return values...\n")
	/*
	It is possible to return named values from a function. If a return value is named, it can be considered as being
	declared as a variable in the first line of the function.
	 */
	area2, perimeter2 := rectPropsNamedReturn(10.8, 5.6)
	fmt.Printf("Area %.2f Perimeter %.2f\n", area2, perimeter2)

	fmt.Printf("\nBeginning of blank identifier...\n")
	/*
	_ is know as the blank identifier in Go. It can be used in place of any value of any type.
	The rectProps function returns the area and perimeter of the rectangle. What if we only need the area and want to
	discard the perimeter. This is where _ is of use.
	*/
	area3, _ := rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f ", area3)
}