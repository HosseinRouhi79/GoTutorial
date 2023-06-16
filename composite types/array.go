package main

import "fmt"

func main() {

var myArr [6]string= [6]string{"a", "b", "c", "d", "e"}
var myArr2 = [7]string{"a", "b", "c", "d", "e"}
var myArr3 = [...]int{1,5,6,7}
array := [...]int{1,5,6,7}

var search string
fmt.Println("Enter search string")
fmt.Scanln(&search)

for i, item := range myArr {

	if item == search {
		fmt.Println("The search index is: ", i)
	}

}

fmt.Println(myArr)
fmt.Println(myArr2)
fmt.Println(myArr3)
fmt.Println(array)

}