package services

import (
	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/repositories"
)

type TransectionsService struct {
	transectionsRepo *repositories.TransectionsRepository
}

func NewTransectionsService(transectionsRepo *repositories.TransectionsRepository) *TransectionsService {
	return &TransectionsService{
		transectionsRepo: transectionsRepo,
	}
}

func (s *TransectionsService) GetAllTransectionsByMachineIdWithLimit(machineId uint, limit int) ([]models.CashierTransections, error) {
	transections, err := s.transectionsRepo.GetAllTransectionsByMachineIdWithLimit(machineId, limit)
	return transections, err
}
