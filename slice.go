package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	size := 3
	count := 0
	fmt.Println("Enter Digits, Enter x to terminate: ")
	sli := make([]int, size)
	var input string
	for {
		fmt.Scan(&input)

		if "x" == input || "X" == input {
			return
		}

		var n int
		var err error

		n, err = strconv.Atoi(input)
		if err != nil {
			println(err.Error())
			continue
		}
		count += 1
		// add to slice
		if count <= size {
			// inserting the elemets at 0th index,
			// as we are sorting the slice later - the elements will be at
			// their correct places.
			sli[0] = n
		} else {
			sli = append(sli, n)
		}

		// sort the slice
		sort.Slice(sli, func(p, q int) bool {
			return sli[p] < sli[q]
		})

		// print the slice
		fmt.Printf("sli: %v\n", sli)
	}
}
