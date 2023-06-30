package main

import "fmt"

func main() {
	arr := []int{7, 2, 6, 8}
	fmt.Println(twoSum(arr, 9))
}

func twoSum(nums []int, target int) []int {
	arr := []int{}
	myMap := make(map[int]int)

	for index, value := range nums {
		cmp := target - value
		if _, ok := myMap[cmp]; ok {
			arr = append(arr, index)
			arr = append(arr, myMap[cmp])
		} else {
			myMap[value] = index
		}
	}

	return arr
}

// its the optimized code
