package service

import (
	"crud-crm/pkg/models"
	"crud-crm/pkg/repository"
)

type CrequestService struct {
	creqRepository *repository.CrequestRepository
}

func NewCrequestSerivce(creqRepo *repository.CrequestRepository) *CrequestService {
	return &CrequestService{creqRepository: creqRepo}
}

func (s *CrequestService) UpdatePost(id uint, creq *models.Crequests) error {
	extCreq, err := s.creqRepository.GetCrequestByID(id)
	if err != nil {
		return err
	}

	//update message fields based on recaving data
	extCreq.FIO = creq.FIO
	extCreq.Data = creq.Data
	extCreq.IsPaid = creq.IsPaid
	extCreq.Price = creq.Price
	extCreq.Status = creq.Status

	err = s.creqRepository.UpdateCrequest(*extCreq)
	if err != nil {
		return err
	}

	return nil
}
