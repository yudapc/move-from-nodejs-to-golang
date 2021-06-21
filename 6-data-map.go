package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sample Data Map")

	dataMap := make(map[string]string)
	dataMap["ID"] = "12"
	dataMap["Name"] = "Yuda Cogati"
	dataMap["Age"] = "30"

	fmt.Println(dataMap)
	fmt.Println(dataMap["Name"])

	for key, value := range dataMap {
		fmt.Println(key + ": " + value)
	}
}
