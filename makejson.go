//Author: b.sushant
/**
Write a program which prompts the user to first enter a name,
and then enter an address.
Your program should create a map and add the name and address
to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
*/
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	fmt.Println("Enter First Name: ")
	fmt.Scan(&name)
	fmt.Println("Enter Address   : ")
	fmt.Scan(&address)

	infoMap := map[string]string{
		"name":    name,
		"address": address}

	jsonInfo, err := json.Marshal(infoMap)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s", jsonInfo)

}
