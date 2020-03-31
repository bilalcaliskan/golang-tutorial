package concurrency

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
	fmt.Printf("\nBeginning of introduction to channels...\n")
	/*
	Channels can be thought as pipes using which Goroutines communicate. Similar to how water flows from one end to
	another in a pipe, data can be sent from one end and received from the another end using channels.
	 */

	fmt.Printf("\nBeginning of declaring channels...\n")
	/*
	Each channel has a type associated with it. This type is the type of data that the channel is allowed to transport.
	No other type is allowed to be transported using the channel.
		chan T is a channel of type T
	The zero value of a channel is nil. nil channels are not of any use and hence the channel has to be defined using
	make similar to maps and slices.
	 */
	var a chan int // this is a nil channel
	if a == nil {
		fmt.Println("Channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
	/*
	As usual the short hand declaration is also a valid and concise way to define a channel.
	 */
	b := make(chan int)
	if b == nil {
		fmt.Println("Channel b is nil, going to define it")
	} else {
		fmt.Println("Channel b is not nil")
	}

	fmt.Printf("\nBeginning of sending and receiving from channel...\n")
	/*
	The syntax to send and receive data from a channel are given below:
		data := <- a // read from channel a
		a <- data // write to channel a
	The direction of the arrow with respect to the channel specifies whether the data is sent or received.
	Sends and receives to a channel are blocking by default. When a data is sent to a channel, the control is blocked
	in the send statement until some other Goroutine reads from that channel. Similarly when data is read from a channel,
	the read is blocked until some Goroutine writes data to that channel. This property of channels is what helps
	Goroutines communicate effectively without the use of explicit locks or conditional variables that are quite common
	in other programming languages.
	*/

	fmt.Printf("\nBeginning of example program...\n")
	done := make(chan bool)
	go helloWorldWithoutSleep(done)
	<- done
	/*
	Above line of code is blocking which means that until some Goroutine writes data to the done channel, the control
	will not move to the next line of code. It blocks main goroutine in that case. Hence this eliminates the need for
	the time.Sleep natively which was present in the original program to prevent the main Goroutine from exiting.
	The line of code <-done receives data from the done channel but does not use or store that data in any variable.
	This is perfectly legal.
	Now we have our main Goroutine blocked waiting for data on done channel. The hello Goroutine receives this channel
	as parameter, prints Hello world goroutine and then writes to the done channel. When this write is complete, the
	main Goroutine receives the data from the done channel, it is unblocked and then the text main goroutine is printed.
	 */
	fmt.Println("main goroutine")

	fmt.Printf("\nBeginning of another example program...\n")
	/*
	We will structure the program such that the squares are calculated in a separate Goroutine, cubes in another
	Goroutine and the final summation happens in the main Goroutine.
	*/
	var number = 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <- sqrch, <- cubech
	fmt.Println("Final output", squares + cubes)

	fmt.Printf("\nBeginning of a deadlocks...\n")
	/*
	One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then
	it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will
	panic at runtime with Deadlock.
	Similarly if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write
	data on that channel, else the program will panic.
	 */
	// commented below 2 line to get rid of deadlock error
	// ch := make(chan int)
	// ch <- 5
	/*
	In the program above, a channel ch is created and we send 5 to the channel in line ch <- 5. In this program no other
	Goroutine is receiving data from the channel ch. Hence this program will panic with the runtime error.
	 */

	fmt.Printf("\nBeginning of unidirectional channels...\n")
	/*
	All the channels we discussed so far are bidirectional channels, that is data can be both sent and received on them.
	It is also possible to create unidirectional channels, that is channels that only send or receive data.
	 */
	chnl := make(chan <- int)
	go sendData(chnl)
	/*
	In the above program, we create send only channel sendch. chan <- int denotes a send only channel as the arrow is
	pointing to chan. We try to receive data from a send only channel in below line. This is not allowed and when the
	program is run, the compiler will complain stating,
	 */
	// fmt.Println(<-chnl)
	/*
	All is well but what is the point of writing to a send only channel if it cannot be read from!
	This is where channel conversion comes into use. It is possible to convert a bidirectional channel to a send only or
	receive only channel but not the vice versa.
	 */
	chnl2 := make(chan int)
	go sendData(chnl2)
	fmt.Println(<- chnl2)
	/*
	In the above program, a bidirectional channel chnl is created. It is passed as a parameter to the sendData Goroutine.
	The sendData function converts this channel to a send only channel in line no. 5 in the parameter
	sendch chan<- int. So now the channel is send only inside the sendData Goroutine but it's bidirectional in the main
	Goroutine. This program will print 10 as the output.
	 */

	fmt.Printf("\nBeginning of closing channels and for range loops on channels...\n")
	/*
	Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.
	Receivers can use an additional variable while receiving data from the channel to check whether the channel has
	been closed.
		v, ok := <- ch
	In the above statement ok is true if the value was received by a successful send operation to a channel. If ok is
	false it means that we are reading from a closed channel. The value read from a closed channel will be the zero
	value of the channel's type. For example if the channel is an int channel, then the value received from a closed
	channel will be 0.
	*/
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
	/*
	The for range form of the for loop can be used to receive values from a channel until it is closed.
	Below code will be exit when the ch2 channel closed on the producer function.
	*/
	ch2 := make(chan int)
	go producer(ch2)
	for v := range ch2 {
		fmt.Println("Received", v)
	}
	/*
	In the above program, for range loop receives data from the ch2 channel until it is closed. Once ch2 is closed, the
	loop automatically exits.
	 */
}