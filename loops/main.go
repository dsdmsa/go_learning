package main

import (
	"fmt"
	"time"
)

var chan1 = make(chan string)
var chan2 = make(chan string)

func task1() {
	time.Sleep(1 * time.Second)
	chan1 <- "msg"
}

func task2() {
	time.Sleep(2 * time.Second)
	chan2 <- "msg2"
}

func main() {

	go task1()
	go task2()

	for i := 0; i < 2; i++ {

		select {
		case msg := <-chan1:
			fmt.Println("recieved ", msg)
		case msg2 := <-chan2:
			fmt.Println("recieved ", msg2)
		}

	}

}
