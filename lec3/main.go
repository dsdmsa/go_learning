package main

import (
	"fmt"
	"myapp/pkg1"
)

func main() {

	var myVar = ""
	fmt.Println(myVar)

	neString := pkg1.PublicVar
	fmt.Println(neString)

	pkg1.Exported()

}
