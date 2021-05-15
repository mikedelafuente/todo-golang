package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/mux"
	"weekendproject.app/todo/models"
)

type ToDoService struct {
	ToDoItems map[string]models.ToDoItem
}

func NewToDoService() *ToDoService {
	items := make(map[string]models.ToDoItem)
	return &ToDoService{ToDoItems: items}
}

func (s ToDoService) DeleteToDoItem(id string) bool {
	f := s.ToDoItems[id]
	deleted := false
	if f.Id == id {
		delete(s.ToDoItems, id)
		deleted = true
	} else {
		deleted = false
	}

	return deleted
}

func (s ToDoService) HandleDeleteToDoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleted := s.DeleteToDoItem(id)

	if deleted {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s ToDoService) PatchToDoItem(id string, description string, dueDate time.Time) (item *models.ToDoItem, found bool) {
	tdi := s.GetToDoItem(id)
	isFound := false

	if len(tdi.Id) > 0 {
		isFound = true

		if !dueDate.IsZero() {
			tdi.DueDate = dueDate
		}

		if len(description) > 0 {
			tdi.Description = description
		}

		s.ToDoItems[tdi.Id] = tdi
	}

	return &tdi, isFound
}

func (s ToDoService) HandlePatchToDoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item models.ToDoItem

	json.NewDecoder(r.Body).Decode(&item)

	tdi, found := s.PatchToDoItem(id, item.Description, item.DueDate)

	if !found {
		handleResponse(w, []byte{}, http.StatusNotFound)
	} else {
		b, _ := marshalFormat(tdi)

		handleResponse(w, b, http.StatusCreated)
	}

}

func (s ToDoService) PostToDoItem(description string, dueDate time.Time) *models.ToDoItem {
	tdi := models.NewToDoItem(fmt.Sprintf("%s %d", description, len(s.ToDoItems)+1))
	if !dueDate.IsZero() {
		tdi.DueDate = dueDate
	}
	s.ToDoItems[tdi.Id] = *tdi
	return tdi
}

func (s ToDoService) HandlePostToDoItem(w http.ResponseWriter, r *http.Request) {
	description := "[No description]"

	var item models.ToDoItem
	item.Description = description

	json.NewDecoder(r.Body).Decode(&item)

	tdi := s.PostToDoItem(item.Description, item.DueDate)

	b, _ := marshalFormat(tdi)

	handleResponse(w, b, http.StatusCreated)
}

func (s ToDoService) ListToDoItems() []models.ToDoItem {
	tmp := []models.ToDoItem{}
	for _, value := range s.ToDoItems { // Order not specified
		tmp = append(tmp, value)
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].DueDate.Before(tmp[j].DueDate)
	})

	return tmp
}

func (s ToDoService) HandleListToDoItems(w http.ResponseWriter, r *http.Request) {
	tmp := s.ListToDoItems()

	b, _ := marshalFormat(tmp)

	if len(s.ToDoItems) > 0 {
		handleResponse(w, b, http.StatusOK)
	} else {
		handleResponse(w, b, http.StatusNotFound)
	}
}

func (s ToDoService) GetToDoItem(id string) models.ToDoItem {
	item := s.ToDoItems[id]
	return item
}

func (s ToDoService) HandleGetToDoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	item := s.GetToDoItem(id)

	if len(item.Id) > 0 {
		b, _ := marshalFormat(item)
		handleResponse(w, b, http.StatusOK)
	} else {
		handleResponse(w, []byte{}, http.StatusNotFound)
	}
}
