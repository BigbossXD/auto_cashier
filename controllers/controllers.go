package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/models/requests"
	"github.com/BigbossXD/auto_cashier/models/responses"
	"github.com/BigbossXD/auto_cashier/orm"
	"github.com/BigbossXD/auto_cashier/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetMaximum(c echo.Context) error {

	machineId := c.QueryParam("machineId")
	machineIdBind, err := strconv.Atoi(machineId)
	if err != nil {
		utils.Logger.Sugar().Error("Invalid machineId")
		response := responses.ErrorBaseResponse{
			Code:    "00404",
			Message: "Invalid machineId",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	configs := []models.CashierConfigs{}
	result := orm.Db.Order("money_value DESC").Where("machine_id = ?", machineIdBind).Find(&configs)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00402",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    configs,
	}
	return c.JSON(http.StatusOK, response)
}

func Deposit(c echo.Context) error {

	var validate = validator.New()

	depositRequest := &requests.DepositRequest{}

	if err := c.Bind(depositRequest); err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err := validate.Struct(depositRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00401",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	for _, v := range depositRequest.Items {
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(result.Error.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00402",
				Message: result.Error.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		if (configs.CurrentAmount + v.Amount) > configs.MaximumAmount {
			utils.Logger.Sugar().Error("exceeded the limit!")
			response := responses.ErrorBaseResponse{
				Code:    "00403",
				Message: "exceeded the limit!",
			}
			return c.JSON(http.StatusBadRequest, response)
		}

	}

	SessionId := uuid.New().String()

	for _, v := range depositRequest.Items {
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(result.Error.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00502",
				Message: result.Error.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount + v.Amount
		orm.Db.Save(configs)

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

	configsAll := []models.CashierConfigs{}
	result := orm.Db.Where("machine_id = ?", depositRequest.MachineId).Find(&configsAll)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00503",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    configsAll,
	}
	return c.JSON(http.StatusOK, response)

}

func Withdraw(c echo.Context) error {

	var validate = validator.New()

	withdrawRequest := &requests.WithdrawRequest{}

	if err := c.Bind(withdrawRequest); err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err := validate.Struct(withdrawRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00401",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	for _, v := range withdrawRequest.Items {
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(result.Error.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00402",
				Message: err.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		if configs.CurrentAmount < v.Amount {
			utils.Logger.Sugar().Error("exceeded the limit!")
			response := responses.ErrorBaseResponse{
				Code:    "00403",
				Message: "exceeded the limit!",
			}
			return c.JSON(http.StatusBadRequest, response)
		}

	}

	SessionId := uuid.New().String()

	for _, v := range withdrawRequest.Items {
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(result.Error.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00502",
				Message: result.Error.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount - v.Amount
		orm.Db.Save(configs)

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

	configsAll := []models.CashierConfigs{}
	result := orm.Db.Where("machine_id = ?", withdrawRequest.MachineId).Find(&configsAll)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00503",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    configsAll,
	}
	return c.JSON(http.StatusOK, response)

}

func Receive(c echo.Context) error {

	var validate = validator.New()

	receiveRequest := &requests.ReceiveRequest{}
	receiveTotal := float32(0.00)

	if err := c.Bind(receiveRequest); err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err := validate.Struct(receiveRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00401",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	for _, v := range receiveRequest.Items {
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(result.Error.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00402",
				Message: result.Error.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		if (configs.CurrentAmount + v.Amount) > configs.MaximumAmount {
			utils.Logger.Sugar().Error("exceeded the limit!")
			response := responses.ErrorBaseResponse{
				Code:    "00403",
				Message: "exceeded the limit!",
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		receiveTotal = receiveTotal + float32(v.Amount)*configs.MoneyValue

	}

	if receiveTotal < receiveRequest.Price {
		utils.Logger.Sugar().Error("not enough change!")
		response := responses.ErrorBaseResponse{
			Code:    "00403",
			Message: "not enough change!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	configsAll := []models.CashierConfigs{}
	result := orm.Db.Order("money_value DESC").Where("machine_id = ?", receiveRequest.MachineId).Find(&configsAll)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00503",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	receivePrice := receiveRequest.Price
	change, flag := calulateChange(receivePrice, receiveTotal, configsAll)
	if !flag {
		utils.Logger.Sugar().Error("not enough change!")
		response := responses.ErrorBaseResponse{
			Code:    "00403",
			Message: "not enough change!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	SessionId := uuid.New().String()

	for _, v := range receiveRequest.Items {
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(result.Error.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00505",
				Message: result.Error.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}
		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount + v.Amount
		orm.Db.Save(configs)

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
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(result.Error.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00505",
				Message: result.Error.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}
		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount - v.Amount
		orm.Db.Save(configs)

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

	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    change,
	}
	return c.JSON(http.StatusOK, response)

}

func GetTransection(c echo.Context) error {
	limit := 100
	limitQuery := c.QueryParam("limit")
	limitBind, err := strconv.Atoi(limitQuery)
	if err == nil {
		limit = limitBind
	}

	machineId := c.QueryParam("machineId")
	machineIdBind, err := strconv.Atoi(machineId)
	if err != nil {
		utils.Logger.Sugar().Error("Invalid machineId")
		response := responses.ErrorBaseResponse{
			Code:    "00404",
			Message: "Invalid machineId",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	transections := []models.CashierTransections{}
	result := orm.Db.Where("machine_id = ?", machineIdBind).Order("created_at desc").Limit(limit).Find(&transections)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00402",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    transections,
	}
	return c.JSON(http.StatusOK, response)
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
