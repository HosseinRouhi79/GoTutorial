package main

import (
	"fmt"
	"reflect"
)

func main() {

	myMap := make(map[string]string)
	fmt.Println(myMap)

	myMap2 := map[string]string{
		"215479962": "Ali",
		"215479963": "Vahid",
		"215479964": "Milad",
		"215479965": "Reza",
	}
	fmt.Println(myMap2)
	fmt.Println(reflect.TypeOf(myMap2))
	myMap2["215479965"] = "Mahtab" //update element
	fmt.Println(myMap2)

	delete(myMap2, "215479965") //delete element
	fmt.Println(myMap2)
}
