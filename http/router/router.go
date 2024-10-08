package router

import (
	"go-hypefast/http/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	handler handler.Handler
}

func NewRouter(handler *handler.Handler) *Router {
	return &Router{handler: *handler}
}

func (r *Router) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	r.CheckRouters(subRouter)

	log.Print("Starting server on :8080")
	return http.ListenAndServe(":8080", router)
}

func (r *Router) CheckRouters(router *mux.Router) {
	router.HandleFunc("/products", r.handler.GetProducts).Methods("GET")
	router.HandleFunc("/products", r.handler.CreateProduct).Methods("POST")
	router.HandleFunc("/orders", r.handler.ListOrder).Methods("GET")
	router.HandleFunc("/orders", r.handler.CreateOrder).Methods("POST")
}
