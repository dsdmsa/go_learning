package main

import (
	"bufio"
	"fmt"
	"os"
)

const PROMPT = "and press ENTER when ready."

func main() {

	var num1 = 1
	var num2 = 5
	var num3 = 4
	var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("guess number game")
	fmt.Println("==================")
	fmt.Println("")

	fmt.Println(" inpiut a num, 1-10", PROMPT)
	reader.ReadString('\n')

	fmt.Println(" multiply your number by", num1, PROMPT)

}
