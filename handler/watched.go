package handler

import (
	"mediaislam/helper"
	"mediaislam/user"
	"mediaislam/watched"
	"net/http"

	"github.com/gin-gonic/gin"
)

type watchedHandler struct {
	service watched.Service
}

func NewWatchedHandler(service watched.Service) *watchedHandler {
	return &watchedHandler{service}
}

func (h *watchedHandler) GetWatched(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.UserTable)
	userID := currentUser.ID

	watcheds, err := h.service.GetWatched(userID)
	if err != nil {
		response := helper.APIResponse("Failed to Get watched", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("list of watched", http.StatusOK, "success", watched.FormatWatchedList(watcheds))
	c.JSON(http.StatusOK, response)
}

func (h *watchedHandler) CreateWatched(c *gin.Context) {
	var input watched.CreateWatchedInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create watched", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.UserTable)
	input.User = currentUser

	newWatched, err := h.service.CreateWatched(input)
	if err != nil {
		response := helper.APIResponse("Failed to create watched", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("watched created", http.StatusCreated, "success", watched.FormatCreateWatched(newWatched))
	c.JSON(http.StatusCreated, response)
}
