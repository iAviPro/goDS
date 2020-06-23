package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// InputValues : To input values
func InputValues() []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Data To Be Sorted: ")
	scanner.Scan()
	valuesToSort := strings.Fields(scanner.Text())
	fmt.Printf("Values Entered: %+v \n", valuesToSort)
	return valuesToSort
}
