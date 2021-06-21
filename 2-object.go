package main

import (
	"fmt"
)

func main() {
	fmt.Println("Example 2 Object JS to Go")

	//const Post = {
	//    ID: 300
	//    Title: "Moving from NodeJS to Go",
	//    Author: "Christopher Ganga",
	//    Difficulty: "Beginner",
	//}

	type Post struct {
		ID         int
		Title      string
		Author     string
		Difficulty string
	}

	post := Post{
		ID:         300,
		Title:      "Moving from NodeJS to Go",
		Author:     "Christopher Ganga",
		Difficulty: "Beginner",
	}

	fmt.Println(post)
	fmt.Println(post.ID)
	fmt.Println(post.Title)
	fmt.Println(post.Author)
	fmt.Println(post.Difficulty)
}
