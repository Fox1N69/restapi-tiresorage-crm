package handlers

import (
	"crud-crm/pkg/controllers"
	"crud-crm/pkg/repository"
	"strconv"
)

// Создание типа Handler
type Handler struct {
	//Добавляем поле mainRepo
	repository repository.MainRepository
	controller controllers.Controllers
}

// Создаем новый экземпляр Handler
func NewHandler(mainRepo repository.MainRepository, controller controllers.Controllers) *Handler {
	return &Handler{
		repository: mainRepo,
		controller: controller,
	}
}

func (h *Handler) ConvertUint(str string) (uint, error) {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(num), nil
}
