package main

// Scanning TCP ports range from 1 to 1024.
// For every port, if the number of the port is not 1024, keep scanning.
// Now we have a int (port number), that we need to convert using Sprintf()
// to use with the second argument string Dial(network, address string).
// This scan is really, really slow

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ { // i = 1, "while" i is not 1024 keep increasing its value
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered
			continue // keep scanning even if the connection is not successful on the current port
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
