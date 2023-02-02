package repositories

import (
	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/utils"
	"gorm.io/gorm"
)

type TransectionsRepository struct {
	db *gorm.DB
}

func NewTransectionsRepo(db *gorm.DB) *TransectionsRepository {
	return &TransectionsRepository{
		db: db,
	}
}

func (r *TransectionsRepository) GetAllTransectionsByMachineIdWithLimit(machineId uint, limit int) ([]models.CashierTransections, error) {
	var transections []models.CashierTransections
	result := r.db.Where("machine_id = ?", machineId).Order("created_at desc").Limit(limit).Find(&transections)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return transections, nil
}
