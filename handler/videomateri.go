package handler

import (
	"mediaislam/helper"
	"mediaislam/videomateri"
	"net/http"

	"github.com/gin-gonic/gin"
)

type videomateriHandler struct {
	service videomateri.Service
}

func NewVideomateriHandler(service videomateri.Service) *videomateriHandler {
	return &videomateriHandler{service}
}

func (h *videomateriHandler) CreateVideomateri(c *gin.Context) {
	var input videomateri.CreateVideomateriInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create videomateri", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newVideomateri, err := h.service.CreateVideomateri(input)
	if err != nil {
		response := helper.APIResponse("Failed to create videomateri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("videomateri created", http.StatusCreated, "success", videomateri.FormatVideomateri(newVideomateri))
	c.JSON(http.StatusCreated, response)
}

func (h *videomateriHandler) UpdateVideomateri(c *gin.Context) {
	var input videomateri.GetVideomateriDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("error to get videomateri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData videomateri.CreateVideomateriInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update videomateri", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateVideomateri, err := h.service.UpdateVideomateri(input, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update videomateri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("videomateri updated", http.StatusOK, "success", videomateri.FormatVideomateri(updateVideomateri))
	c.JSON(http.StatusOK, response)
}
		
