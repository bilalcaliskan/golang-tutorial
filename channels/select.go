package channels

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func processForSelect(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successfull"
}

func RunSelect() {
	fmt.Printf("\nBeginning of introduction...\n")
	fmt.Println("The select statement is used to choose from multiple send/receive channel operations. The select " +
		"statement blocks until one of the send/receive operation is ready. If multiple operations are ready, one of " +
		"them is chosen at random. The syntax is similar to switch_statement except that each of the case statement will be a " +
		"channel operation.")
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <- output1:
		fmt.Println(s1)
	case s2 := <- output2:
		fmt.Println(s2)
	}

	fmt.Printf("\nBeginning of practical use of select...\n")
	fmt.Println("Lets assume we have a mission critical application and we need to return the output to the user " +
		"as quickly as possible. The database for this application is replicated and stored in different servers " +
		"across the world. Assume that the functions server1 and server2 are in fact communicating with 2 such servers. " +
		"The response time of each server is dependant on the load on each and the network delay. We send the request " +
		"to both the servers and then wait on the corresponding channels for the response using the select statement. " +
		"The server which responds first is chosen by the select and the other response is ignored. This way we can " +
		"send the same request to multiple servers and return the quickest response to the user.")

	fmt.Printf("\nBeginning of deadlock and default case...\n")
	// We have commented out the below block of codes because The select statement will block forever since no
	//other Goroutine is writing to this channel and hence will result in deadlock. This program will panic at runtime.
	// If a default case is present, this deadlock will not happen since the default case will be executed when no
	//other case is ready.
	/*ch := make(chan string)
	select {
	case <- ch:
	}*/
	ch := make(chan string)
	select {
	case <- ch:
	default:
		fmt.Println("default case executed to prevent deadlock")
	}
	fmt.Println("Similarly the default case will be executed even if the select has only nil channels")
	var nilChannel chan string
	select {
	case v := <- nilChannel:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")
	}

	fmt.Printf("\nBeginning of empty select...\n")
	// The below line commented out because the select statement will block until one of its cases is executed. In this
	// case, the select statement does not have any cases and hence it will block forever resulting in a deadlock.
	// this program will panic with the following output: fatal error: all goroutines are asleep - deadlock!
	/*
	select {

	}
	 */

	fmt.Printf("\nBeginning of default case...\n")
	ch2 := make(chan string)
	go processForSelect(ch2)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <- ch2:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}