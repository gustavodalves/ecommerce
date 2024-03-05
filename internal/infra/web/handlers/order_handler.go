package web

import (
	"encoding/json"
	"net/http"

	"github.com/gustavodalves/ecommerce/internal/app/service"
)

type OrderHandler struct {
	Service service.OrderService
}

func NewHandler(
	s service.OrderService,
) *OrderHandler {
	return &OrderHandler{
		Service: s,
	}
}

type Message struct {
	Message string
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	err := h.Service.CreateOrder()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Message{Message: "Order Created"})
}
