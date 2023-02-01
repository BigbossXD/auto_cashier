package controllers

import (
	"net/http"

	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/models/requests"
	"github.com/BigbossXD/auto_cashier/models/responses"
	"github.com/BigbossXD/auto_cashier/orm"
	"github.com/BigbossXD/auto_cashier/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func FindMachine(c echo.Context) error {
	machine := []models.Machine{}
	result := orm.Db.Find(&machine)
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
		Data:    machine,
	}
	return c.JSON(http.StatusOK, response)
}

func FindMachineFull(c echo.Context) error {

	machine := []models.Machine{}
	result := orm.Db.Find(&machine)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00402",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	errorMachineResponse := []responses.ErrorMachineResponse{}
	result = orm.Db.Model(&models.CashierConfigs{}).Select("machine_id").Where("current_amount = maximum_amount").Group("machine_id").Find(&errorMachineResponse)
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
		Data:    errorMachineResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func FindMachineEmpty(c echo.Context) error {

	machine := []models.Machine{}
	result := orm.Db.Find(&machine)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00402",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	errorMachineResponse := []responses.ErrorMachineResponse{}
	result = orm.Db.Model(&models.CashierConfigs{}).Select("machine_id").Where("current_amount = 0").Group("machine_id").Find(&errorMachineResponse)
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
		Data:    errorMachineResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateMachine(c echo.Context) error {

	var validate = validator.New()

	createMachineRequest := &requests.CreateMachineRequest{}

	if err := c.Bind(createMachineRequest); err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err := validate.Struct(createMachineRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00401",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	machine := &models.Machine{}
	machine.Name = createMachineRequest.Name
	orm.Db.Save(machine)

	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    machine,
	}
	return c.JSON(http.StatusOK, response)

}

func UpdateMachine(c echo.Context) error {

	var validate = validator.New()
	machine := &models.Machine{}
	updateMachineRequest := &requests.UpdateMachineRequest{}

	if err := c.Bind(updateMachineRequest); err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err := validate.Struct(updateMachineRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00401",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	result := orm.Db.Where("id = ?", updateMachineRequest.MachineId).First(&machine)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00402",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	machine.ID = updateMachineRequest.MachineId
	machine.Name = updateMachineRequest.Name
	orm.Db.Save(machine)

	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    machine,
	}
	return c.JSON(http.StatusOK, response)

}

func DeleteMachine(c echo.Context) error {
	configId := c.Param("id")

	machine := &models.Machine{}

	result := orm.Db.Where("id = ?", configId).First(&machine)
	if result.Error != nil {
		utils.Logger.Sugar().Error(result.Error.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00402",
			Message: result.Error.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	orm.Db.Delete(&machine)

	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
	}
	return c.JSON(http.StatusOK, response)
}
