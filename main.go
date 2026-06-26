package main

import (
	"fmt"
	"net"
	"time"
)

func scanPort(port int, results chan<- int) {
	address := fmt.Sprintf("scanme.nmap.org:%d", port) // სამიზნე მისამართი
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		results <- 0 // პორტი დახურულია
		return
	}
	conn.Close()
	results <- port // პორტი ღიაა
}

func main() {
	results := make(chan int)

	// ვუშვებთ სკანირებას 1-დან 1024-მდე პორტებზე
	for i := 1; i <= 1024; i++ {
		go scanPort(i, results)
	}

	for i := 1; i <= 1024; i++ {
		port := <-results
		if port != 0 {
			fmt.Printf("პორტი %d ღიაა!\n", port)
		}
	}
}