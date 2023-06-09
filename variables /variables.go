package main

import "fmt"

func main() {
	//1
	var str string = "hello world";
	var num int = 12;

	fmt.Println(str);
	fmt.Println(num);

	//2
	var str2 = "hello world2!!!";
	var num2 = 13;

	fmt.Printf("str2: %s , type: %T\n", str2, str2);
	fmt.Printf("num2: %d , type: %T\n", num2, num2);


	//3
	str3 := "hello world3!!!";
	num3 := 14;

	fmt.Println(str3)
	fmt.Println(num3)

	//4
	var (
		a = 1
		b = 2
		c = 3
	)

	fmt.Println(a,b,c)

}