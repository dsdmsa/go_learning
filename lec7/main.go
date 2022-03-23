package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	go listen()

	fmt.Println("press any key or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' {
			break
		}
		keyPressChan <- char
	}

}

func listen() {
	for {
		key := <-keyPressChan
		fmt.Println("you pressed", string(key))
	}
}

// func main() {
// 	go doSomething("hello world")
// 	fmt.Println("message")
// 	for {
// 	}
// }

// func doSomething(s string) {
// 	until := 0
// 	for {
// 		fmt.Println("s is", s)
// 		until += 1
// 		if until == 6 {
// 			break
// 		}
// 	}
// }
