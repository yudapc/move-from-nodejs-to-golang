package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSON parse & stringify")
	type FooBar struct {
		Foo string `json:"foo"`
	}

	dataJsonString := `{"foo":"bar"}`

	//parse
	fooBar := new(FooBar)
	err := json.Unmarshal([]byte(dataJsonString), fooBar)
	if err != nil {
		panic(err)
	}
	fmt.Println(fooBar)
	fmt.Println(fooBar.Foo)

	// stringify
	marshalled, err := json.Marshal(fooBar)
	jsonString := string(marshalled)
	fmt.Println(jsonString)
}
