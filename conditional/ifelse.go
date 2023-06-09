package main

import "fmt"
func main(){

	var salaray float64;
	var tax float64;

	fmt.Print("Enter your salaray:")
	fmt.Scanln(&salaray)

	if salaray > 0 && salaray <= 7_300_000 {
		tax = 0
	}else if salaray > 7_300_000 && salaray <= 12_000_000 {
		tax = 0.05
	}else if salaray > 12_000_000 && salaray <= 15_000_000 {
		tax = 0.07
	}else if salaray > 15_000_000 && salaray <= 20_000_000 {
		tax = 0.10
	}else {
		tax = 15
	}

	fmt.Printf("Your pure salary is %.2f\n", salaray)
	fmt.Printf("Your tax is %.2f\n", tax)
	fmt.Printf("Your salary is %.2f\n", salaray - tax*salaray)
}