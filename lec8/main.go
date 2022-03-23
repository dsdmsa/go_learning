package main

import "fmt"

type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name    string
	Sound   string
	NumLegs int
}

type Cat struct {
	Name    string
	Sound   string
	NumLegs int
	HasTail bool
}

func (d *Cat) Says() string {
	return d.Sound
}

func (d *Cat) HowManyLegs() int {
	return d.NumLegs
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumLegs
}

func main() {
	dog := Dog{
		Name:    "dog",
		Sound:   "wof",
		NumLegs: 4,
	}

	cat := Cat{
		Name:    "cat",
		Sound:   "mew",
		NumLegs: 4,
	}

	Ridle(&dog)
	Ridle(&cat)

}

func Ridle(a Animal) {
	riddle := fmt.Sprintf("this animal says %s and has %d legs", a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
