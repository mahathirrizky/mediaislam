package handler

import (
	"mediaislam/helper"
	"mediaislam/materi"
	"mediaislam/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type materiHandler struct {
	service materi.Service
}

func NewMateriHandler(service materi.Service) *materiHandler {
	return &materiHandler{service}
}

func (h *materiHandler) GetMateriList(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	datamateri, err := h.service.GetMateri(userID)
	if err != nil {
		response := helper.APIResponse("error to get materi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("list of materi", http.StatusOK, "success", materi.FormatMateriList(datamateri))
	c.JSON(http.StatusOK, response)
}

func (h *materiHandler) GetMateri(c *gin.Context) {
	var input materi.GetMateriDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("error to get materi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	materiDetail, err := h.service.GetMateriByID(input)
	if err != nil {
		response := helper.APIResponse("error to get materi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("materi detail", http.StatusOK, "success", materi.FormatMateri(materiDetail))
	c.JSON(http.StatusOK, response)
}

func (h *materiHandler) CreateMateri(c *gin.Context) {
	var input materi.CreateMateriInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create materi", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.UserTable)
	input.User = currentUser

	newMateri, err := h.service.CreateMateri(input)
	if err != nil {
		response := helper.APIResponse("Failed to create materi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create materi", http.StatusOK, "success", materi.FormatMateri(newMateri))
	c.JSON(http.StatusOK, response)
}

func (h *materiHandler) UpdateMateri(c *gin.Context) {
	var inputID materi.GetMateriDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update materi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData materi.CreateMateriInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update materi", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.UserTable)
	inputData.User = currentUser

	updatedMateri, err := h.service.UpdateMateri(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update materi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to update materi", http.StatusOK, "success", materi.FormatMateri(updatedMateri))
	c.JSON(http.StatusOK, response)
}
