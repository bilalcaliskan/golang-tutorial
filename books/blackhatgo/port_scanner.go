package main

import (
	"fmt"
	"net"
)

func runPortScanner() {
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
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successfull!")
	} else {
		fmt.Println("An error occured while establishing connection!")
	}

	/*
	Scanning a single port at a time isn’t useful, and it certainly isn’t efficient. TCP ports range from 1 to 65535; but
	for testing, let’s scan ports 1 to 1024.
	*/
}