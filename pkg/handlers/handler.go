package handlers

import (
	"crud-crm/pkg/repository"
	"strconv"
)

// Создание типа Handler
type Handler struct {
	//Добавляем поле mainRepo
	mainRepo repository.MainRepository
}

// Создаем новый экземпляр Handler
func NewHandler(mainRepo repository.MainRepository) *Handler {
	return &Handler{mainRepo: mainRepo}
}

func (h *Handler) ConvertUint(str string) (uint, error) {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(num), nil
}
