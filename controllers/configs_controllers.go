package controllers

import (
	"net/http"
	"strconv"

	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/models/requests"
	"github.com/BigbossXD/auto_cashier/models/responses"
	"github.com/BigbossXD/auto_cashier/orm"
	"github.com/BigbossXD/auto_cashier/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FindConfig(c echo.Context) error {

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
	result := orm.Db.Where("machine_id = ?", machineIdBind).Find(&configs)
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

func CreateConfig(c echo.Context) error {
	var validate = validator.New()

	configs := &models.CashierConfigs{}

	createConfigRequest := &requests.CreateConfigRequest{}
	if err := c.Bind(createConfigRequest); err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	err := validate.Struct(createConfigRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00401",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	result := orm.Db.Where("money_value = ? and machine_id = ?", createConfigRequest.MoneyValue, createConfigRequest.MachineId).First(&configs)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			configs.MachineId = createConfigRequest.MachineId
			configs.MoneyValue = createConfigRequest.MoneyValue
			configs.MaximumAmount = createConfigRequest.MaximumAmount
			configs.CurrentAmount = 0
			orm.Db.Save(configs)
			response := responses.SuccessBaseResponse{
				Code:    "00000",
				Message: "Success!",
				Data:    configs,
			}
			return c.JSON(http.StatusCreated, response)
		} else {
			utils.Logger.Sugar().Error(result.Error.Error())
			response := responses.ErrorBaseResponse{
				Code:    "00402",
				Message: result.Error.Error(),
			}
			return c.JSON(http.StatusConflict, response)
		}
	}
	utils.Logger.Sugar().Error("duplicate value")
	response := responses.ErrorBaseResponse{
		Code:    "00404",
		Message: "duplicate value",
	}
	return c.JSON(http.StatusConflict, response)
}

func UpdateConfig(c echo.Context) error {

	var validate = validator.New()

	configs := &models.CashierConfigs{}

	updateConfigRequest := &requests.UpdateConfigRequest{}
	if err := c.Bind(updateConfigRequest); err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	err := validate.Struct(updateConfigRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00401",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	result := orm.Db.Where("ID = ?", updateConfigRequest.ConfigId).First(&configs)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00402",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusConflict, response)
	}

	configs.ID = updateConfigRequest.ConfigId
	configs.MaximumAmount = updateConfigRequest.MaximumAmount
	orm.Db.Save(configs)
	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success!",
		Data:    configs,
	}
	return c.JSON(http.StatusCreated, response)
}

func DeleteConfig(c echo.Context) error {

	configId := c.Param("id")

	configs := &models.CashierConfigs{}

	result := orm.Db.Where("id = ?", configId).First(&configs)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00402",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	orm.Db.Delete(&configs)

	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
	}
	return c.JSON(http.StatusOK, response)
}
