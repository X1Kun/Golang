package main

import (
	"fmt"
)

func fpn(c chan int, quit chan int) {
	a, b := 1, 1

	for {
		select {
		case c <- a:
			temp := a + b
			a = b
			b = temp
			// fmt.Println("in", a)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// <-c
			fmt.Println("out", <-c)
		}

		quit <- 1
	}()
	fpn(c, quit)

}
