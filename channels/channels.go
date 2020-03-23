package channels

import (
	"fmt"
	"time"
)

func sendData(sendch chan<- int) {
	sendch <- 10
}

func helloWorldWithoutSleep(done chan bool) {
	fmt.Println("helloWorldWithoutSleep goroutine")
	done <- true
}

func helloWorldWithSleep(done2 chan bool) {
	fmt.Println("helloWorldWithSleep goroutine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("helloWorldWithSleep goroutine awake and going to write to done2 channel")
	done2 <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func RunChannels() {
	fmt.Printf("\nBeginning of declaring channels...\n")
	var a chan int
	if a == nil {
		fmt.Println("Channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
	fmt.Println("Sends and receives to a channel are blocking by default. What does this mean? When a data is " +
		"sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that " +
		"channel. Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to " +
		"that channel.")
	fmt.Println("This property of channels is what helps Goroutines communicate effectively without the use of " +
		"explicit locks or conditional variables that are quite common in other programming languages.")

	fmt.Printf("\nBeginning of example program...\n")
	done := make(chan bool)
	go helloWorldWithoutSleep(done)
	<- done
	fmt.Println("main goroutine")

	fmt.Printf("\nBeginning of another example program...\n")
	done2 := make(chan bool)
	fmt.Println("Main goroutine going to call helloWorld go goroutine")
	go helloWorldWithSleep(done2)
	<- done2
	fmt.Println("Main goroutine received data")

	fmt.Printf("\nBeginning of a little bit complex program...\n")
	var number = 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <- sqrch, <- cubech
	fmt.Println("Final output", squares + cubes)

	fmt.Printf("\nBeginning of a deadlocks...\n")
	fmt.Println("One important factor to consider while using channels is deadlock. If a Goroutine is sending " +
		"data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does " +
		"not happen, then the program will panic at runtime with Deadlock.")
	fmt.Println("Similarly if a Goroutine is waiting to receive data from a channel, then some other Goroutine " +
		"is expected to write data on that channel, else the program will panic.")
	// commented below 2 line to get rid of deadlock error
	// ch := make(chan int)
	// ch <- 5

	fmt.Printf("\nBeginning of unidirectional channels...\n")
	fmt.Println("All the channels we discussed so far are bidirectional channels, that is data can be both sent " +
		"and received on them. It is also possible to create unidirectional channels, that is channels that only " +
		"send or receive data.")
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)

	fmt.Printf("\nBeginning of closing channels and for range loops on channels...\n")
	fmt.Println("Senders have the ability to close the channel to notify receivers that no more data will be " +
		"sent on the channel. Receivers can use an additional variable while receiving data from the channel to " +
		"check whether the channel has been closed.")
	fmt.Println("v, ok := <- ch")
	fmt.Println("In the above statement ok is true if the value was received by a successful send operation " +
		"to a channel. If ok is false it means that we are reading from a closed channel. The value read from a " +
		"closed channel will be the zero value of the channel's type. For example if the channel is an int channel, " +
		"then the value received from a closed channel will be 0.")
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <- ch
		if ok == false {
			fmt.Println("Received", v, ok)
			break
		}
		fmt.Println("Received", v, ok)
	}
	fmt.Println("Below code will be exit when the ch2 channel closed on the producer function")
	ch2 := make(chan int)
	go producer(ch2)
	for v := range ch2 {
		fmt.Println("Received", v)
	}
}