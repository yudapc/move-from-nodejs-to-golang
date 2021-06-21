package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Payload json to struct")

	type User struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Status bool   `json:"status"`
	}

	payload := `{"id": 123, "name": "Cogati", "age": 30, "status": true}`
	user := new(User)

	err := json.Unmarshal([]byte(payload), user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
	fmt.Println(user.ID)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Status)
}
