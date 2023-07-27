package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan string, 1)

	ch1 <- 999
	ch2 <- "message"

	select {
	case msg := <-ch1:
		fmt.Printf("Channel 1: %d", msg)
	case msg := <-ch2:
		fmt.Printf("Channel 2: %v", msg)
	}
}
