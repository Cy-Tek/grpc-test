package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-test/protofiles/todos"
	"grpc-test/service"
	"net"
)

type server struct{}

func (s server) CreateTodo(_ context.Context, request *todos.CreateTodoRequest) (*todos.Todo, error) {
	todo, err := service.AddTodo(request.GetTitle(), request.GetListId(), request.GetComplete())
	if err != nil {
		return nil, err
	}

	return &todos.Todo{
		Id:       todo.Id,
		ListId:   todo.ListId,
		Title:    todo.Title,
		Complete: todo.Complete,
	}, nil
}

func (s server) GetTodo(_ context.Context, request *todos.GetTodoRequest) (*todos.Todo, error) {
	todo, err := service.GetTodo(request.GetId())
	if err != nil {
		return nil, err
	}

	return &todos.Todo{
		Id:       todo.Id,
		ListId:   todo.ListId,
		Title:    todo.Title,
		Complete: todo.Complete,
	}, nil
}

func (s server) GetList(_ context.Context, request *todos.GetListRequest) (*todos.TodoList, error) {
	list, err := service.GetList(request.GetId())
	if err != nil {
		return nil, err
	}

	todoResults := make([]*todos.Todo, len(list.Todos))
	for i, todo := range list.Todos {
		todoResults[i] = &todos.Todo{
			Id:       todo.Id,
			ListId:   todo.ListId,
			Title:    todo.Title,
			Complete: todo.Complete,
		}
	}

	return &todos.TodoList{
		Id:    list.Id,
		Title: list.Title,
		Todos: todoResults,
	}, nil
}

func (s server) UpdateTodo(_ context.Context, request *todos.UpdateTodoRequest) (*todos.Todo, error) {
	todo, err := service.UpdateTodo(request.GetId(), request.GetMask().GetPaths(), request.GetTodo())
	if err != nil {
		return nil, err
	}

	return &todos.Todo{
		Id:       todo.Id,
		ListId:   todo.ListId,
		Title:    todo.Title,
		Complete: todo.Complete,
	}, nil
}

func (s server) DeleteTodo(_ context.Context, request *todos.DeleteTodoRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, service.DeleteTodo(request.GetId())
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server started at http://localhost:8080")

	s := grpc.NewServer()
	todos.RegisterTodoServiceServer(s, server{})

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
