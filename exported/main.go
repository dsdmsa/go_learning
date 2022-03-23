package main

import (
	"log"
	"myapp/staff"
)

var employee = []staff.Employee{
	{FirstName: "Jphn", LastName: "smt", Salary: 10, FullTime: true},
	{FirstName: "Sally", LastName: "asd", Salary: 60, FullTime: true},
	{FirstName: "Mark", LastName: "asdasd", Salary: 100, FullTime: true},
	{FirstName: "Opa", LastName: "sdsds", Salary: 1400, FullTime: false},
	{FirstName: "mar", LastName: "cc", Salary: 19900, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employee,
	}

	staff.OverPayedLimit = 110
	log.Println(myStaff.All())
	log.Println(myStaff.Overpayed())

}
