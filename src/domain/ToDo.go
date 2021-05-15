package domain

import (
	"fmt"
	"sort"
	"time"
	"weekendproject/todo/domain/models"
)

type ToDo struct {
	ToDoItems map[string]models.ToDoItem
}

func NewToDo() *ToDo {
	items := make(map[string]models.ToDoItem)
	return &ToDo{ToDoItems: items}
}

func (s ToDo) GetToDoItem(id string) models.ToDoItem {
	item := s.ToDoItems[id]
	return item
}

func (s ToDo) ListToDoItems() []models.ToDoItem {
	tmp := []models.ToDoItem{}
	for _, value := range s.ToDoItems { // Order not specified
		tmp = append(tmp, value)
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].DueDate.Before(tmp[j].DueDate)
	})

	return tmp
}

func (s ToDo) PatchToDoItem(id string, description string, dueDate time.Time) (item *models.ToDoItem, found bool) {
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

func (s ToDo) PostToDoItem(description string, dueDate time.Time) *models.ToDoItem {
	tdi := models.NewToDoItem(fmt.Sprintf("%s %d", description, len(s.ToDoItems)+1))
	if !dueDate.IsZero() {
		tdi.DueDate = dueDate
	}
	s.ToDoItems[tdi.Id] = *tdi
	return tdi
}

func (s ToDo) DeleteToDoItem(id string) bool {
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
