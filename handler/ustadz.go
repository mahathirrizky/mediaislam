package handler

import (
	"mediaislam/helper"
	"mediaislam/user"
	"mediaislam/ustadz"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ustadzHandler struct {
	ustadzService ustadz.Service
}

func NewUstadzHandler(ustadzService ustadz.Service) *ustadzHandler {
	return &ustadzHandler{ustadzService}
}

func (h *ustadzHandler) GetUstadzList(c *gin.Context) {
	ustadzID, _ := strconv.Atoi(c.Query("ustadz_id"))

	dataustadz, err := h.ustadzService.GetUstadz(ustadzID)
	if err != nil {
		response := helper.APIResponse("error to get ustadz", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("list of ustadz", http.StatusOK, "success", ustadz.FormatUstadzList(dataustadz))
	c.JSON(http.StatusOK, response)
}

func (h *ustadzHandler) GetUstadz(c *gin.Context) {
	var input ustadz.GetUstadzDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("error to get ustadz", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ustadzDetail, err := h.ustadzService.GetUstadzByID(input)
	if err != nil {
		response := helper.APIResponse("error to get ustadz", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("ustadz detail", http.StatusOK, "success", ustadz.FormatUstadz(ustadzDetail))
	c.JSON(http.StatusOK, response)
}

func (h *ustadzHandler) RegisterUstadz(c *gin.Context) {
	var input ustadz.RegisterUstadzInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create ustadz", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.UserTable)
	if currentUser.Role == "user" {
		response := helper.APIResponse("Failed to create ustadz, must contributor who can register ustadz", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUstadz, err := h.ustadzService.RegisterUstadz(input)
	if err != nil {
		response := helper.APIResponse("Failed to create ustadz", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create ustadz", http.StatusOK, "success", ustadz.FormatUstadz(newUstadz))
	c.JSON(http.StatusOK, response)
}

func (h *ustadzHandler) UpdateUstadz(c *gin.Context) {
	var inputID ustadz.GetUstadzDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update ustadz", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData ustadz.RegisterUstadzInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update ustadz", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedustadz, err := h.ustadzService.UpdateUstadz(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update ustadz", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to update ustadz", http.StatusOK, "success", ustadz.FormatUstadz(updatedustadz))
	c.JSON(http.StatusOK, response)
}
