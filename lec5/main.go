package main

import (
	"bufio"
	"fmt"
	"log"
	"myapp/game"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwndADog        bool
}

func main() {

	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("what is your name")
	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("what is your favorite number?")
	user.OwndADog = checkDog("do you have a dog?")

	fmt.Printf("Your name is %s. You are %d years old. Your favorite number is %.2f., wunds a dog %t\n",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
		user.OwndADog)

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

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, erro := strconv.ParseFloat(userInput, 64)
		if erro != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func checkDog(s string) bool {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		prompt()

		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if char == 'y' || char == 'Y' {
			return true
		} else {
			return false
		}

	}
}
