package handlers

import "crud-crm/pkg/repository"

type Handler struct {
	clientRepo *repository.ClientRepository
}

func NewHandler(clientRepo *repository.ClientRepository) *Handler {
	return &Handler{clientRepo: clientRepo}
}

