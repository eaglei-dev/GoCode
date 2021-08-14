package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	size := 3
	fmt.Println("Enter Digits, Enter x to terminate: ")
	sli := make([]int, 0, size)
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
		// add to slice
		sli = append(sli, n)

		// sort the slice
		sort.Slice(sli, func(p, q int) bool {
			return sli[p] < sli[q]
		})

		// print the slice
		fmt.Printf("sli: %v\n", sli)
	}
}
