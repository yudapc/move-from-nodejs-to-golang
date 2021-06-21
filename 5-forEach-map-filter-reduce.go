package main

import (
	"fmt"
	"strings"
)

func main() {
	array := []string{"a", "b", "c"}

	// sample .forEach
	for i, value := range array {
		fmt.Println(i, value)
	}

	// sample Map
	mapped := make([]string, len(array))
	for i, value := range array {
		mapped[i] = strings.ToUpper(value)
	}

	fmt.Println(mapped)

	// sample Filter
	var filtered []string
	for i, value := range array {
		if i%2 == 0 {
			filtered = append(filtered, value)
		}
	}

	fmt.Println(filtered)

	// sample Reduce
	var reduced []string
	for i, value := range array {
		if i%2 == 0 {
			reduced = append(reduced, strings.ToUpper(value))
		}
	}

	fmt.Println(reduced)
}
