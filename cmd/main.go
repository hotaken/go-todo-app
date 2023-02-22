package main

import (
	"github.com/hotaken/go-todo-app"
	"github.com/hotaken/go-todo-app/pkg/handler"
	"github.com/hotaken/go-todo-app/pkg/repository"
	"github.com/hotaken/go-todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
