package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const PROMPT = "and press ENTER when ready."

func main() {
	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	var num1 = rand.Intn(8) + 2
	var num2 = rand.Intn(8) + 2
	var num3 = rand.Intn(8) + 2

	reader := bufio.NewReader(os.Stdin)

	runGame(*reader, num1, num2, num3)

}

func runGame(reader bufio.Reader, num1 int, num2 int, num3 int) {

	fmt.Println("guess number game")
	fmt.Println("==================")
	fmt.Println("")

	fmt.Println(" inpiut a num, 1 and 10", PROMPT)
	reader.ReadString('\n')

	fmt.Println(" multiply your number by", num1, PROMPT)
	reader.ReadString('\n')

	fmt.Println(" multiply your number by", num2, PROMPT)
	reader.ReadString('\n')

	fmt.Println(" divide the result by your original number", PROMPT)
	reader.ReadString('\n')

	fmt.Println(" now substract", num3, PROMPT)
	reader.ReadString('\n')

	fmt.Println("the answer is", num1*num2-num3)

}
