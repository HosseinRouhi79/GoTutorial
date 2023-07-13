package main

import (
	"fmt"
	"time"
)

func main() {
	closure([]string{"Ali", "Ehsan", "Shima", "Vahid"})
	closureGo([]string{"Ali", "Ehsan", "Shima", "Vahid", "Vahid2", "Vahid3", "Vahid4", "Vahid5", "Zahra", "Vahid7", "Vahid8"})
}

func closure(input []string) {
	for index, value := range input {
		func() {
			fmt.Printf("index: %d , value: %s\n", index, value)
		}()
	}
}

func closureGo(input []string) {
	for i, name := range input {
		go func(i int, name string) {
			name = "*" + name
			println(i, name)
		}(i, name)
	}
	time.Sleep(time.Second * 1)
}
