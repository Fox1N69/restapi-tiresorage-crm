package handlers

import "crud-crm/pkg/repository"

type Handler struct {
	mainRepo repository.MainRepository
}

func NewHandler(mainRepo repository.MainRepository) *Handler {
	return &Handler{mainRepo: mainRepo}
}
