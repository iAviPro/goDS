package main

import (
	"fmt"
	"strconv"
)

func bubbleSortIntValue(str []string) []int {
	arr := make([]int, len(str))
	if _, e := strconv.Atoi(str[0]); e == nil {
		arr = convertStringToIntArray(str)
	} else {
		fmt.Errorf("\nInput Not of Int Type \n %w ", e)
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Size of Output: ", len(arr))
	return arr
}

func bubbleSortStringValues(str []string) []string {
	var arr []string
	arr = str
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Size of Output: ", len(arr))
	return arr
}

func convertStringToIntArray(str []string) []int {
	var arr []int
	for _, val := range str {
		intVal, _ := strconv.Atoi(val)
		arr = append(arr, intVal)
	}
	fmt.Printf("%+v \n", arr)
	return arr
}

// RunBubbleSort : Driver code for bubble sort
func RunBubbleSort() {
	input := InputValues()
	//sorted := bubbleSortIntValue(input)
	sorted := bubbleSortStringValues(input)
	fmt.Printf("Sorted By BubbleSort : %+v \n", sorted)

}
