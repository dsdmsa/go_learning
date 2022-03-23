package staff

var OverPayedLimit = 100
var underPayedLimit = 90

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpayed() []Employee {
	var overpayed []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverPayedLimit {
			overpayed = append(overpayed, x)
		}
	}
	return overpayed

}

func (e *Office) Undepayed() []Employee {
	var underpayed []Employee

	for _, x := range e.AllStaff {
		if x.Salary < underPayedLimit {
			underpayed = append(underpayed, x)
		}
	}
	return underpayed
}
