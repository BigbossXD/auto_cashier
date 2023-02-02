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

type MachineController struct {
	machineService *services.MachineService
}

func NewMachineController(machineService *services.MachineService) *MachineController {
	return &MachineController{
		machineService: machineService,
	}
}

func (h *MachineController) FindMachineList(c echo.Context) error {
	var machineList []models.Machine
	machineList, err := h.machineService.GetAllMachines()
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
		Data:    machineList,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *MachineController) FindMachineFull(c echo.Context) error {
	var errorMachineResponse []responses.ErrorMachineResponse
	errorMachineResponse, err := h.machineService.GetAllMachinesFullStatus()
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
		Data:    errorMachineResponse,
	}
	return c.JSON(http.StatusOK, response)

}

func (h *MachineController) FindMachineEmpty(c echo.Context) error {

	var errorMachineResponse []responses.ErrorMachineResponse
	errorMachineResponse, err := h.machineService.GetAllMachinesEmptyStatus()
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
		Data:    errorMachineResponse,
	}
	return c.JSON(http.StatusOK, response)

}

func (h *MachineController) CreateMachine(c echo.Context) error {

	var validate = validator.New()
	createMachineRequest := &requests.CreateMachineRequest{}
	if err := c.Bind(createMachineRequest); err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00440",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	err := validate.Struct(createMachineRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00441",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	machine := &models.Machine{}
	machine.Name = createMachineRequest.Name
	machine, err = h.machineService.CreateMachine(machine)
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
		Data:    machine,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *MachineController) UpdateMachine(c echo.Context) error {

	var validate = validator.New()
	machine := &models.Machine{}
	updateMachineRequest := &requests.UpdateMachineRequest{}

	if err := c.Bind(updateMachineRequest); err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00440",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err := validate.Struct(updateMachineRequest)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00441",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	machine.ID = updateMachineRequest.MachineId
	machine.Name = updateMachineRequest.Name
	machine, err = h.machineService.UpdateMachine(machine)
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
		Data:    machine,
	}
	return c.JSON(http.StatusCreated, response)

}

func (h *MachineController) DeleteMachine(c echo.Context) error {
	Id := c.Param("id")
	machineId, err := strconv.Atoi(Id)
	if err != nil {
		utils.Logger.Sugar().Error(err.Error())
		response := responses.ErrorBaseResponse{
			Code:    "00400",
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	machine, err := h.machineService.DeleteMachine(uint(machineId))
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
		Data:    machine,
	}
	return c.JSON(http.StatusCreated, response)
}
