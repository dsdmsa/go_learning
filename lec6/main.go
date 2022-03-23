package main

import (
	"fmt"
	"sort"
)

var myint int
var myInt16 int16
var yInt32 int32
var yInt64 int64

var myuint uint // unsignint

var myfloat float32
var myfloatibig float64

type Car struct {
	NumberOfTyres int
	Luxury        bool
	Make          string
	Year          int
	Modle         string
}

type Animal struct {
	Name  string
	Sount string
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sount)
}

func main() {

	dog := Animal{
		Name:  "lealea",
		Sount: "mumu",
	}

	cat := Animal{
		Name:  "cat",
		Sount: "meau",
	}

	cat.Says()

	dog.Says()

	return
	myT := sumAll(2, 3, 4, 6, 7, 66) // variatic function
	fmt.Println(myT)

	return
	z := addTwo(2, 4) // do not use naked func
	fmt.Println(z)

	return
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["t"] = 2
	intMap["th"] = 3

	for i, x := range intMap {
		fmt.Println(i, x)
	}

	// delete(intMap, "t")

	el, ok := intMap["t"]

	if ok {
		fmt.Println(el, " is in map")
	} else {
		fmt.Println(el, "not in map")
	}

	return

	var animals []string // slices
	animals = append(animals, "dog")
	animals = append(animals, "xog2")
	animals = append(animals, "dog3")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println(animals[1])
	fmt.Println(animals[0:2])

	fmt.Println(len(animals))

	fmt.Println(sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println(sort.StringsAreSorted(animals))

	x := 10

	myFIrstPointer := &x

	fmt.Println("x is ", x)
	fmt.Println("my pointer is", myFIrstPointer)

	*myFIrstPointer = 15
	fmt.Println("x is now", x)

	changeMyVal(&x)
	fmt.Println("x is now now 2:", x)

	var myInt int
	myInt = 10

	fmt.Println(myInt)

	var myCar Car
	myCar.Luxury = true
	myCar.Make = "BMR"

	my2car := Car{
		Luxury: true,
		Year:   2,
		Modle:  "BBR",
	}

	fmt.Printf("my car %s\n", my2car.Modle)

	var myStrings [3]string
	myStrings[0] = "cat0"
	myStrings[1] = "cat1"
	myStrings[2] = "cat2"

}

func changeMyVal(num *int) {
	*num = 100
}

func addTwo(x, y int) (sum int) {
	sum = x + y
	return
}

func sumAll(nums ...int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	return total
}
