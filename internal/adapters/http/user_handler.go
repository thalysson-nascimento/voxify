package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/thalysson-nascimento/voxify/internal/application"
	"github.com/thalysson-nascimento/voxify/internal/core"
)

type UserHandler struct {
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{service: service}
}


func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var coreUser core.User

	// Decodificando o JSON do corpo da requisição
	if err := json.NewDecoder(r.Body).Decode(&coreUser); err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	// Chamando o serviço para criar o usuário
	userModel, err := h.service.CreateUser(coreUser)
	if err != nil {
		http.Error(w, "Erro ao criar o usuário", http.StatusInternalServerError)
		return
	}

	// Respondendo com o modelo de usuário criado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(userModel); err != nil {
		log.Println("Erro ao enviar a resposta:", err)
	}
}


func (handler *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.service.GetAllUsers()
	if err != nil {
			http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Erro ao enviar resposta", http.StatusInternalServerError)
		return
	}
}