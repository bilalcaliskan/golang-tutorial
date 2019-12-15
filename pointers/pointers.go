package pointers

import "fmt"

func ChangeWithAddress(val *int) {
	*val = 55
}

func ChangeWithValue(val int) {
	val = 50
}

func Hello() *int {
	i := 5
	return &i
}

func Run()  {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}