package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

/*
The function above generates a random number and writes it to the channel data and then calls Done on the waitgroup to
notify that it is done with its task.
 */
func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

/*
The consume function creates a file named concurrent. It then reads the random numbers from the data channel and writes
to the file. Once it has read and written all the random numbers, it writes true to the done channel to notify that
it's done with its task.
 */
func consume(data chan int, done chan bool) {
	f, err := os.Create("exercises/files/concurrent.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func writeFiles() {
	fmt.Printf("\nBeginning of writing string to a file...\n")
	/*
	One of the most common file writing operation is writing string to a file. This is quite simple to do. It consists
	of the following steps.
	1- Create the file
	2- Write the string to the file
	 */
	f, err := os.Create("exercises/files/test4.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		l, err := f.WriteString("Hello World")
		if err != nil {
			fmt.Println(err)
			f.Close()
		} else {
			fmt.Println(l, "bytes written successfully")
			err = f.Close()
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	fmt.Printf("\nBeginning of writing bytes to a file...\n")
	/*
	Writing bytes to a file is quite similar to writing string to a file. We will use the Write method to write bytes
	to a file.
	 */
	f, err = os.Create("exercises/files/test3.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
		n2, err := f.Write(d2)
		if err != nil {
			fmt.Println(err)
			f.Close()
		} else {
			fmt.Println(n2, "bytes written successfully")
			err = f.Close()
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	fmt.Printf("\nBeginning of writing strings to a file line by line...\n")
	f, err = os.Create("exercises/files/lines.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
	} else {
		d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}
		for _, v := range d {
			// The Fprintln function takes a io.writer as parameter and appends a new line, just what we wanted.
			fmt.Fprintln(f, v)
			if err != nil {
				fmt.Println(err)
			}
		}
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("file written successfully")
		}
	}

	fmt.Printf("\nBeginning of appending to a file...\n")
	/*
	The file has to be opened in append and write only mode. These flags are passed parameters are passed to the Open
	function. After the file is opened in append mode, we add the new line to the file.
	*/
	f, err = os.OpenFile("exercises/files/lines.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		newLine := "File handling is easy"
		_, err = fmt.Fprintln(f, newLine)
		if err != nil {
			fmt.Println(err)
			f.Close()
		} else {
			err = f.Close()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("file appended successfully")
			}
		}
	}

	fmt.Printf("\nBeginning of writing to a file concurrently...\n")
	/*
	When multiple goroutines write to a file concurrently, we will end up with a race condition. Hence concurrent
	writes to a file should be co-ordinated using a channel.
	We will write a program that creates 100 goroutines. Each of this goroutine will generate a random number
	concurrently, thus generating hundred random numbers in total. These random numbers will be written to a file. We
	will solve this problem by using the following approach.
	1- Create a channel which will be used to read and write the generated random numbers.
	2- Create 100 producer goroutines. Each goroutine will generate a random number and will also write the random
	number to a channel
	3- Create a consumer goroutine which will read from the channel and write the generated random number to the
	file. Thus we have only one goroutine writing to a file concurrently thereby avoiding race condition :)
	4- Close the file once done.
	*/
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	/*
	The goroutine call in line no. 169 calls wait() on the waitgroup to wait for all 100 goroutines to finish creating
	random numbers. After that it closes the channel.
	*/
	go func() {
		wg.Wait()
		close(data)
	}()
	/*
	Once the channel is closed and the consume goroutine has finished writing all generated random numbers to the file,
	it writes true to the done channel in line no. 37 and the main goroutine is unblocked and prints File written
	successfully.
	 */
	d := <- done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}