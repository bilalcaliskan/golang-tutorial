package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func processSingle(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func RunBufferedPools() {
	fmt.Printf("\nBeginning of description of buffered channels...\n")
	/*
	All the channels we discussed in the previous tutorial were basically unbuffered. As we discussed in the channels
	tutorial in detail, sends and receives to an unbuffered channel are blocking.
	It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer
	is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.
	Buffered channels can be created by passing an additional capacity parameter to the make function which specifies
	the size of the buffer.
		ch := make(chan type, capacity)
	Capacity in the above syntax should be greater than 0 for a channel to have a buffer. The capacity for an
	unbuffered channel is 0 by default and hence we omitted the capacity parameter while creating channels.
	 */
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	/*
	In the program above, we create a buffered channel with a capacity of 2. Since the channel has a capacity of 2, it
	is possible to write 2 strings into the channel without being blocked. We write 2 strings to the channel and the
	channel does not block.
	*/

	fmt.Printf("\nBeginning of another example on buffered channels...\n")
	/*
	Lets look at one more example of buffered channel in which the values to the channel are written in a concurrent
	Goroutine and read from the main Goroutine.
	 */
	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(2 * time.Second)
	for v := range ch2 {
		fmt.Println("read value", v,"from ch")
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("\nBeginning of deadlocks...\n")
	ch3 := make(chan string, 2)
	ch3 <- "naveen"
	ch3 <- "paul"
	//ch3 <- "steve"
	/**
	If we uncomment above line, there will be a deadlock. ch3 channel has capacity 2 and to be able to write that
	channel, somebody must read from that channel. But there is no concurrent routine reading from that channel. Some
	Goroutine must read from the channel in order for the write to proceed, but in this case there is no concurrent
	routine reading from this channel. Hence there will be a deadlock and the program will panic at run time.
	 */
	fmt.Println(<- ch3)
	fmt.Println(<- ch3)

	fmt.Printf("\nBeginning of length vs capacity...\n")
	ch4 := make(chan string, 3)
	ch4 <- "naveen"
	ch4 <- "paul"
	// below will print 3
	fmt.Println("capacity is", cap(ch4))
	// below will print 2
	fmt.Println("length is", len(ch4))
	fmt.Println("read value", <- ch4)
	// below will print 1
	fmt.Println("new length is", len(ch4))

	fmt.Printf("\nBeginning of waitgroup...\n")
	/*
	A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all
	Goroutines finish executing. Lets say we have 3 concurrently executing Goroutines spawned from the main Goroutine.
	The main Goroutines needs to wait for the 3 other Goroutines to finish before terminating. This can be accomplished
	using WaitGroup.
	 */
	/*
	WaitGroup is a struct type and we are creating a zero value variable of type WaitGroup in line no.18. The way
	WaitGroup works is by using a counter. When we call Add on the WaitGroup and pass it an int, the WaitGroup's
	counter is incremented by the value passed to Add. The way to decrement the counter is by calling Done() method on
	the WaitGroup. The Wait() method blocks the Goroutine in which it's called until the counter becomes zero.
	In the above program, we call wg.Add(1) inside the for loop which iterates 3 times. So the counter now becomes 3.
	The for loop also spawns 3 process Goroutines and then wg.Wait() called makes the main Goroutine to wait until the
	counter becomes zero. The counter is decremented by the call to wg.Done in the process Goroutine. Once all the 3
	spawned Goroutines finish their execution, that is once wg.Done() has been called three times, the counter will
	become zero, and the main Goroutine will be unblocked.
	 */
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		/*
		It is important to pass the address of wg. If the address is not passed, then each Goroutine will have its
		own copy of the WaitGroup and main will not be notified when they finish executing.
		 */
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished executing")

	fmt.Printf("\nBeginning of worker pool implementation...\n")
	/*
	One of the important uses of buffered channel is the implementation of worker pool.
	In general, a worker pool is a collection of threads which are waiting for tasks to be assigned to them. Once they
	finish the task assigned, they make themselves available again for the next task.
	We will implement worker pool using buffered channels. Our worker pool will carry out the task of finding the sum of
	a digits of the input number. For example if 234 is passed, the output would be 9 (2 + 3 + 4). The input to the worker
	pool will be list of pseudo random integers.
	The following are the core functionalities of our worker pool:
		- Creation of a pool of Goroutines which listen on an input buffered channel waiting for jobs to be assigned
		- Addition of jobs to the input buffered channel
		- Writing results to an output buffered channel after job completion
		- Read and print results from the output buffered channel
	Check the worker_pools_demo.go for the code.
	*/
}