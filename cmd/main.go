package main

import (
	"fmt"
	nethttp "net/http"

	"github.com/thalysson-nascimento/voxify/internal/adapters/db"
	"github.com/thalysson-nascimento/voxify/internal/adapters/http"
	"github.com/thalysson-nascimento/voxify/internal/application"
)

func main() {
	repo, err := db.NewPostgresRepository()
	if err != nil {
		panic(err)
	}

	service := application.NewUserService(repo)
	handler := http.NewUserHandler(service)
	router := http.Router(handler)

	fmt.Println("Servidor rodando na porta 8080...")
	if err := nethttp.ListenAndServe(":8080", router); err != nil {
		// Trate o erro adequadamente
		fmt.Printf("Erro ao iniciar servidor: %v\n", err)
	}
}
