package main 

import (
	"fmt"
)

type salary interface {
	monthly_salary() float64
}

type Google_employee struct {
	hours float64
 	days float64
	hourly_pay float64
	weeks int
}

type Netflix_employee struct {
	hours float64
 	days float64
	hourly_pay float64
	weeks int
}

func (g Google_employee) monthly_salary() float64 {
	return ((g.hours * g.hourly_pay) * g.days) * float64(g.weeks)
}

func (n Netflix_employee) monthly_salary() float64 {
	return ((n.hours * n.hourly_pay) * n.days) * float64(n.weeks)
}

func main() {
	employee1 := Google_employee {
		hours: 10,
		days: 4,
		hourly_pay: 4.50,
		weeks: 3,
	}

	employee2 := Netflix_employee {
		hours: 11,
		days: 5,
		hourly_pay: 5,
		weeks: 4,
	}

	fmt.Println("the monthly salary for employee1: ",employee1.monthly_salary())
	fmt.Println("the monthly salary for employee2: ",employee2.monthly_salary())
}