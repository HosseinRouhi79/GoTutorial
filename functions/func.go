package main

import "fmt"

func main() {

	finalPrice, message := calculatePrice(2, "silver")
	fmt.Println(finalPrice, message)
}

func calculatePrice(days int, tip string) (float64, string) {
	var price float64
	var message string
	switch tip {
	case "gold":
		price = float64(days * 20)
		message = "gold"
	case "silver":
		price = float64(days * 15)
		message = "silver"
	case "bronze":
		price = float64(days * 10)
		message = "bronze"
	}
	return float64(price), message
}
