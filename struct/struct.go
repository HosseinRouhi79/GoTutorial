package main

import "fmt"

func main() {
	type Car struct {
		Model   int
		Company string
	}
	car1 := Car{
		1965, "Benz",
	}
	var car2 = new(Car)
	car3 := car2
	car2.Model = 2002
	car2.Company = "BMW"
	fmt.Println(car1)
	fmt.Println(car3)

}
