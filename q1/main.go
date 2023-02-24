package main 

import (
	"fmt"
)

//create an interface named Salary
//in the interface create a signiture method named monthly_salary and returns a float64
type salary interface {
	monthly_salary() float64
}

//struct for the interface
//struct type for the google employee
type Google_employee struct {
	hours float64
 	days float64
	hourly_pay float64
	weeks int
}

//struct for the interface
//struct type for the netflix employee
type Netflix_employee struct {
	hours float64
 	days float64
	hourly_pay float64
	weeks int
}

//use the struct type we created to implement the monthly_salary() of interface
func (g Google_employee) monthly_salary() float64 {
	return ((g.hours * g.hourly_pay) * g.days) * float64(g.weeks)
}

func (n Netflix_employee) monthly_salary() float64 {
	return ((n.hours * n.hourly_pay) * n.days) * float64(n.weeks)
}

//main function
func main() {

	//assign values to the types we created
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

	//print out the result
	fmt.Println("the monthly salary for employee1: ",employee1.monthly_salary())
	fmt.Println("the monthly salary for employee2: ",employee2.monthly_salary())
}