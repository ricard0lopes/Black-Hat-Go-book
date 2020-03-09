/*
  To avoid inconsistencies we can use a pool of goroutines to manage
  the concurrent work being performed.
  We can create a number of worker goroutines as a resource pool with a for loop.
  Then we will use a channel on our main() function to provide work.
  To start we create a new program that has 100 workers, consumes a
  channel of int, and prints them to the screen.
  We still use the WaitGroup to block execution.
*/

package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) { // The channel(chan) will be used to receive the work
	// and the WaitGroup to track when a single work item as been completed
	for p := range ports { // use range to continuously receive rom the ports channel
		// looping until the channel is closed.
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100) // Create a channel with an int value of 100
	// meaning it can hold 100 items before the sender will block, allowing all the workers to start immediatly.
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ { // start the desired number of workers
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i // send a port on the ports channel to the worker
	}
	wg.Wait()
	close(ports) // after all the work is completed we close the channel
}
