package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

func singlePortScan(host string, proto string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	_, err := net.Dial(proto, address)
	if err == nil {
		fmt.Println("Connection successfull!")
	}
}

func multiPortScan(host string, proto string) {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("%s:%d", host, i)
		conn, err := net.Dial(proto, address)
		if err != nil {
			// port is closed or filtered
			continue
		}
		conn.Close()
		fmt.Printf("%d open on host %s\n", i, host)
	}
}

// The “Too Fast” Scanner Version
func multiPortScanConcurrently(host string, proto string, wg *sync.WaitGroup) {
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", host, j)
			conn, err := net.Dial(proto, address)
			if err != nil {
				// port is closed or filtered
				return
			}
			conn.Close()
			fmt.Printf("%d open on host %s\n", j, host)
		}(i)
	}
}

func multiPortScanConcurrentlyUsingWorkerPool(host string, proto string, wg *sync.WaitGroup) {
	ports := make(chan int, 100)
	for i := 0; i < cap(ports); i++ {
		go worker(host, proto, ports, wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}

func worker(host string, proto string, ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Printf("%s:%d\n", host, p)
		wg.Done()
	}
}

func multiPortScanConcurrentlyUsingMultipleChannels(host, proto string) {
	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go workerMultiChannel(host, proto, ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <- results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}

func workerMultiChannel(host, proto string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial(proto, address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func runPortScanner() {
	proto := "tcp"
	host := "scanme.nmap.org"
	// Scanning single port
	/*
	The first step in creating the port scanner is understanding how to initiate a connection from a client to a server.
	Throughout this example, you’ll be connecting to and scanning scanme.nmap.org, a service run by the Nmap project.
	To do this, you’ll use Go’s net package: net.Dial(network, address string).
	The first argument is a string that identifies the kind of connection to initiate. This is because Dial isn’t just
	for TCP; it can be used for creating connections that use Unix sockets, UDP, and Layer 4 protocols that exist only
	in your head (the authors have been down this road, and suffice it to say, TCP is very good). There are a few strings
	you can provide, but for the sake of brevity, you’ll use the string tcp.
	The second argument tells Dial(network, address string) the host to which you wish to connect. Notice it’s a single
	string, not a string and an int. For IPv4/TCP connections, this string will take the form of host:port. For example,
	if you wanted to connect to scanme.nmap.org on TCP port 80, you would supply scanme.nmap.org:80.

	Now you know how to create a connection, but how will you know if the connection is successful? You’ll do this through
	error checking: Dial(network, address string) returns Conn and error, and error will be nil if the connection is
	successful. So, to verify your connection, you just check whether error equals nil.
	*/
	// singlePortScan(host, proto, 80)




	// Scanning Multiple Ports
	/*
	Scanning a single port at a time isn’t useful, and it certainly isn’t efficient. TCP ports range from 1 to 65535; but
	for testing, let’s scan ports 1 to 1024.
	*/
	// multiPortScan(host, proto)




	// Synchronized Scanning Using WaitGroup
	/*
	The previous scanner scanned multiple ports in a single go (pun intended). But your goal now is to scan multiple
	ports concurrently, which will make your port scanner faster. To do this, you’ll harness the power of goroutines.
	Go will let you create as many goroutines as your system can handle, bound only by available memory.
	The most naive way to create a port scanner that runs concurrently is to wrap the call to Dial(network, address string)
	in a goroutine.
	*/
	/*
	If we dont define a wait group, the code you just ran launches a single goroutine per connection, but the main goroutine
	who calls the method does not wait for child goroutines to complete. Therefore, the code completes and exits as soon
	as the for loop finishes its iterations, which may be faster than the network exchange of packets between your code
	and the target ports. You may not get accurate results for ports whose packets were still in-flight.
	There are a few ways to fix this. One is to use WaitGroup from the sync package, which is a thread-safe way to control
	concurrency. WaitGroup is a struct type and can be created like so:
		var wg sync.WaitGroup
	After created that, we will pass it to the function
	Once you’ve created WaitGroup, you can call a few methods on the struct. The first is Add(int), which increases an
	internal counter by the number provided. Next, Done() decrements the counter by one. Finally, Wait() blocks the
	execution of the goroutine in which it’s called, and will not allow further execution until the internal counter
	reaches zero. You can combine these calls to ensure that the main goroutine waits for all connections to finish.
	*/
	// var wg sync.WaitGroup
	// multiPortScanConcurrently(host, proto, &wg)
	// wg.Wait()
	/*
	Above function remains largely identical to our initial version. However, you’ve added code that explicitly tracks
	the remaining work. In this version of the program, you create sync.WaitGroup which acts as a synchronized counter.
	You increment this counter via wg.Add(1) each time you create a goroutine to scan a port and a deferred call to wg.Done()
	decrements the counter whenever one unit of work has been performed. Your main() function calls wg.Wait(), which blocks
	until all the work has been done and your counter has returned to zero.

	This version of the program is better, but still incorrect. If you run this multiple times against multiple hosts, you
	might see inconsistent results. Scanning an excessive number of hosts or ports simultaneously may cause network or
	system limitations to skew your results. Go ahead and change 1024 to 65535, and the destination server to your
	localhost 127.0.0.1 in your code. If you want, you can use Wireshark or tcpdump to see how fast those connections
	are opened.
	*/




	// Port Scanning Using a Worker Pool
	/*
	To avoid inconsistencies, you’ll use a pool of goroutines to manage the concurrent work being performed. Using a for
	loop, you’ll create a certain number of worker goroutines as a resource pool. Then, in your main() “thread,” you’ll
	use a channel to provide work.
	To start, create a new program that has 100 workers, consumes a channel of int, and prints them to the screen.
	The worker(int, *sync.WaitGroup) function takes two arguments: a channel of type int and a pointer to a WaitGroup.
	The channel will be used to receive work, and the WaitGroup will be used to track when a single work item has been
	completed.
	*/
	// var wg sync.WaitGroup
	// multiPortScanConcurrentlyUsingWorkerPool(host, proto, &wg)
	/*
	We have created a channel in function above like that;
		ports := make❷(chan int, 100)
	This allows the channel to be buffered, which means you can send it an item without waiting for a receiver to read
	the item. Buffered channels are ideal for maintaining and tracking work for multiple producers and consumers. You’ve
	capped the channel at 100, meaning it can hold 100 items before the sender will block. This is a slight performance
	increase, as it will allow all the workers to start immediately.
	Next, you use a for loop to start the desired number of workers—in this case, 100. In the worker(int, *sync.WaitGroup)
	function, you use range to continuously receive from the ports channel, looping until the channel is closed.
	After all the work has been completed, you close the channel. Waitgroups stops blocking the execution so compiler can
	reach the closing the channel line.
	You might notice something interesting here: the numbers are printed in no particular order. Welcome to the wonderful
	world of parallelism.
	*/




	// Multichannel Communication
	/*
	To complete the port scanner, you could plug in your code from earlier in the section, and it would work just fine.
	However, the printed ports would be unsorted, because the scanner wouldn’t check them in order.
	To solve this problem, you need to use a separate thread to pass the result of the port scan back to your main thread
	to order the ports before printing. Another benefit of this modification is that you can remove the dependency of a
	WaitGroup entirely, as you’ll have another method of tracking completion. For example, if you scan 1024 ports, you’re
	sending on the worker channel 1024 times, and you’ll need to send the result of that work back to the main thread 1024
	times. Because the number of work units sent and the number of results received are the same, your program can know
	when to close the channels and subsequently shut down the workers.
	*/
	multiPortScanConcurrentlyUsingMultipleChannels(host, proto)
	/*
	If the port is closed, you’ll send a zero, and if its open, you will send the port to the results channel inside
	function workerMultiChannel. Also, you create a separate channel to communicate the results from the worker to the
	main thread. You then use a slice to store the results so you can sort them later. Next, you need to send to the
	workers in a separate goroutine because the result-gathering loop needs to start before more than 100 items of work
	can continue.
	The result-gathering loop receives on the results channel 1024 times. If the port doesn’t equal 0, it’s appended to
	the slice. After closing the channels, you’ll use sort to sort the slice of open ports. All that’s left is to loop
	over the slice and print the open ports to screen.
	*/



	// Final Words
	/*
	There you have it: a highly efficient port scanner. Take some time to play around with the code—specifically, the number
	of workers. The higher the count, the faster your program should execute. But if you add too many workers, your results
	could become unreliable. When you’re writing tools for others to use, you’ll want to use a healthy default value that
	caters to reliability over speed. However, you should also allow users to provide the number of workers as an option.

	*/
}