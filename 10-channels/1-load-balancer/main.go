package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}


func main() {
	data := make(chan int)
	qtdWorkers := 1000

	for i := 1; i <= qtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 1; i <= 10000; i++ {
		data <- i
	}

	close(data)
}