package error_handling

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
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

func RunErrorHandling() {
	fmt.Printf("\nBeginning of deferred functions...\n")
	fmt.Println("The finished() function will be called just before the largest function returns.")
	nums := []int{68, 2394, 3, 1234, 4949}
	largest(nums)

	fmt.Printf("\nBeginning of deferred methods...\n")
	p := person{"John", "Smith"}
	defer p.fullName()
	fmt.Printf("Welcome ")

	fmt.Printf("\nBeginning of arguments evaluation...\n")
	fmt.Println("The arguments of a deferred function are evaluated when the defer statement is executed and " +
		"not when the actual function call is done.")
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

	fmt.Printf("\nBeginning of stack of defers...\n")
	fmt.Println("When a function has multiple defer calls, they are pushed on to a stack and executed in Last " +
		"In First Out (LIFO) order.")
	name := "Naveen"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}

	fmt.Printf("\nBeginning of practical use of defer...\n")
	fmt.Println("Defer is used in places where a function call should be executed irrespective of the code flow.")
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

	fmt.Printf("\nBeginning of errors...\n")
	fmt.Println("If a function or method returns an error, then by convention it has to be the last value " +
		"returned from the function.")
	fmt.Println("The idiomatic way of handling error in Go is to compare the returned error to nil. A nil value " +
		"indicates that no error has occurred and a non nil value indicates the presence of an error.")
	file, err := os.Open("/test.txt")
	if err != nil {
		// "Error() string" method of error interface is just a string, so we can print it
		fmt.Println(err)
	} else {
		fmt.Println(file.Name(), "opened successfully")
	}

	fmt.Printf("\nBeginning of extracting more information from errors method 1...\n")
	fmt.Println("Asserting the underlying struct type and getting more information from the struct fields.")
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		//return
	} else {
		fmt.Println(f.Name(), "opened successfully")
	}

	fmt.Printf("\nBeginning of extracting more information from errors method 2...\n")
	addr, err := net.LookupHost("golangbot.com")
	if err, ok := err.(*net.DNSError); ok {
		fmt.Println(ok)
		if err.Timeout() {
			fmt.Println(err.Name)
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println(err.Name)
			fmt.Println("temporary error")
		} else {
			fmt.Println(err.Name)
			fmt.Println("generic error: ", err)
		}
		//return
	} else if err == nil {
		fmt.Println(addr)
	}

	fmt.Printf("\nBeginning of extracting more information from errors method 3...\n")
	files, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		//return
	} else if err == nil {
		fmt.Println("matched files", files)
	}

	fmt.Printf("\nBeginning of ignoring errors...\n")
	files2, _ := filepath.Glob("[")
	fmt.Println("matched files", files2)
}