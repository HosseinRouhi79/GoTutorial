package main

import "fmt"

func main() {
	price, message := calculatePrice2(30, "bronze")
	fmt.Printf("the price is :%.2f and the message is %s\n", price, message)
}
func calculatePrice2(days int, tip string) (price float64, message string) {
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
	return
}
