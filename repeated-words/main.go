package main

import (
	"fmt"
	"strings"
)

func main() {
	myMap, fixed := calculateRepeatedWords()
	fmt.Println(myMap)
	fmt.Println(fixed)
}
func calculateRepeatedWords() (myMap map[string]int, fixed string) {

	var sentence string
	fmt.Scanln(&sentence)

	myMap = make(map[string]int)
	var arrfixed = []string{}
	arrSentence := strings.Split(sentence, "")
	for _, value := range arrSentence {
		if _, ok := myMap[value]; ok {
			myMap[value]++
		} else {
			myMap[value] = 1
			arrfixed = append(arrfixed, value)
		}
	}
	fixed = strings.Join(arrfixed, "")

	return
}
