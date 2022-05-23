package mdl_two

import "fmt"

type Employee struct {
	Name   string
	Salary float32
}

func (e *Employee) salaryTaxFree() float32 {
	if e.Salary <= 50000 {
		return e.Salary * 0.83
	}

	return e.Salary * 0.9
}

func Exercise01() {
	fmt.Println("====== EXERCISE 01 ======")

	var empList = map[string]*Employee{
		"emp1": {"Joao", 49000},
	}

	empList["emp2"] = &Employee{
		Name:   "Federico",
		Salary: 200000,
	}

	for _, r := range empList {
		fmt.Printf("%s has a tax free salary of %f\n", r.Name, r.salaryTaxFree())
	}
}
