package controllers

import (
	"fmt"
	"math"
	"net/http"

	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/models/requests"
	"github.com/BigbossXD/auto_cashier/models/responses"
	"github.com/BigbossXD/auto_cashier/orm"
	"github.com/BigbossXD/auto_cashier/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

func GetMaximum(c echo.Context) error {

	configs := []models.CashierConfigs{}
	result := orm.Db.Find(&configs)
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
			utils.Logger.Sugar().Error(err.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00402",
				Message: err.Error(),
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

	for _, v := range depositRequest.Items {
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(err.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00502",
				Message: err.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount + v.Amount
		orm.Db.Save(configs)

	}

	configsAll := []models.CashierConfigs{}
	result := orm.Db.Find(&configsAll)
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
			utils.Logger.Sugar().Error(err.Error())
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

	for _, v := range withdrawRequest.Items {
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(err.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00502",
				Message: err.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount - v.Amount
		orm.Db.Save(configs)

	}

	configsAll := []models.CashierConfigs{}
	result := orm.Db.Find(&configsAll)
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
			utils.Logger.Sugar().Error(err.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00402",
				Message: err.Error(),
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
		utils.Logger.Sugar().Error("not paying enough!")
		response := responses.ErrorBaseResponse{
			Code:    "00403",
			Message: "not paying enough!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	configsAll := []models.CashierConfigs{}
	result := orm.Db.Find(&configsAll)
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
			Code:    "00504",
			Message: "not enough change!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	for _, v := range receiveRequest.Items {
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(err.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00505",
				Message: err.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}
		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount + v.Amount
		orm.Db.Save(configs)
	}

	for _, v := range change {
		configs := &models.CashierConfigs{}
		result := orm.Db.Where("ID = ?", v.ID).First(&configs)
		if result.Error != nil {
			utils.Logger.Sugar().Error(err.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00505",
				Message: err.Error(),
			}
			return c.JSON(http.StatusBadRequest, response)
		}
		configs.ID = v.ID
		configs.CurrentAmount = configs.CurrentAmount - v.Amount
		orm.Db.Save(configs)
	}

	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    change,
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

	fmt.Println("moneyChange", " => ", moneyChange)
	fmt.Println("checkChange", " => ", checkChange)
	fmt.Println("Change List", " => ", change)
	if moneyChange <= 0 {
		return change, true
	}
	return change, false

}