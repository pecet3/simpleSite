package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"simpleSite/model"
)

func refreshPosts(w http.ResponseWriter) {
	posts, err := model.GetAllPosts()
	if err != nil {
		log.Fatal("error during get all tasks: ", err)
	}

	tmp := template.Must(template.ParseFiles("./static/index.html"))

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

	tmp := template.Must(template.ParseFiles("./static/index.html"))

	err = tmp.ExecuteTemplate(w, "Posts", posts)
	if err != nil {
		log.Fatal("error executing index.html: ", err)
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
		log.Fatal("error:", err)
	}

	refreshPosts(w)
}

func templating(w http.ResponseWriter, fileName string, filePath string, data interface{}) error {
	tmp := template.Must(template.ParseFiles("./static/index.html"))

	err := tmp.ExecuteTemplate(w, fileName, data)
	if err != nil {
		return err
	}
	return nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		tmp := template.Must(template.ParseFiles("./static/register.html"))

		err := tmp.Execute(w, "register")
		if err != nil {
			log.Fatal("error:", err)
		}

	} else if method == "POST" {
		user := &model.User{}

		user.Name = r.FormValue("name")
		user.Password = r.FormValue("password")
		repeatedPassword := r.FormValue("repeatedPassword")

		if repeatedPassword != user.Password {
			//send htmx
			return
		}
		var err error
		user, err = user.RegisterUser()
		if err != nil {
			log.Fatal("error register:", err)
		}
		fmt.Println(user)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		tmp := template.Must(template.ParseFiles("./static/login.html"))

		err := tmp.Execute(w, "login")
		if err != nil {
			log.Fatal("error:", err)
		}

	}
	if method == "POST" {
		user := &model.User{}
		fmt.Println(method)
		user.Name = r.FormValue("name")
		user.Password = r.FormValue("password")

		name := user.Name
		var err error
		usersDb, err := model.GetUserByName(name)
		if err != nil {
			log.Fatal("error register:", err)
		}
		fmt.Println(usersDb)

		if usersDb[0].Password != user.Password {
			fmt.Println("lol")
		}

		fmt.Println(usersDb)

		fmt.Println("p1: ", user.Password, "p2: ", usersDb[0].Password)
	}
}
