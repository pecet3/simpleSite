package view

import (
	"html/template"
	"net/http"
	"simpleSite/model"
)

func ShowIndex(posts []model.Post, w http.ResponseWriter) error {
	tmp := template.Must(template.ParseFiles("./view/static/index.html"))

	err := tmp.Execute(w, posts)
	if err != nil {
		return err
	}
	return nil
}
