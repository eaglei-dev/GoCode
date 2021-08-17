// Author : b.sushant
/**Write a program which
- reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order,
separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and
create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice,
and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file,
your program should iterate through your slice of structs
and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	size = 20
)

type NameInfo struct {
	fname string
	lname string
}

func setName(s string, nameStruct *NameInfo) {
	names := strings.Split(s, " ")
	lname := names[0]
	fname := names[1]

	if len(lname) > size {
		nameStruct.fname = fname[:size]
	} else {
		nameStruct.fname = fname
	}

	if len(lname) > size {
		nameStruct.lname = lname[:size]
	} else {
		nameStruct.lname = lname
	}
}

func main() {
	fmt.Println("Enter file name: ")
	var fileName string
	fmt.Scan(&fileName)

	// for each line of file create a NameInfo struct
	f, _ := os.Open(fileName)

	scanner := bufio.NewScanner(f)

	sli := make([]NameInfo, 0, 4)
	for scanner.Scan() {
		s := scanner.Text()
		var nameStruct NameInfo
		setName(s, &nameStruct)
		sli = append(sli, nameStruct)
	}

	for i, val := range sli {
		fmt.Printf("%d : %s %s\n", i, val.fname, val.lname)
	}

}
