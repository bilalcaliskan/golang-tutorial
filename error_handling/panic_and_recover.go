package error_handling

import "fmt"

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func deferredFullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func recoverFullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func RunPanicAndRecover() {
	fmt.Printf("\nBeginning of introduction...\n")
	/*
	The idiomatic way to handle abnormal conditions in a program in Go is using errors. Errors are sufficient for most
	of the abnormal conditions arising in the program. But there are some situations where the program cannot simply
	continue executing after an abnormal situation. In this case we use panic to terminate the program. When a function
	encounters a panic, its execution is stopped, any deferred functions are executed and then the control returns to
	its caller. This process continues until all the functions of the current goroutine have returned at which point
	the program prints the panic message, followed by the stack trace and then terminates. This concept will be more
	clear when we write a sample program.
	It is possible to regain control of a panicking program using recover. Panic and recover can be considered similar
	to try-catch-finally idiom in other languages except that it is rarely used and when used is more elegant and
	results in clean code.
	 */

	fmt.Printf("\nBeginning of when to use panic...\n")
	/*
	One important factor is that you should avoid panic and recover and use errors where ever possible. Only in cases
	where the program just cannot continue execution should a panic and recover mechanism be used.
	There are two valid use cases for panic:
		1- An unrecoverable error where the program cannot simply continue its execution; one example would be a web
		server which fails to bind to the required port. In this case it's reasonable to panic as there is nothing else
		to do if the port binding itself fails.
		2- A programmer error; Let's say we have a method which accepts a pointer as a parameter and someone calls
		this method using nil as argument. In this case we can panic as it's a programmer error to call a method with
		nil argument which was expecting a valid pointer.
	 */

	fmt.Printf("\nBeginning of panic example...\n")
	/*
	The signature of the built in panic function is provided like that; func panic(interface{})
	 */
	firstName := "Elon"
	lastName := "Mask"
	// uncomment below line to panic
	//fullName(&firstName, nil)
	fullName(&firstName, &lastName)
	fmt.Println("returned normally from main")

	fmt.Printf("\nBeginning of defer while panicking...\n")
	/*
	When a function encounters a panic, its execution is stopped, any deferred functions are executed and then the
	control returns to its caller. This process continues until all the functions of the current goroutine have
	returned at which point the program prints the panic message, followed by the stack trace and then terminates.
	 */
	defer fmt.Println("deferred call in main goroutine")
	firstName = "Hasan"
	lastName = "Huseyin"
	// uncomment below line to panic. First deferred call on deferredFullName will be run, then the deferred call on main
	// when deferredFullName function panics, any deferred function calls are first executed and then the control
	//returns to the caller whose deferred calls are executed and so on until the top level caller is reached. Caller
	// is the main goroutine in our case.
	//deferredFullName(&firstName, nil)
	fmt.Println("returned normally from main")
	/*
	When all the deferred calls are executed from down to up, control reaches the top level function and hence the
	program prints the panic message followed by the stack trace and then terminates.
	 */

	fmt.Printf("\nBeginning of recover...\n")
	/*
	Recover is a builtin function which is used to regain control of a panicking goroutine. The signature of recover
	function is like that; func recover() interface{}.
	Recover is useful only when called inside deferred functions. Executing a call to recover inside a deferred
	function stops the panicking sequence by restoring normal execution and retrieves the error value passed to the
	call of panic. If recover is called outside the deferred function, it will not stop a panicking sequence.
	 */
	defer fmt.Println("deferred call in main")
	firstName = "Elon"
	recoverFullName(&firstName, nil)
	fmt.Println("returned normally from main")
	/*
	After execution of recover(), the panicking stops and the control returns to the caller, in this case the main
	function and the program continues to execute normally from line 29 in main right after the panic.
	 */
}