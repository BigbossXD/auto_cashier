package controllers

import (
	"net/http"
	"strconv"

	"github.com/BigbossXD/auto_cashier/models/requests"
	"github.com/BigbossXD/auto_cashier/models/responses"
	"github.com/BigbossXD/auto_cashier/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *ConfigsController) GetMaximum(c echo.Context) error {

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

	configs, err := h.configsService.GetConfigsByMachineIDOrderByValueDesc(uint(machineIdBind))
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

func (h *ConfigsController) Deposit(c echo.Context) error {

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

	configs, err := h.configsService.Deposit(depositRequest)
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

func (h *ConfigsController) Withdraw(c echo.Context) error {

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

	configs, err := h.configsService.Withdraw(withdrawRequest)
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

func (h *ConfigsController) Receive(c echo.Context) error {

	var validate = validator.New()
	receiveRequest := &requests.ReceiveRequest{}
	// receiveTotal := float32(0.00)
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

	configs, err := h.configsService.Receive(receiveRequest)
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
