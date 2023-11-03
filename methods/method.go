package main

import "fmt"

type Employee struct {
	ID             int
	Name           string
	SalaryPerMonth float64
	Type           int
}

func main() {
	employee1 := Employee{1, "Alireza", 700, 2}
	fmt.Println(employee1.SalaryPerMonth)
	fmt.Println(employee1.calculateYearlySalary())
}
func (employee Employee) calculateYearlySalary() (yearSalary float64) {
	yearSalary = employee.SalaryPerMonth * 12
	return
}
