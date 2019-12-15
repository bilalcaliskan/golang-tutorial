package arrays

import "fmt"

func changeLocal(numbers [5]int) {
	numbers[0] = 55
	fmt.Println("inside function", numbers)
}

func Run()  {
	numbers := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function", numbers)
	changeLocal(numbers)
	fmt.Println("after passing to function", numbers)

	floatArray := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of floatArray is", len(floatArray))
	for i := 0; i < len(floatArray); i++ {
		fmt.Printf("%d th element of floatArray is %.2f\n", i, floatArray[i])
	}

	intArray := [...]int{12, 24, 36, 48, 60}
	sum := 0
	fmt.Printf("length of intArray is %d and sum of intArray is %d", len(intArray), sum)
	for i, v := range intArray {
		fmt.Printf("%d the element of intArray is %d\n", i, v)
		sum += v
	}
	fmt.Printf("Sum of all elements of intArray is %d", sum)

	for _, v := range intArray {
		fmt.Printf("element of intArray is %d\n", v)
	}

}