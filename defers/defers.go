package defers

import (
	"fmt"
	"sync"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func printA(a int) {
	fmt.Println()
	fmt.Println("value of a in deferred function", a)
}

func RunDefers() {
	fmt.Printf("\nBeginning of introduction...\n")
	/*
	Defer statement is used to execute a function call just before the surrounding function where the defer statement
	is present returns.
	In the below example, finished() function will be called just before the largest function returns. defer in the first
	line of largest function states it.
	*/
	nums := []int{68, 2394, 3, 1234, 4949}
	largest(nums)

	fmt.Printf("\nBeginning of deferred methods...\n")
	/*
	Defer is not restricted only to functions. It is perfectly legal to defer a method call too.
	*/
	p := person{
		firstName: "John",
		lastName: "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")

	fmt.Printf("\nBeginning of arguments evaluation...\n")
	/*
	The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual
	function call is done.
	*/
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)
	/*
	Output of above program says that although the value of a changes to 10 after the defer statement is executed, the
	actual deferred function call printA(a) still prints 5.
	*/

	fmt.Printf("\nBeginning of stack of defers...\n")
	/*
	When a function has multiple defer calls, they are pushed on to a stack and executed in Last In First Out (LIFO)
	order.
	*/
	name := "Naveen"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
	fmt.Println()
	/*
	In the program above, the for range loop in line no. 11, iterates the string and calls defer fmt.Printf("%c", v).
	These deferred calls will be added to a stack.
	The stack is a last in first out datastructure. The defer call that is pushed to the stack last will be pulled out
	and executed first.
	*/

	fmt.Printf("\nBeginning of practical use of defer...\n")
	/*
	Defer is used in places where a function call should be executed irrespective of the code flow.
	Deferred function will be called just before the surrounding function returns.
	*/
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	r4 := rect{10, 100}
	rects := []rect{r1, r2, r3, r4}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
	/*
	In the program above, we have removed the 3 wg.Done() calls in the area method of rect and replaced it with a single
	defer wg.Done() call. This makes the code more simple and understandable.
	There is one more advantage of using defer in the above program. Let's say we add another return path to the area
	method using a new if condition. If the call to wg.Done() was not deferred, we have to be careful and ensure that
	we call wg.Done() in this new return path. But since the call to wg.Done() is defered, we need not worry about adding
	new return paths to this method.
	 */
}