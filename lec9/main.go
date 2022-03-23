package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOpPasenegers int
}

type Car struct {
	make     string
	model    string
	year     int
	electric bool
	hybrit   bool
	Vehicle  Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println(" NUmber of pasenger", v.NumberOpPasenegers)
	fmt.Println(" NUmber of pasenweelsger", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.make)
	fmt.Println("Year:", c.year)
	c.Vehicle.showDetails()
}

func main() {

	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOpPasenegers: 6,
	}

	volvo := Car{
		make:    "volvo",
		year:    9919,
		Vehicle: suv,
	}

	volvo.show()

}
