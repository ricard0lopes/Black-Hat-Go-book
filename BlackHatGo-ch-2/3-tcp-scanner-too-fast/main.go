/*
 Scan multiple ports concurrently using goroutines.
 Wrap the call Dial(network, address string) in a goroutine.
 This program will exit almost immediatly because the goroutine
 doesn't know to wait for the connection to take place
*/

package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
}
