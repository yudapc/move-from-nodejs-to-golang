package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Payload struct to json string")

	type User struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Status bool   `json:"status"`
	}

	user := User{
		ID:     123,
		Name:   "Cogati",
		Age:    30,
		Status: true,
	}

	marshalled, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	userSerializer := string(marshalled)
	fmt.Println(userSerializer)
}
