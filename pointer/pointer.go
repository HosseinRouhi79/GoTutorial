package main



func main (){
	i := 10
	// j := 20

	var z *int = &i
	*z = *z + 2
	println(*z)
	println(i)
}

// func originalNum (num int){

// 	num = num + 10
// 	println(num)
// }