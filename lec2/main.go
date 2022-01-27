package main

import (
	"bufio"
	"fmt"
	"os"
)

const PROMPT = "and press ENTER when ready."

func main() {

	var num1 = 2
	var num2 = 5
	var num3 = 7
	var answer int

	reader := bufio.NewReader(os.Stdin)

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

	answer = num1*num2 - num3
	fmt.Println("the answer is", answer)

}
