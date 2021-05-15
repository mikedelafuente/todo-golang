package main

//import "fmt"
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"weekendproject.app/todo/services"
)

// create a handler struct
type HttpHandler struct{}

var td = services.NewToDoService()

// implement `ServeHTTP` method on `HttpHandler` struct

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "Welcome to the HomePage!<br /> <br />")

	items := td.ListToDoItems()
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
	myRouter.HandleFunc("/todo", td.HandleListToDoItems).Methods(http.MethodGet)
	myRouter.HandleFunc("/todo", td.HandlePostToDoItem).Methods(http.MethodPost)
	myRouter.HandleFunc("/todo/{id}", td.HandleDeleteToDoItem).Methods(http.MethodDelete)
	myRouter.HandleFunc("/todo/{id}", td.HandleGetToDoItem).Methods(http.MethodGet)
	myRouter.HandleFunc("/todo/{id}", td.HandlePatchToDoItem).Methods(http.MethodPatch)

	log.Fatal(http.ListenAndServe("localhost:8080", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
