package main

import (
	"fmt"
	"log"
	model "simpleSite/model"
)

func main() {
	fmt.Println("hello")
	model.ConnectDb()
	err := model.CreatePost("aaa", 1)
	if err != nil {
		log.Println("error", err)
	}
}
