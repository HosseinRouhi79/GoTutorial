package main

import "fmt"

func main() {

	type Car struct {
		Color  string
		MadeIn string
		Model  string
		Name   string
	}

	carMap := map[string]Car{
		"W4E532": {Color: "Blue", Model: "2017", MadeIn: "Germany", Name: "Audi"},
		"Z4EO92": {Color: "White", Model: "2020", MadeIn: "Germany", Name: "BMW"},
		"K2EP12": {Color: "Black", Model: "2014", MadeIn: "USA", Name: "Ford"},
	}

	fmt.Println(carMap)

	newCarMap := carMap // if we make copy like this, the change in copy affects on the reference map
	fmt.Println(newCarMap)
	delete(newCarMap, "Z4EO92")
	fmt.Println(newCarMap)

	for _, car := range carMap {
		fmt.Println(car.Color) // this doesnt show cars by queue, we should use slices
	}
}

//map is not thread safe
//map is reference type like slices and pointers not value type
