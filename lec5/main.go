package main

import (
	"bufio"
	"fmt"
	"myapp/game"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {

	reader = bufio.NewReader(os.Stdin)
	userName := readString("what is your name")

	age := readInt("How old are you?")

	// fmt.Println("your name is "+userName+". Yand you are", age, "years old.")

	fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old", userName, age))

	return

	playAgain := true

	for playAgain {
		game.Play()
		playAgain = game.GetYesOrNo("Would you like to play again (y/n)?")
	}

	fmt.Println("")
	fmt.Println("Goodbye.")
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("please enter an valid name")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, erro := strconv.Atoi(userInput)
		if erro != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}
