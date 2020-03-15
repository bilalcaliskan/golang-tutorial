package goroutines

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

func RunBufferedPools() {
	fmt.Printf("\nBeginning of description of buffered channels...\n")
	fmt.Println("All the channels we discussed in the previous tutorial were basically unbuffered. As we " +
		"discussed in the channels tutorial in detail, sends and receives to an unbuffered channel are blocking.")
	fmt.Println("It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only " +
		"when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.")
	fmt.Println("Buffered channels can be created by passing an additional capacity parameter to the make " +
		"function which specifies the size of the buffer. ch := make(chan type, capacity)")
	fmt.Println("capacity in the above syntax should be greater than 0 for a channel to have a buffer. The " +
		"capacity for an unbuffered channel is 0 by default and hence we omitted the capacity parameter while " +
		"creating channels")
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<- ch)
	fmt.Println(<- ch)

	fmt.Printf("\nBeginning of another example on buffered channels...\n")
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
	/**
	If we uncomment below line, there will be a deadlock. ch3 channel has capacity 2 and to be able to write that
	channel, somebody must read from that channel. But there is no concurrent routine reading from that channel.
	 */
	//ch3 <- "steve"
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
	fmt.Println("To understand worker pools, we need to first understand WaitGroup.")
	fmt.Println("A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control " +
		"is blocked until all Goroutines finish executing. Lets say we have 3 concurrently executing Goroutines " +
		"spawned from the main Goroutine. The main Goroutines needs to wait for the 3 other Goroutines to finish " +
		"before terminating. This can be accomplished using WaitGroup.")
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished executing")
}