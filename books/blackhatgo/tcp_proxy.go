package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type fooReader struct {}

func (fooReader *fooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type fooWriter struct {}

func (fooWriter *fooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func readAndWrite(writer *fooWriter, reader *fooReader) {
	// Create buffer to hold input/output
	input := make([]byte, 4096)
	// Use reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)
	// Use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}

func readAndWriteWithCopying(writer *fooWriter, reader *fooReader) {
	if _, err := io.Copy(writer, reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

// echo is a handler function that simply echoes received data.
func echo(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store received data
	b := make([]byte, 512)
	for {
		// Receive data via conn.Read into a buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		// Send data via conn.Write
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func improvedEcho(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s\n", len(s), s)

	log.Println("Writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
}

func runEchoServer() {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for  {
		// Wait for connection. Create net.Conn on connection established. It blocks execution as it awaits client connections
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency
		go echo(conn)
	}
}

func runImprovedEchoServer() {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for  {
		// Wait for connection. Create net.Conn on connection established. It blocks execution as it awaits client connections
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency
		go improvedEcho(conn)
	}
}

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "joescatcam.website:80")
	if err != nil {
		log.Fatalln("Unable to connect to our reachable host")
	}
	defer dst.Close()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func runJoesProxyCom() {
	// Listen on local port 80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}

func runTcpProxy() {
	// Introduction
	/*
	You can achieve all TCP-based communications by using Go’s built-in net package. The previous section(port scanner)
	focused primarily on using the net package from a client’s perspective, and this section will use it to create TCP
	servers and transfer data. You’ll begin this journey by building the requisite echo server—a server that merely
	echoes a given response back to a client—followed by two much more generally applicable programs: a TCP port forwarder
	and a re-creation of Netcat’s “gaping security hole” for remote command execution.
	*/




	// Using io.Reader and io.Writer
	/*
	To create the examples in this section, you need to use two significant types that are crucial to essentially all
	input/output (I/O) tasks, whether you’re using TCP, HTTP, a filesystem, or any other means: io.Reader and io.Writer.
	Part of Go’s built-in io package, these types act as the cornerstone to any data transmission, local or networked.
	These types are defined in Go’s documentation as follows:
		type Reader interface {
	    	Read(p []byte) (n int, err error)
		}
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	Both types are defined as interfaces, meaning they can’t be directly instantiated. Each type contains the definition
	of a single exported function: Read or Write. As explained in Chapter 1, you can think of these functions as abstract
	methods that must be implemented on a type for it to be considered a Reader or Writer. For example, the following
	contrived type fulfills this contract and can be used anywhere a Reader is accepted:
		type FooReader struct {}
		func (fooReader *FooReader) Read(p []byte) (int, error) {
			// Read some data from somewhere, anywhere.
			return len(dataReadFromSomewhere), nil
		}
	Let’s take this knowledge and create something semi-usable: a custom Reader and Writer that wraps stdin and stdout.
	The code for this is a little contrived since Go’s os.Stdin and os.Stdout types already act as Reader and Writer,
	but then you wouldn’t learn anything if you didn’t reinvent the wheel every now and again, would you?
	*/
	/*var (
		reader fooReader
		writer fooWriter
	)*/
	// readAndWrite(&writer, &reader)
	/*
	In the above readAndWrite() function, the data itself is copied into the byte slice passed to the function. This is
	consistent with the Reader interface prototype definition provided earlier in this section.
	Copying data from a Reader to a Writer is a fairly common pattern—so much so that Go’s io package contains a Copy()
	function that can be used to simplify the main() function. The function prototype is as follows:
		func Copy(dst io.Writer, src io.Reader) (written int64, error)
	This convenience function allows you to achieve the same programmatic behavior as before with below function!
	*/
	// readAndWriteWithCopying(&writer, &reader)
	/*
	Notice that the explicit calls to reader.Read([]byte) and writer.Write([]byte) have been replaced with a single call
	to io.Copy(writer, reader) inside function readAndWriteWithCopying(). Under the covers, io.Copy(writer, reader) calls
	the Read([]byte) function on the provided reader, triggering the FooReader to read from stdin. Subsequently,
	io.Copy(writer, reader) calls the Write([]byte) function on the provided writer, resulting in a call to your
	FooWriter, which writes the data to stdout. Essentially, io.Copy(writer, reader) handles the sequential read-then-write
	process without all the petty details.
	*/




	// Creating the Echo Server
	/*
	As is customary for most languages, you’ll start by building an echo server to learn how to read and write data to
	and from a socket. To do this, you’ll use net.Conn, Go’s stream-oriented network connection, which we introduced when
	you built a port scanner. Based on Go’s documentation for the data type, Conn implements the Read([]byte) and
	Write([]byte) functions as defined for the Reader and Writer interfaces. Therefore, Conn is both a Reader and a
	Writer (yes, this is possible). This makes sense logically, as TCP connections are bidirectional and can be used to
	send (write) or receive (read) data.
	After creating an instance of Conn, you’ll be able to send and receive data over a TCP socket. However, a TCP server
	can’t simply manufacture a connection; a client must establish a connection. In Go, you can use net.Listen(network,
	address string) to first open a TCP listener on a specific port. Once a client connects, the Accept() method creates
	and returns a Conn object that you can use for receiving and sending data.
	*/
	// runEchoServer()
	/*
	echo(net.Conn), which accepts a Conn instance as a parameter. It behaves as a connection handler to perform all
	necessary I/O. The function loops indefinitely, using a buffer to read and write data from and to the connection.
	The data is read into a variable named b and subsequently written back on the connection.
	Now you need to set up a listener that will call your handler. As mentioned previously, a server can’t manufacture
	a connection but must instead listen for a client to connect. Therefore, a listener, defined as tcp bound to port
	20080, is started on all interfaces by using the net.Listen(network, address string) function.
	Next, an infinite loop ensures that the server will continue to listen for connections even after one has been
	received. Within this loop, you call listener.Accept(), a function that blocks execution as it awaits client
	connections. When a client connects, this function returns a Conn instance. Recall from earlier discussions in this
	section that Conn is both a Reader and a Writer (it implements the Read([]byte) and Write([]byte) interface methods).
	The Conn instance is then passed to the echo(net.Conn) handler function. This call is prefaced with the go keyword,
	making it a concurrent call so that other connections don’t block while waiting for the handler function to complete.
	This is likely overkill for such a simple server, but we’ve included it again to demonstrate the simplicity of Go’s
	concurrency pattern, in case it wasn’t already clear. At this point, you have two lightweight threads running
	concurrently:
		- The main thread loops back and blocks on listener.Accept() while it awaits another connection.
		- The handler goroutine, whose execution has been transferred to the echo(net.Conn) function, proceeds to run,
		processing the data.

	*/




	// Improving the Code by Creating a Buffered Listener
	/*
	Above function runEchoServer() works perfectly fine but relies on fairly low-level function calls, buffer tracking,
	and iterative reads/writes. This is a somewhat tedious, error-prone process. Fortunately, Go contains other packages
	that can simplify this process and reduce the complexity of the code. Specifically, the bufio package wraps Reader
	and Writer to create a buffered I/O mechanism. The updated version of echo(net.Conn) function is detailed here, and
	an explanation of the changes follows.
	*/
	// runImprovedEchoServer()
	/*
	No longer are you directly calling the Read([]byte) and Write([]byte) functions on the Conn instance; instead, you’re
	initializing a new buffered Reader and Writer via NewReader(io.Reader) and NewWriter(io.Writer). These calls both
	take, as a parameter, an existing Reader and Writer (remember, the Conn type implements the necessary functions to
	be considered both a Reader and a Writer).
	Both buffered instances contain complementary functions for reading and writing string data. ReadString(byte) takes
	a delimiter character used to denote how far to read, whereas WriteString(byte) writes the string to the socket.
	When writing data, you need to explicitly call writer.Flush() to flush write all the data to the underlying writer
	(in this case, a Conn instance).
	Although the previous example simplifies the process by using buffered I/O, you can reframe it to use the
	Copy(Writer, Reader) convenience function. Recall that this function takes as input a destination Writer and a
	source Reader, simply copying from source to destination.
	In this example, you’ll pass the conn variable as both the source and destination because you’ll be echoing the
	contents back on the established connection:
		func echo(conn net.Conn) {
	    	defer conn.Close()
	    	// Copy data from io.Reader to io.Writer via io.Copy().
			if _, err := io.Copy(conn, conn); err != nil {
				log.Fatalln("Unable to read/write data")
			}
		}
	*/




	// Proxying a TCP Client
	/*
	Now that you have a solid foundation, you can take what you’ve learned up to this point and create a simple port
	forwarder to proxy a connection through an intermediary service or host. As mentioned earlier in this chapter, this
	is useful for trying to circumvent restrictive egress controls or to leverage a system to bypass network segmentation.
	Before laying out the code, consider this imaginary but realistic problem: Joe is an underperforming employee who
	works for ACME Inc. as a business analyst making a handsome salary based on slight exaggerations he included on his
	resume. (Did he really go to an Ivy League school? Joe, that’s not very ethical.) Joe’s lack of motivation is
	matched only by his love for cats—so much so that Joe installed cat cameras at home and hosted a site,
	joescatcam.website, through which he could remotely monitor the dander-filled fluff bags. One problem, though: ACME is
	onto Joe. They don’t like that he’s streaming his cat cam 24/7 in 4K ultra high-def, using valuable ACME network
	bandwidth. ACME has even blocked its employees from visiting Joe’s cat cam website.
	Joe has an idea. “What if I set up a port-forwarder on an internet-based system I control,” Joe says, “and force the
	redirection of all traffic from that host to joescatcam.website?” Joe checks at work the following day and confirms
	he can access his personal website, hosted at the joesproxy.com domain. Joe skips his afternoon meetings, heads to a
	coffee shop, and quickly codes a solution to his problem. He’ll forward all traffic received at http://joesproxy.com
	to http://joescatcam.website.
	Here’s Joe’s code, which he runs on the joesproxy.com server:
	*/
	runJoesProxyCom()
	/*
	Start by examining Joe’s handle(net.Conn) function. Joe connects to joescatcam.website (recall that this unreachable
	host isn’t directly accessible from Joe’s corporate workstation). Joe then uses Copy(Writer, Reader) two separate
	times. The first instance ensures that data from the inbound connection is copied to the joescatcam.website connection.
	The second instance ensures that data read from joescatcam.website is written back to the connecting client’s
	connection. Because Copy(Writer, Reader) is a blocking function, and will continue to block execution until the
	network connection is closed, Joe wisely wraps his first call to Copy(Writer, Reader) in a new goroutine. This
	ensures that execution within the handle(net.Conn) function continues, and the second Copy(Writer, Reader) call can
	be made.
	Joe’s proxy listens on port 80 and relays any traffic received from a connection to and from port 80 on
	joescatcam.website. Joe, that crazy and wasteful man, confirms that he can connect to joescatcam.website via
	joesproxy.com by connecting with curl.
	At this point, Joe has done it. He’s living the dream, wasting ACME-sponsored time and network bandwidth while he
	watches his cats. Today, there will be cats!
	*/




	// Replicating Netcat for Command Execution
	/*

	*/
}