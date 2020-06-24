package main

// Insertion Sort for Integer / String values

import (
	"fmt"
	"strconv"
)

// insertion sort for integer type
func insertionSortIntValues(values []string) []string {
	str := values
	fmt.Printf("Length of the entered Slice : %d \n", len(str))
	// Start from Exterior loop with 1 till the length of array
	for i := 1; i < len(str); i++ {
		// Interior loop : Till j = i till j>0; keep on comparing a > b and swap
		for j := i; j > 0; j-- {
			a, e := strconv.Atoi(str[j-1])
			if e != nil {
				fmt.Printf("Error Converting %s \n Error: %w", str[j-1], e)
			}
			b, er := strconv.Atoi(str[j])
			if er != nil {
				fmt.Printf("Error Converting %s \n Error: %w", str[j], er)
			}

			// Sorts in ascending order
			if a > b {
				// convert back int to string
				str[j-1], str[j] = strconv.Itoa(b), strconv.Itoa(a)
			}
		}
	}
	return str
}

// insertion sort for String type
func insertionSortStringValues(values []string) []string {
	str := values
	fmt.Printf("Length of the entered Slice : %d \n", len(str))
	for i := 1; i < len(str); i++ {
		for j := i; j > 0; j-- {
			if str[j-1] > str[j] {
				// swap
				str[j-1], str[j] = str[j], str[j-1]
			}
		}
	}
	return str
}

// RunInsertionSort : driver code
func RunInsertionSort() {
	values := InputValues()
	//str := insertionSortIntValues(values)
	str := insertionSortStringValues(values)
	fmt.Printf("Sorted By InsertionSort : %+v \n", str)
}
