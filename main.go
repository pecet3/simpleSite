package main

import (
	model "simpleSite/model"
	"simpleSite/router"
)

func main() {
	model.ConnectDb()
	router.SetupAndRun()
}
