package main

import (
	"fmt"
	"sort"
)

func main() {
	sort2(1, 3, 4, 12, 43, 2, 11, 38, 24, 100, 99, 96, 97)
}
func sort2(numbers ...int) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	fmt.Println(numbers)
}
