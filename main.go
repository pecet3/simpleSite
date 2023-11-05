package main

import (
	"fmt"
	model "simpleSite/model"
	"simpleSite/router"
)

func main() {
	fmt.Println("hello")
	model.ConnectDb()
	router.SetupAndRun()
}
