package main

import (
	"fmt"
)

func main() {
	fmt.Println("Object from js to golang")

	//const Post = {
	//    ID: 300
	//    Title: "Moving from NodeJS to Go",
	//    Author: "Christopher Ganga",
	//    Difficulty: "Beginner",
	//}

	post := map[string]interface{}{
		"ID":         300,
		"Title":      "Moving from NodeJS to Go",
		"Author":     "Christopher Ganga",
		"Difficulty": "Beginner",
	}
	fmt.Println(post)
	fmt.Println(post["ID"])
	fmt.Println(post["Title"])
	fmt.Println(post["Author"])
	fmt.Println(post["Difficulty"])
}
