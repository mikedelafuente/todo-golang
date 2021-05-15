package main

//import "fmt"
import (
	"fmt"
	"log"
	"net/http"

	"weekendproject/todo/services"

	"github.com/gorilla/mux"
)

// create a handler struct
type HttpHandler struct{}

var toDoService = services.NewToDoService()

// implement `ServeHTTP` method on `HttpHandler` struct

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "Welcome to the HomePage!<br /> <br />")

	items := toDoService.ToDo.ListToDoItems()
	if len(items) > 0 {
		fmt.Fprintf(w, "Your tasks for today are:<br/><ul>")

		for _, v := range items {
			fmt.Fprintf(w, "<li>%s</li>", v.Description)
		}
		fmt.Fprintf(w, "</ul>")
	} else {
		fmt.Fprintf(w, "No tasks for today!")
	}

	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/todo", toDoService.HandleListToDoItems).Methods(http.MethodGet)
	myRouter.HandleFunc("/todo", toDoService.HandlePostToDoItem).Methods(http.MethodPost)
	myRouter.HandleFunc("/todo/{id}", toDoService.HandleDeleteToDoItem).Methods(http.MethodDelete)
	myRouter.HandleFunc("/todo/{id}", toDoService.HandleGetToDoItem).Methods(http.MethodGet)
	myRouter.HandleFunc("/todo/{id}", toDoService.HandlePatchToDoItem).Methods(http.MethodPatch)

	log.Fatal(http.ListenAndServe("localhost:8080", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
