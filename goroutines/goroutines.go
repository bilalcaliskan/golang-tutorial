package goroutines

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func RunGoroutines() {
	fmt.Printf("\nBeginning of starting Goroutines...\n")
	fmt.Println("When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the " +
		"control does not wait for the Goroutine to finish executing. The control returns immediately to the next line " +
		"of code after the Goroutine call and any return values from the Goroutine are ignored.")
	fmt.Println("The main Goroutine should be running for any other Goroutines to run. If the main Goroutine " +
		"terminates then the program will be terminated and no other Goroutine will run.")
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main goroutine")

	fmt.Printf("\nBeginning of multiple Goroutines...\n")
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main goroutine terminated")
}