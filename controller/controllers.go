package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"simpleSite/model"
	"simpleSite/view"
)

func refreshPosts(w http.ResponseWriter) {
	posts, err := model.GetAllPosts()
	if err != nil {
		log.Fatal("error during get all tasks: ", err)
	}

	tmp := template.Must(template.ParseFiles("./view/static/index.html"))

	err = tmp.ExecuteTemplate(w, "Posts", posts)
	if err != nil {
		log.Fatal("error executing index.html: ", err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	posts, err := model.GetAllPosts()
	if err != nil {
		log.Fatal("error get all posts: ", err)
	}

	err = view.ShowIndex(posts, w)
	if err != nil {
		log.Fatal("error show index: ", err)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := &model.Post{}
	post.Content = r.FormValue("content")
	post.UserId = 1

	err := r.ParseForm()
	if err != nil {
		log.Fatal("error during parsing form data: ", err)
	}

	post, err = post.CreatePost()
	if err != nil {
		log.Fatal("error during creating task:", err)
	}

	refreshPosts(w)
}

func templating(w http.ResponseWriter, fileName string, filePath string, data interface{}) error {
	tmp := template.Must(template.ParseFiles("./view/static/index.html"))

	err := tmp.ExecuteTemplate(w, fileName, data)
	if err != nil {
		return err
	}
	return nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	fmt.Println(method)

	if method == "GET" {
		tmp := template.Must(template.ParseFiles("./view/static/register.html"))

		err := tmp.Execute(w, "register")
		if err != nil {
			log.Fatal("error during creating task:", err)
		}

	}
}
