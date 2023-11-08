package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"simpleSite/model"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("tajny_klucz_jwt")

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
	tokenString := r.Header.Get("Authorization")
	fmt.Println(tokenString)
	if tokenString == "" {
		// Brak tokenu JWT w nagłówku, możesz zwrócić błąd lub inny komunikat
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Nieprawidłowa metoda podpisywania tokena")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		// Token JWT jest nieprawidłowy lub wygasły, możesz zwrócić błąd autoryzacji
		return
	}

	if err != nil {
		log.Fatal("error get all posts: ", err)
	}

	tmp := template.Must(template.ParseFiles("./static/index.html"))

	err = tmp.Execute(w, posts)
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
			fmt.Println()
			return
		}

		fmt.Println(usersDb)

		fmt.Println("p1: ", user.Password, "p2: ", usersDb[0].Password)

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.Name

		tokenString, err := token.SignedString(jwtKey)

		w.Header().Set("Authorization", "Bearer "+tokenString)
		response := map[string]interface{}{"token": tokenString}
		json.NewEncoder(w).Encode(response)
	}
}
