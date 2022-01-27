package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	myMap := make(map[int]string)
	myMap[1] = "map 1"
	myMap[2] = "map 2"
	myMap[3] = "map 3"
	myMap[4] = "map 4"

	fmt.Println("Press anu key, ESC to quit")

	for {

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("pressed %s", myMap[i])
		fmt.Println(t)

	}

	fmt.Println("exit")

}
