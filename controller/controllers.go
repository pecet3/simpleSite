package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"simpleSite/model"
)

func GetPosts() {
	posts, err := model.GetAllPosts()

	if err != nil {
		log.Println("error get all posts: ", err)
	}

	fmt.Println(posts)
}

func Index(w http.ResponseWriter, r *http.Request) {
	posts, err := model.GetAllPosts()
	if err != nil {
		log.Fatal("error during get all tasks: ", err)
	}
	tmp := template.Must(template.ParseFiles("./view/static/index.html"))

	err = tmp.Execute(w, posts)
	if err != nil {
		log.Fatal("error executing index.html: ", err)
	}
}
