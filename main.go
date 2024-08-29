package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const URL = "https://jsonplaceholder.typicode.com/posts"

type Posts struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetJson(url string, data interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(data)
}

func GetPost() {
	var posts []Post
	err := GetJson(URL, &posts)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for i := 0; i < len(posts); i++ {
			fmt.Println(posts[i].Title)
		}
	}
}

func main() {
	GetPost()
}
