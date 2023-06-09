package main

import "fmt"


func main(){
	str := "Hello World 😀 😍!";
	fmt.Printf("str: %s , len: %d , char: %c \n", str, len(str), str[17])

	rune1 := []rune("Hello World 😀 😍!");
	fmt.Printf("rune: %v , len: %d , char: %c \n", rune1, len(rune1), rune1[1])


	emoji := 128519;
	fmt.Printf("emoji: %c ", emoji)

}