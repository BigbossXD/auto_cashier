package repositories

import (
	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/models/responses"
	"github.com/BigbossXD/auto_cashier/utils"
	"gorm.io/gorm"
)

type MachineRepository struct {
	db *gorm.DB
}

func NewMachineRepo(db *gorm.DB) *MachineRepository {
	return &MachineRepository{
		db: db,
	}
}

func (r *MachineRepository) GetMachineById(ID uint) (*models.Machine, error) {
	var machines models.Machine
	result := r.db.Where("ID = ?", ID).First(&machines)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return &machines, nil
}

func (r *MachineRepository) GetAllMachines() ([]models.Machine, error) {
	var machines []models.Machine
	result := r.db.Find(&machines)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return machines, nil
}

func (r *MachineRepository) GetAllMachinesFullStatus() ([]responses.ErrorMachineResponse, error) {
	errorMachineResponse := []responses.ErrorMachineResponse{}
	result := r.db.Model(&models.CashierConfigs{}).Select("machine_id").Where("current_amount = maximum_amount").Group("machine_id").Find(&errorMachineResponse)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return errorMachineResponse, nil
}

func (r *MachineRepository) GetAllMachinesEmptyStatus() ([]responses.ErrorMachineResponse, error) {
	errorMachineResponse := []responses.ErrorMachineResponse{}
	result := r.db.Model(&models.CashierConfigs{}).Select("machine_id").Where("current_amount = 0").Group("machine_id").Find(&errorMachineResponse)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return errorMachineResponse, nil
}

func (r *MachineRepository) CreateMachine(machine *models.Machine) (*models.Machine, error) {
	result := r.db.Create(machine)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return machine, nil
}

func (r *MachineRepository) UpdateMachine(machine *models.Machine) (*models.Machine, error) {
	result := r.db.Save(machine)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return machine, nil
}

func (r *MachineRepository) DeleteMachine(machine *models.Machine) (*models.Machine, error) {
	result := r.db.Save(machine)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		return nil, result.Error
	}
	return machine, nil
}
