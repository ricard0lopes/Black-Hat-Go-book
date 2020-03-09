/*
  FINAL IMPROVEMENTS:
  In order to sort the ports before printing w need to use a separate thread
  to pass the result of the port scan back to the main thread.
  A benefit of this modificaton is that we can remove the dependency of a WaitGroup entirely.

  Added to the program:
    - Commandline arguments
    - Name of common ports' version
  Usage:
    $ go build main.go
    $ ./main <address>
*/

package main

import (
	"fmt"
	"net"
	"os"
	"sort"
)

func worker(ports, results chan int) {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "ADDRESS")
		return
	} else {
		for p := range ports {
			cmd := os.Args[1]
			address := fmt.Sprintf("%v:%d", cmd, p)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				results <- 0
				continue
			}
			conn.Close()
			results <- p
		}
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		var service string
		switch port {
		case 22:
			service = "SSH"
		case 80:
			service = "HTTP"
		case 443:
			service = "HTTPS"
		case 21:
			service = "FTP"
		case 25:
			service = "SMTP"
		case 110:
			service = "POP3"
		case 143:
			service = "IMAP"
		case 3306:
			service = "MySQL"
		default:
			service = "Unknown Service"
		}

		fmt.Printf("Port %d %v is open\n", port, service)
	}
}
