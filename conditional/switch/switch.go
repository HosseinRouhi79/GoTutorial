package main

import "fmt"

func main(){
	var num int;
	println("Enter number: ")
	fmt.Scanln(&num)



	//we can give the selected parameter a default value in switch below

/*

	switch score {
	case 10: do smt
	case 15: do another thing
	}

*/




	switch  {
		case num%2 == 0:
			println("Even")
		case num%2 != 0:
			println("Odd")
		default :print("Unknown")
	}
}