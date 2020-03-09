/*
  Using WaitGroup from sync package will allow us to control concurrency.
  We can create the WaitGroup struct with 'var wg sync.WaitGroup'.
  Once created we can call a few methods on the struct:
    1. Add(int) -> increases an internal counter by the number provided.
    2. Done() -> decrements the counter by one.
    3. Wait() -> blocks the execution of the goroutine in which it's called,
    and will not allow further execution until the internal counter reaches zero.
  Now we can create a synchronized scanning using WaitGroup
*/

package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup // synchronized counter
	for i := 1; i <= 65535; i++ { // decided to increase the ports range for experimental purposes
		wg.Add(1) // increment the counter each time it creates a goroutine to scan a port
		go func(j int) {
			defer wg.Done() // decrements the counter whenever one unit of work has been performed
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait() // blocks until all the work has been done and the counter returned to zero
}
