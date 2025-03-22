package http

import "github.com/gorilla/mux"

func Router(userHandler *UserHandler) *mux.Router {
	router := mux.NewRouter();

	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router .HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")

	return router
}
