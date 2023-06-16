package main

import (
	"fmt"
)

func main(){
	var number int
	var markerArray []int
	var minArray []int
	var indexes []int
	var values []int
	countMap := make(map[int]int)

	fmt.Println("Enter the number of markers")
	fmt.Scanln(&number)
	for i := 1; i <= number; i++ {
		var color int
		fmt.Println("Enter marker color")
		fmt.Scanln(&color)
		markerArray = append(markerArray, color)
	}

	for _, value := range markerArray {
		if  _, ok := countMap[value]; ok {
			countMap[value]++
		}else {
			countMap[value] = 1
		}
	}
	for index, value := range countMap {

		indexes = append(indexes, index)
		values = append(values, value)

	}

	min := values[0]
	for _, value := range indexes {

		if value<=min {
			min = value
		}

	}
	for index, value := range countMap {

		if value == min {
			minArray = append(minArray, index)			
		}

	}
	min1 := minArray[0]
	for _, value := range minArray {

		if value<=min1 {
			min1 = value
		}

	}

	fmt.Println(indexes)
	fmt.Println(values)
	fmt.Println(minArray)
	fmt.Println(min1)
	
}