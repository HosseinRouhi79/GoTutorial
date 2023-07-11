package main

import "fmt"

func main() {
	final := randCode([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}, []string{"*", "#", "@", "$"}, 5)
	fmt.Printf("%v", final)
}
func randCode(refArray []string, desArray []string, step int) (finalArray []string) {
	lenRef := len(refArray)
	lenDes := len(desArray)
	var k = 0
	finalArray = []string{}
	for i := 1; i < lenRef; i++ {
		if i%step != 0 || i < step {
			finalArray = append(finalArray, refArray[i])
		} else if i%step == 0 {
			finalArray = append(finalArray, refArray[i])
			for k < lenDes {
				finalArray = append(finalArray, desArray[k])
				k++
				break
			}
		}
	}
	return
}
