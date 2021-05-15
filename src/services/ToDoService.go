package services

import (
	"encoding/json"
	"net/http"

	"weekendproject/todo/domain"
	"weekendproject/todo/domain/models"

	"github.com/gorilla/mux"
)

type ToDoService struct {
	ToDo domain.ToDo
}

func NewToDoService() *ToDoService {
	logic := domain.NewToDo()
	return &ToDoService{ToDo: *logic}
}

func (s ToDoService) HandleDeleteToDoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleted := s.ToDo.DeleteToDoItem(id)

	if deleted {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s ToDoService) HandlePatchToDoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item models.ToDoItem

	json.NewDecoder(r.Body).Decode(&item)

	tdi, found := s.ToDo.PatchToDoItem(id, item.Description, item.DueDate)

	if !found {
		handleResponse(w, []byte{}, http.StatusNotFound)
	} else {
		b, _ := marshalFormat(tdi)

		handleResponse(w, b, http.StatusCreated)
	}

}

func (s ToDoService) HandlePostToDoItem(w http.ResponseWriter, r *http.Request) {
	description := "[No description]"

	var item models.ToDoItem
	item.Description = description

	json.NewDecoder(r.Body).Decode(&item)

	tdi := s.ToDo.PostToDoItem(item.Description, item.DueDate)

	b, _ := marshalFormat(tdi)

	handleResponse(w, b, http.StatusCreated)
}

func (s ToDoService) HandleListToDoItems(w http.ResponseWriter, r *http.Request) {
	tmp := s.ToDo.ListToDoItems()

	b, _ := marshalFormat(tmp)

	if len(s.ToDo.ToDoItems) > 0 {
		handleResponse(w, b, http.StatusOK)
	} else {
		handleResponse(w, b, http.StatusNotFound)
	}
}

func (s ToDoService) HandleGetToDoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	item := s.ToDo.GetToDoItem(id)

	if len(item.Id) > 0 {
		b, _ := marshalFormat(item)
		handleResponse(w, b, http.StatusOK)
	} else {
		handleResponse(w, []byte{}, http.StatusNotFound)
	}
}
