package services

import (
	"errors"
	"math"

	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/models/requests"
	"github.com/BigbossXD/auto_cashier/models/responses"
	"github.com/BigbossXD/auto_cashier/orm"
	"github.com/BigbossXD/auto_cashier/repositories"
	"github.com/BigbossXD/auto_cashier/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConfigsService struct {
	configsRepo *repositories.ConfigsRepository
}

func NewConfigsService(configsRepo *repositories.ConfigsRepository) *ConfigsService {
	return &ConfigsService{
		configsRepo: configsRepo,
	}
}

func (s *ConfigsService) GetConfigsById(ID uint) (*models.CashierConfigs, error) {
	configs, err := s.configsRepo.GetConfigsById(ID)
	return configs, err
}

func (s *ConfigsService) GetConfigsByMachineID(machineId uint) (*[]models.CashierConfigs, error) {
	configs, err := s.configsRepo.GetConfigsByMachineID(machineId)
	return configs, err
}

func (s *ConfigsService) GetAllConfigs() ([]models.CashierConfigs, error) {
	configs, err := s.configsRepo.GetAllConfigss()
	return configs, err
}

func (s *ConfigsService) GetConfigsByMoneyValueAndMachineID(moneyValue float32, machineId uint) (*models.CashierConfigs, error) {
	configs, err := s.configsRepo.GetConfigsByMoneyValueAndMachineID(moneyValue, machineId)
	return configs, err
}

func (s *ConfigsService) GetConfigsByMachineIDOrderByValueDesc(machineId uint) ([]models.CashierConfigs, error) {
	configs, err := s.configsRepo.GetConfigsByMachineIDOrderByValueDesc(machineId)
	return configs, err
}

func (s *ConfigsService) CreateConfigs(data *requests.CreateConfigRequest) (*models.CashierConfigs, error) {
	_, err := s.configsRepo.GetConfigsByMoneyValueAndMachineID(data.MoneyValue, data.MachineId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			configs := &models.CashierConfigs{}
			configs.MachineId = data.MachineId
			configs.MoneyValue = data.MoneyValue
			configs.MaximumAmount = data.MaximumAmount
			configs.CurrentAmount = 0
			configs, err := s.configsRepo.CreateConfigs(configs)
			if err != nil {
				return nil, err
			}
			return configs, err
		}
		return nil, err
	}
	err = errors.New("duplicate value")
	return nil, err
}

func (s *ConfigsService) UpdateConfigs(data *requests.UpdateConfigRequest) (*models.CashierConfigs, error) {

	configs, err := s.configsRepo.GetConfigsById(data.ConfigId)
	if err != nil {
		return nil, err
	}

	configs.MaximumAmount = data.MaximumAmount

	configs, err = s.configsRepo.UpdateConfigs(configs)
	if err != nil {
		return nil, err
	}
	return configs, err
}

func (s *ConfigsService) DeleteConfigs(ID uint) (*models.CashierConfigs, error) {

	deleteConfigs, err := s.configsRepo.GetConfigsById(ID)
	if err != nil {
		return nil, err
	}

	deleteConfigs, err = s.configsRepo.DeleteConfigs(deleteConfigs)
	return deleteConfigs, err
}

func (s *ConfigsService) Deposit(depositRequest *requests.DepositRequest) ([]models.CashierConfigs, error) {

	var err error
	for _, v := range depositRequest.Items {
		configs, err := s.configsRepo.GetConfigsById(v.ID)
		if err != nil {
			return nil, err
		}

		if (configs.CurrentAmount + v.Amount) > configs.MaximumAmount {
			utils.Logger.Sugar().Error("exceeded the limit!")
			err = errors.New("exceeded the limit!")
			return nil, err
		}
	}

	SessionId := uuid.New().String()

	for _, v := range depositRequest.Items {
		configs, err := s.configsRepo.GetConfigsById(v.ID)
		if err != nil {
			return nil, err
		}

		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount + v.Amount
		configs, err = s.configsRepo.UpdateConfigs(configs)

		// waiting refactor
		if v.Amount > 0 {
			transection := &models.CashierTransections{}
			transection.MachineId = depositRequest.MachineId
			transection.SessionId = SessionId
			transection.Type = models.DEPOSIT
			transection.MoneyValue = configs.MoneyValue
			transection.Amount = v.Amount
			orm.Db.Save(transection)
		}

	}

	configsAll, err := s.configsRepo.GetAllConfigss()
	if err != nil {
		return nil, err
	}

	return configsAll, err

}

