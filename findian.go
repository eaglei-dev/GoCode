package main

import (
	"fmt"
	"strings"
)

func main() {
	// check if string starts with i/I ends with n/N and contains a/A
	fmt.Println("Enter a string: ")
	var inStr string
	fmt.Scan(&inStr)
	s := strings.ToLower(inStr)
	if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
