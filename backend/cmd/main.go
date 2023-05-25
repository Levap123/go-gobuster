package main

import (
	"log"
	"net/http"

	"github.com/Levap123/go-gobuster/internal/handlers"
	"github.com/Levap123/go-gobuster/internal/services"
	"github.com/Levap123/go-gobuster/internal/usecase"
)

func main() {
	routes, err := usecase.ReadTxt("directory-list-lowercase-2.3-medium.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	services := services.New(routes)
	handlers := handlers.New(services)

	log.Println("started")
	http.ListenAndServe("0.0.0.0:8080", handlers.InitRoutes())
}
