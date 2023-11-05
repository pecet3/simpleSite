package router

import (
	"log"
	"net/http"
	"simpleSite/controller"

	"github.com/gorilla/mux"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	todos, err := model.GetAllTasks()
// 	if err != nil {
// 		log.Fatal("error during get all tasks: ", err)
// 	}

// 	tmp := template.Must(template.ParseFiles("./pkg/templates/index.html"))

// 	err = tmp.Execute(w, todos)
// 	if err != nil {
// 		log.Fatal("error executing index.html: ", err)
// 	}

// }
// func sendPost(w http.ResponseWriter) {
// 	todos, err := modelCreatePost()
// 	if err != nil {
// 		log.Fatal("error during get all tasks: ", err)
// 	}

// 	tmp := template.Must(template.ParseFiles("./pkg/templates/index.html"))

// 	err = tmp.ExecuteTemplate(w, "Todos", todos)
// 	if err != nil {
// 		log.Fatal("error executing index.html: ", err)
// 	}

// }

func SetupAndRun() {
	mux := mux.NewRouter()
	port := ":5000"

	mux.HandleFunc("/", controller.Index)

	log.Fatal(http.ListenAndServe(port, mux))
}
