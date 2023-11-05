package view

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"simpleSite/model"
)

func ShowIndex(posts []model.Post, w http.ResponseWriter) {
	fmt.Println(posts, "aaaaaaaaaaaaaa")
	tmp := template.Must(template.ParseFiles("./view/static/index.html"))

	err := tmp.Execute(w, posts)
	if err != nil {
		log.Fatal("error executing index.html: ", err)
	}

}
