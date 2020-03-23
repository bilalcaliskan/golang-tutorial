package error_handling

import (
	"fmt"
	"runtime/debug"
	"time"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func recoveryWithStackTrace() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
		debug.PrintStack()
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked!")
}

func c() {
	defer recovery()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from c")
}

func d() {
	defer recoveryWithStackTrace()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from c")
}

func RunPanicAndRecover2() {
	fmt.Printf("\nBeginning of panic, recover and goroutines...\n")
	/*
		Recover works only when it is called from the same goroutine. It's not possible to recover from a panic that has
		happened in a different goroutine.
	*/
	a()
	fmt.Println("normally returned from main goroutine")
	/*
	The panic will not be recovered. This is because the recovery function is present in a different gouroutine and the
	panic is happening in function b() in a different goroutine. Hence recovery is not possible.
	*/

	fmt.Printf("\nBeginning of runtime panics...\n")
	/*
	Panics can also be caused by runtime errors such as array out of bounds access. This is equivalent to a call of
	the built-in function panic with an argument defined by interface type runtime.Error.
	 */
	c()
	fmt.Println("normally returned from main")

	fmt.Printf("\nBeginning of getting stack trace after recover...\n")
	/*
	If we recover a panic, we loose the stack trace about the panic. Even in the program above after recovery, we lost
	the stack trace. There is a way to print the stack trace using the PrintStack function of the Debug package.
	 */
	d()
	fmt.Println("normally returned from main")
}