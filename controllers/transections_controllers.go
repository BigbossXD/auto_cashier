package controllers

import (
	"net/http"
	"strconv"

	"github.com/BigbossXD/auto_cashier/models/responses"
	"github.com/BigbossXD/auto_cashier/services"
	"github.com/BigbossXD/auto_cashier/utils"
	"github.com/labstack/echo/v4"
)

type TransectionsController struct {
	transectionsService *services.TransectionsService
}

func NewTransectionsController(transectionsService *services.TransectionsService) *TransectionsController {
	return &TransectionsController{
		transectionsService: transectionsService,
	}
}

func (h *TransectionsController) GetTransection(c echo.Context) error {
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

	transections, err := h.transectionsService.GetAllTransectionsByMachineIdWithLimit(uint(machineIdBind), limit)
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
		Data:    transections,
	}
	return c.JSON(http.StatusOK, response)

}