func (s *ConfigsService) Withdraw(withdrawRequest *requests.WithdrawRequest) ([]models.CashierConfigs, error) {

	var err error
	for _, v := range withdrawRequest.Items {
		configs, err := s.configsRepo.GetConfigsById(v.ID)
		if err != nil {
			return nil, err
		}

		if configs.CurrentAmount < v.Amount {
			err = errors.New("exceeded the limit!")
			return nil, err
		}
	}

	SessionId := uuid.New().String()

	for _, v := range withdrawRequest.Items {
		configs, err := s.configsRepo.GetConfigsById(v.ID)
		if err != nil {
			return nil, err
		}

		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount - v.Amount
		configs, err = s.configsRepo.UpdateConfigs(configs)

		// waiting refactor
		if v.Amount > 0 {
			transection := &models.CashierTransections{}
			transection.MachineId = withdrawRequest.MachineId
			transection.SessionId = SessionId
			transection.Type = models.WITHDRAW
			transection.MoneyValue = configs.MoneyValue
			transection.Amount = v.Amount
			orm.Db.Save(transection)
		}
	}

	configsAll, err := s.configsRepo.GetAllConfigss()
	if err != nil {
		return nil, err
	}
	return configsAll, err
}

func (s *ConfigsService) Receive(receiveRequest *requests.ReceiveRequest) ([]responses.ChangeItemRequest, error) {

	var err error
	receiveTotal := float32(0.00)

	for _, v := range receiveRequest.Items {

		configs, err := s.configsRepo.GetConfigsById(v.ID)
		if err != nil {
			return nil, err
		}

		if (configs.CurrentAmount + v.Amount) > configs.MaximumAmount {
			utils.Logger.Sugar().Error("exceeded the limit!")
			err = errors.New("exceeded the limit!")
			return nil, err
		}

		receiveTotal = receiveTotal + float32(v.Amount)*configs.MoneyValue

	}

	if receiveTotal < receiveRequest.Price {
		err = errors.New("not enough change!")
		return nil, err
	}

	configsAll, err := s.configsRepo.GetConfigsByMachineIDOrderByValueDesc(receiveRequest.MachineId)
	if err != nil {
		return nil, err
	}

	receivePrice := receiveRequest.Price
	change, flag := calulateChange(receivePrice, receiveTotal, configsAll)
	if !flag {
		err = errors.New("not enough change!")
		return nil, err
	}

	SessionId := uuid.New().String()

	for _, v := range receiveRequest.Items {
		configs, err := s.configsRepo.GetConfigsById(v.ID)
		if err != nil {
			return nil, err
		}

		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount + v.Amount
		orm.Db.Save(configs)

		// waiting refactor
		if v.Amount > 0 {
			transection := &models.CashierTransections{}
			transection.MachineId = receiveRequest.MachineId
			transection.SessionId = SessionId
			transection.Type = models.RECEIVE
			transection.MoneyValue = configs.MoneyValue
			transection.Amount = v.Amount
			orm.Db.Save(transection)
		}
	}

	for _, v := range change {
		configs, err := s.configsRepo.GetConfigsById(v.ID)
		if err != nil {
			return nil, err
		}

		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount - v.Amount
		orm.Db.Save(configs)

		// waiting refactor
		if v.Amount > 0 {
			transection := &models.CashierTransections{}
			transection.MachineId = receiveRequest.MachineId
			transection.SessionId = SessionId
			transection.Type = models.CHANGE
			transection.MoneyValue = configs.MoneyValue
			transection.Amount = v.Amount
			orm.Db.Save(transection)
		}
	}

	return change, err
}

func calulateChange(priceTotal float32, reciveTotal float32, configs []models.CashierConfigs) ([]responses.ChangeItemRequest, bool) {

	moneyChange := reciveTotal - priceTotal
	change := []responses.ChangeItemRequest{}
	checkChange := float32(0)
	for _, v := range configs {

		countAmount := int32(moneyChange / v.MoneyValue)
		if countAmount >= 1 && countAmount <= v.CurrentAmount {
			moneyChange = float32(math.Mod(float64(moneyChange), float64(v.MoneyValue)))
			changeItem := responses.ChangeItemRequest{ID: v.ID, MoneyValue: v.MoneyValue, Amount: countAmount}
			change = append(change, changeItem)
			checkChange = checkChange + (float32(countAmount) * v.MoneyValue)
		} else {
			changeItemIgnore := responses.ChangeItemRequest{ID: v.ID, MoneyValue: v.MoneyValue, Amount: 0}
			change = append(change, changeItemIgnore)
		}
	}

	if moneyChange <= 0 {
		return change, true
	}
	return change, false

}
