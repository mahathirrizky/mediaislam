package handler

import (
	"mediaislam/helper"
	"mediaislam/submateri"
	"net/http"

	"github.com/gin-gonic/gin"
)

type submateriHandler struct {
	service submateri.Service
}

func NewSubmateriHandler(service submateri.Service) *submateriHandler {
	return &submateriHandler{service}
}

func (h *submateriHandler) CreateSubmateri(c *gin.Context) {
	var input submateri.CreateSubmateriInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create submateri", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSUbmateri, err := h.service.CreateSubmateri(input)
	if err != nil {
		response := helper.APIResponse("Failed to create submateri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("submateri created", http.StatusCreated, "success", submateri.FormatCreateSubmateri(newSUbmateri))
	c.JSON(http.StatusCreated, response)
}

func (h *submateriHandler) UpdateSubmateri(c *gin.Context){
	var inputID submateri.GetSubmateriDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("error to get submateri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData submateri.CreateSubmateriInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update submateri", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateSubmateri, err := h.service.UpdateSubmateri(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update submateri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("submateri updated", http.StatusOK, "success", submateri.FormatCreateSubmateri(updateSubmateri))
	c.JSON(http.StatusOK, response)
}
