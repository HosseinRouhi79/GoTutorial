package main

func main() {
	defer println("This is defer") // It runs the line after any jobs in func, like finally in php
	println("Hello World")
	println("GoodBye")
}
