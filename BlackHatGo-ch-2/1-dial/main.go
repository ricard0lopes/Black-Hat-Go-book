// Testing Port Availability.
// Initiate connection from client to server.
// Connecting and scanning scanme.nmap-org
// using Go's net package: net.Dial(network, address string).

package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80") // returns Conn and error
	if err == nil {
		// error will be nil if the connection is successful
		fmt.Println("Connection successful")
	}
}
