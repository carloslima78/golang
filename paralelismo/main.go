package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", workerId, x)
	}
}

func main() {
	ch := make(chan int)

	// go worker(1, ch)
	// go worker(2, ch)

	for i := 0; i < 10; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 100; i++ {
		ch <- i
	}
}
