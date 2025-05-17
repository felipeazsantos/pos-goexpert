package main

import "time"

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 4)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 4)
		c2 <- 2
	}()

	select {
	case v := <-c1:
		println("c1", v)
	case v := <-c2:
		println("c2", v)
	case <-time.After(time.Second * 3):
		println("timeout")
	}
}	

