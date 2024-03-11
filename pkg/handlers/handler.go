package handlers

import "crud-crm/pkg/repository"

//Создание типа Handler
type Handler struct {
	//Добавляем поле mainRepo
	mainRepo repository.MainRepository
}

//Создаем новый экземпляр Handler
func NewHandler(mainRepo repository.MainRepository) *Handler {
	return &Handler{mainRepo: mainRepo}
}
