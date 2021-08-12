package main

import "fmt"

func main() {
	fmt.Println("Enter floating point number: ")
	// take float input from user
	var floatNumber float32
	fmt.Scanln(&floatNumber)
	// truncate it to integer
	var intPart int
	intPart = int(floatNumber)
	// print that integer
	fmt.Println("Truncated Intrger is ", intPart)
}
