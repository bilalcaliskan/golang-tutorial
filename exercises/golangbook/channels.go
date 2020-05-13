package main

import (
	"fmt"
	"time"
)

/*
We can specify a direction on a channel type thus restricting it to either sending or receiving. For example pinger's
function signature can be changed to this:
	func pinger(c chan<- string)
Now c can only be sent to. Attempting to receive from c will result in a compiler error.
A channel that doesn't have these restrictions is known as bi-directional. A bi-directional channel can be passed to a
function that takes send-only or receive-only channels, but the reverse is not true.
*/
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

/*
Similarly we can change printer to this:
	func printer(c <-chan string)
A channel that doesn't have these restrictions is known as bi-directional. A bi-directional channel can be passed to a
function that takes send-only or receive-only channels, but the reverse is not true.
*/
func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg, cap(c), len(c))
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func runChannels() {
	/*
	A channel type is represented with the keyword chan followed by the type of the things that are passed on the channel
	(in this case we are passing strings)
	The <- (left arrow) operator is used to send and receive messages on the channel. c <- "ping" means send "ping".
	msg := <- c means receive a message and store it in msg. The fmt line could also have been written like
	this: fmt.Println(<-c)
	Using a channel like this synchronizes the two goroutines. When pinger attempts to send a message on the channel it
	will wait until printer is ready to receive the message. (this is known as blocking)
	*/
	var c = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	/*
	It's also possible to pass a second parameter to the make function when creating a channel:
		c := make(chan int, 1)
	This creates a buffered channel with a capacity of 1. Normally channels are synchronous; both sides of the channel
	will wait until the other side is ready. A buffered channel is asynchronous; sending or receiving a message will
	not wait unless the channel is already full.
	*/

	var input string
	fmt.Scanln(&input)
}