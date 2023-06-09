package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	str := "this is beginning beginning just beginning"

	//split
	strArr := strings.Split(str, " ")

	fmt.Printf("type: %T,    value: %s \n", strArr,strArr)
	fmt.Println(strArr[3])

	newArr := []string{"Hi", "there"} //slice

	fmt.Printf("type: %T, value: %s \n", newArr, newArr)


	//join
	arrStr := strings.Join(newArr, " ")
	fmt.Println(arrStr)
	fmt.Println(newArr[0])

	//replace
	repStr := strings.Replace(str,"beginning","start",1)
	fmt.Println(repStr)

	//replace all
	repStrAll := strings.ReplaceAll(str,"beginning","start")
	fmt.Println(repStrAll)

	//compare & equalFold
	x := "HI darling"
	y := "Hi darling"

	fmt.Println(strings.EqualFold(x,y))
	fmt.Println(strings.Compare(x,y))
	
	

	//contains
	fmt.Println(strings.Contains(x,"HqI"))
	
	//strconv
	strConv := strconv.Itoa(42);

	fmt.Printf("type: %T\n", strConv)


}