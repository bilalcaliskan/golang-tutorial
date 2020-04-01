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
	// go b() // commented out to prevent runtime panic
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
	fmt.Println(n[3]) // this line will cause to "runtime error: index out of range [3] with length 3"
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
	The panic will not be recovered. This is because the recovery function is present in a different goroutine(main
	goroutine in that case) and the panic is happening in different goroutine(b goroutine in that case). Hence recovery
	is not possible. If you change line 25 to "b()", panic will be recovered cause b function will run on main goroutine
	also.
	*/

	fmt.Printf("\nBeginning of runtime panics...\n")
	/*
	Panics can also be caused by runtime errors such as array out of bounds access. This is equivalent to a call of
	the built-in function panic with an argument defined by interface type runtime.Error.
	If you comment out line 35, runtime panic will be occured.
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
	/*
	From the output you can understand that first the panic is recovered and Recovered runtime error: index out of range
	is printed. Following that the stack trace is printed. Then normally returned from main is printed after the panic
	has recovered.
	 */
}