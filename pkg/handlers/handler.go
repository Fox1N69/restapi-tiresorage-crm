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

func ConverUint(str string) uint {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}

	return uint(num)
}
