package repositories

import (
	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/utils"
	"gorm.io/gorm"
)

type ConfigsRepository struct {
	db *gorm.DB
}

func NewConfigsRepo(db *gorm.DB) *ConfigsRepository {
	return &ConfigsRepository{
		db: db,
	}
}

func (r *ConfigsRepository) GetConfigsById(ID uint) (*models.CashierConfigs, error) {
	var configss models.CashierConfigs
	result := r.db.Where("ID = ?", ID).First(&configss)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return &configss, nil
}

func (r *ConfigsRepository) GetConfigsByMachineID(machineId uint) (*[]models.CashierConfigs, error) {
	var configss []models.CashierConfigs
	result := r.db.Where("machine_id = ?", machineId).Find(&configss)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return &configss, nil
}

func (r *ConfigsRepository) GetConfigsByMachineIDOrderByValueDesc(machineId uint) ([]models.CashierConfigs, error) {
	var configs []models.CashierConfigs
	result := r.db.Order("money_value DESC").Where("machine_id = ?", machineId).Find(&configs)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return configs, nil
}

func (r *ConfigsRepository) GetConfigsByMoneyValueAndMachineID(moneyValue float32, machineId uint) (*models.CashierConfigs, error) {
	var configs models.CashierConfigs
	result := r.db.Where("money_value = ? and machine_id = ?", moneyValue, machineId).First(&configs)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return &configs, nil
}

func (r *ConfigsRepository) GetAllConfigss() ([]models.CashierConfigs, error) {
	var configss []models.CashierConfigs
	result := r.db.Find(&configss)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return configss, nil
}

func (r *ConfigsRepository) CreateConfigs(configs *models.CashierConfigs) (*models.CashierConfigs, error) {
	result := r.db.Create(configs)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return configs, nil
}

func (r *ConfigsRepository) UpdateConfigs(configs *models.CashierConfigs) (*models.CashierConfigs, error) {
	result := r.db.Save(configs)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return configs, nil
}

func (r *ConfigsRepository) DeleteConfigs(configs *models.CashierConfigs) (*models.CashierConfigs, error) {
	result := r.db.Delete(&configs)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return configs, nil
}
