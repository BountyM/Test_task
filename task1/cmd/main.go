package main

import (
	"context"
	"log"
	"task1"
	"task1/pkg/handler"
	"task1/pkg/repository"
	"task1/pkg/service"
)

/*
http://localhost:8080/getResults
Request:
Method get
Body JSON

	{
	    "data": [
	        1,
	        2,
	        ...
			1000
		]
	}

Response:

	{
	    "data": [
	        15,
	        16,
	        ...
	        1014
		]
	}

или, если за 5 секунд не удалось просчитать значения

	{
	     "message": "Для некоторых из отправлнных чисел подсчет незавершен и находится в процессе."
	}
*/
func main() {
	cache := repository.NewCache()

	repository := repository.NewRepository(cache)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(task1.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}
}
