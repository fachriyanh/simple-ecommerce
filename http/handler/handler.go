package handler

import (
	"go-hypefast/entity"
	"go-hypefast/service"
	"go-hypefast/utils"
	"net/http"
)

type Handler struct {
	service service.Usecase
}

func NewHandler(service service.Usecase) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload entity.Product
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.service.CreateProduct(&payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Create product Success")
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	product, err := h.service.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var payload entity.Order
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.service.CreateOrder(&payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Create order Success")
}

func (h *Handler) ListOrder(w http.ResponseWriter, r *http.Request) {
	order, err := h.service.ListOrder()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, order)
}
