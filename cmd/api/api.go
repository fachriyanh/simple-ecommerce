package api

import (
	"go-hypefast/http/handler"
	"go-hypefast/http/router"
	"go-hypefast/repository"
	"go-hypefast/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIServer struct {
	address string
	db      *gorm.DB
}

func NewAPIServer(address string, db *gorm.DB) *APIServer {
	return &APIServer{address: address, db: db}
}

func (r *APIServer) Run() error {
	server := mux.NewRouter()
	subRouter := server.PathPrefix("/api/v1").Subrouter()

	repository := repository.NewRepository(r.db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	router := router.NewRouter(handler)
	router.CheckRouters(subRouter)

	log.Print("Starting server on :8080")
	return http.ListenAndServe(":8080", server)
}
