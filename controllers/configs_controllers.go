package controllers

import (
	"net/http"
	"strconv"

	"github.com/BigbossXD/auto_cashier/models"
	"github.com/BigbossXD/auto_cashier/models/requests"
	"github.com/BigbossXD/auto_cashier/models/responses"
	"github.com/BigbossXD/auto_cashier/services"
	"github.com/BigbossXD/auto_cashier/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ConfigsController struct {
	configsService *services.ConfigsService
}

func NewConfigsController(configsService *services.ConfigsService) *ConfigsController {
	return &ConfigsController{
		configsService: configsService,
	}
}

func (h *ConfigsController) FindConfig(c echo.Context) error {

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

	configs, err := h.configsService.GetConfigsByMachineID(uint(machineIdBind))
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
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

func (h *ConfigsController) CreateConfig(c echo.Context) error {

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

	configs, err = h.configsService.CreateConfigs(createConfigRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    configs,
	}
	return c.JSON(http.StatusCreated, response)

}

func (h *ConfigsController) UpdateConfig(c echo.Context) error {

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

	configs, err = h.configsService.UpdateConfigs(updateConfigRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    configs,
	}
	return c.JSON(http.StatusCreated, response)

}

func (h *ConfigsController) DeleteConfig(c echo.Context) error {

	Id := c.Param("id")
	configId, err := strconv.Atoi(Id)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	configs, err := h.configsService.DeleteConfigs(uint(configId))
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	response := responses.SuccessBaseResponse{
		Code:    "00000",
		Message: "Success",
		Data:    configs,
	}
	return c.JSON(http.StatusCreated, response)

}
