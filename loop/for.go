package main

import(
	"strings"
)

func main(){
	slice := []string{"foo", "bar", "zaz", "var"}

	//range for
	for _,value := range slice{      //the first element is index here (_) and second is value here(value)

		println(strings.ToUpper(value))
		println(strings.ToLower(value))
	}

	//normal loop
	for i := 10; i > 0; i-- {
		println(i)
	}
	
}