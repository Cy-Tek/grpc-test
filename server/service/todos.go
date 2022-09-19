package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var (
	service todoService
)

func init() {
	lists := map[string]TodoList{
		"list-1": {
			Id:    "list-1",
			Title: "General Todos",
			Todos: []*Todo{},
		},
		"list-2": {
			Id:    "list-2",
			Title: "Work Todos",
			Todos: []*Todo{},
		},
		"list-3": {
			Id:    "list-3",
			Title: "Personal Todos",
			Todos: []*Todo{},
		},
	}

	service = todoService{
		lists: lists,
		todos: map[string]*Todo{},
	}
}

type Todo struct {
	Id       string
	ListId   string
	Title    string
	Complete bool
}

type TodoOptions struct {
	ListId   *string
	Title    *string
	Complete *bool
}

type TodoList struct {
	Id    string
	Title string
	Todos []*Todo
}

type todoService struct {
	todos map[string]*Todo
	lists map[string]TodoList
}

func AddTodo(title, listId string, complete bool) (*Todo, error) {
	id := uuid.New()
	todo := Todo{
		Id:       id.String(),
		ListId:   listId,
		Title:    title,
		Complete: complete,
	}

	if list, ok := service.lists[listId]; ok {
		list.Todos = append(list.Todos, &todo)
		service.lists[listId] = list
	} else {
		return nil, errors.New(fmt.Sprintf("could not find a list for list id: %s", listId))
	}

	service.todos[todo.Id] = &todo

	return &todo, nil
}

func GetList(id string) (*TodoList, error) {
	if list, ok := service.lists[id]; ok {
		return &list, nil
	}

	return nil, errors.New(fmt.Sprintf("could not find a list for id: %s", id))
}

func GetTodo(id string) (*Todo, error) {
	if todo, ok := service.todos[id]; ok {
		return todo, nil
	}

	return nil, errors.New(fmt.Sprintf("could not find a todo for id: %s", id))
}

func UpdateTodo(id string, options TodoOptions) (*Todo, error) {
	todo, err := GetTodo(id)
	if err != nil {
		return nil, err
	}

	if options.ListId != nil {
		removeTodoFromList(todo)
		todo.ListId = *options.ListId

		if list, ok := service.lists[*options.ListId]; ok {
			list.Todos = append(list.Todos, todo)
			service.lists[*options.ListId] = list
		} else {
			return nil, errors.New(fmt.Sprintf("could not find a list for list id: %s", *options.ListId))
		}
	}

	if options.Complete != nil {
		todo.Complete = *options.Complete
	}

	if options.Title != nil {
		todo.Title = *options.Title
	}

	return todo, nil
}

func DeleteTodo(id string) error {
	if todo, ok := service.todos[id]; ok {
		removeTodoFromList(todo)
		delete(service.todos, id)
		return nil
	}

	return errors.New(fmt.Sprintf("could not find a todo for id: %s", id))
}

func removeTodoFromList(todo *Todo) {
	if list, ok := service.lists[todo.ListId]; ok {
		for i, item := range list.Todos {
			if item.Id == todo.Id {
				lastIndex := len(list.Todos) - 1
				if i != lastIndex {
					copy(list.Todos[i:], list.Todos[i+1:])
				}

				list.Todos[lastIndex] = nil
				list.Todos = list.Todos[:lastIndex]
				service.lists[todo.ListId] = list
				break
			}
		}
	}
}
