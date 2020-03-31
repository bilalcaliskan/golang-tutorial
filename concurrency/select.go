package concurrency

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

func server3(ch chan string) {
	ch <- "from server3"
}

func server4(ch chan string) {
	ch <- "from server4"
}

func processForSelect(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successfull"
}

func RunSelect() {
	fmt.Printf("\nBeginning of introduction...\n")
	/*
	The select statement is used to choose from multiple send/receive channel operations. The select statement blocks
	until one of the send/receive operation is ready. If multiple operations are ready, one of them is chosen at random.
	The syntax is similar to switch except that each of the case statement will be a channel operation.
	*/
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
	/*
	In the program above, the server1 function sleeps for 6 seconds then writes the text from server1 to
	the channel ch. The server2 function sleeps for 3 seconds and then writes from server2 to the channel ch.
	The main function calls the go Goroutines server1 and server2.
	The control reaches the select statement. The select statement blocks until one of its cases is ready. In our
	program above, the server1 Goroutine writes to the output1 channel after 6 seconds whereas the server2 writes to
	the output2 channel after 3 seconds. So the select statement will block for 3 seconds and will wait for server2
	Goroutine to write to the output2 channel
	 */

	fmt.Printf("\nBeginning of practical use of select...\n")
	/*
	The reason behind naming the functions in the above program as server1 and server2 is to illustrate the practical
	use of select.
	Lets assume we have a mission critical application and we need to return the output to the user as quickly as
	possible. The database for this application is replicated and stored in different servers across the world. Assume
	that the functions server1 and server2 are in fact communicating with 2 such servers. The response time of each
	server is dependant on the load on each and the network delay. We send the request to both the servers and then
	wait on the corresponding channels for the response using the select statement. The server which responds first is
	chosen by the select and the other response is ignored. This way we can send the same request to multiple servers
	and return the quickest response to the user :).
	*/

	fmt.Printf("\nBeginning of default case...\n")
	/*
	The default case in a select statement is executed when none of the other case is ready. This is generally used to
	prevent the select statement from blocking.
	 */
	ch0 := make(chan string)
	go processSingle(ch0)
	outer:
		for {
			time.Sleep(1000 * time.Millisecond)
			select {
			case v := <- ch0:
				fmt.Println("received value: ", v)
				break outer
			default:
				fmt.Println("no value received")
			}
		}
	/*
	After calling the process Goroutine concurrently, an infinite for loop is started in the main Goroutine. The infinite
	loop sleeps for 1000 milliseconds (1 second) during the start of each iteration and them performs a select operation.
	During the first 10500 milliseconds, the first case of the select statement namely case v := <-ch: will not be ready
	since the process Goroutine will write to the ch channel only after 10500 milliseconds. Hence thedefault case will
	be executed during this time and the program will print no value received 10 times.
	After 10.5 seconds, the process Goroutine writes process successful to ch in line no. 10. Now the first case of the
	select statement will be executed and the program will print received value:  process successful and then it will
	terminate.
	 */

	fmt.Printf("\nBeginning of deadlock and default case...\n")
	/*
	In the program above, we have created a channel ch in line no. 4. We try to read from this channel inside the select
	in line no. 6. The select statement will block forever since no other Goroutine is writing to this channel and hence
	will result in deadlock. This program will panic at runtime
	 */
	//ch := make(chan string)
	//select {
	//case <- ch:
	//}
	/*
	If a default case is present, this deadlock will not happen since the default case will be executed when no other
	case is ready. The program above is rewritten with a default case below.
	 */
	ch := make(chan string)
	select {
	case <- ch:
	default:
		fmt.Println("default case executed to prevent deadlock")
	}
	/*
	Similarly the default case will be executed even if the select has only nil channels.
	 */
	var nilChannel chan string
	select {
	case v := <- nilChannel:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed to prevent deadlock")
	}
	/*
	In the program above nilChannel is nil and we are trying to read from nilChannel in the select. If the default case
	was not present, the select would have blocked forever and caused a deadlock. Since we have a default case inside
	the select, it will be executed
	 */

	fmt.Printf("\nBeginning of random selection...\n")
	/*
	When multiple cases in a select statement are ready, one of them will be executed at random.
	 */
	output3 := make(chan string)
	output4 := make(chan string)
	go server3(output3)
	go server4(output4)
	time.Sleep(1 * time.Second)
	select {
	case s3 := <- output3:
		fmt.Println(s3)
	case s4 := <- output4:
		fmt.Println(s4)
	}

	fmt.Printf("\nBeginning of empty select...\n")
	/*
	The below block of code commented out because the select statement will block until one of its cases is executed. In
	this case, the select statement does not have any cases and hence it will block forever resulting in a deadlock.
	This program will panic with the following output: fatal error: all goroutines are asleep - deadlock!
	 */
	//select {
	//
	//}
}