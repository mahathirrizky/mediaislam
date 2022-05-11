package handler

import (
	"mediaislam/helper"
	"mediaislam/subscribe"
	"mediaislam/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type subscribeHandler struct {
	service subscribe.Service
}

func NewSubscribeHandler(service subscribe.Service) *subscribeHandler {
	return &subscribeHandler{service}
}

func (h *subscribeHandler) GetSubscribe(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.UserTable)
	userID := currentUser.ID

	subscribes, err := h.service.GetSubscribe(userID)
	if err != nil {
		response := helper.APIResponse("Failed to Get subscribe", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("list of subscribe", http.StatusOK, "success", subscribe.FormatSubscribeList(subscribes))
	c.JSON(http.StatusOK, response)
}

func (h *subscribeHandler) CreateSubscribe(c *gin.Context) {
	var input subscribe.CreateSubscribeInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create subscribe", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.UserTable)
	input.User = currentUser

	newSubscribe, err := h.service.CreateSubscribe(input)
	if err != nil {
		response := helper.APIResponse("Failed to create subscribe", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("subscribe created", http.StatusCreated, "success", subscribe.FormatCreateSubscribe(newSubscribe))
	c.JSON(http.StatusCreated, response)
}


