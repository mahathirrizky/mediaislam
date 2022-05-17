package handler

import (
	"fmt"
	"mediaislam/helper"
	"mediaislam/user"
	"mediaislam/video"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	service video.Service
}

func NewVideoHandler(service video.Service) *VideoHandler {
	return &VideoHandler{service}
}

func (h *VideoHandler) CreateVideo(c *gin.Context) {
	var input video.CreateVideoInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create video", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newVideo, err := h.service.CreateVideo(input)
	if err != nil {
		response := helper.APIResponse("Failed to create video", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("video created", http.StatusCreated, "success", video.FormatVideo(newVideo))
	c.JSON(http.StatusCreated, response)
}

func (h *VideoHandler) UpdateVideo(c *gin.Context) {
	var input video.GetVideoDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("error to get video", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData video.CreateVideoInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update video", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newVideo, err := h.service.UpdateVideo(input, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update video", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("video updated", http.StatusOK, "success", video.FormatVideo(newVideo))
	c.JSON(http.StatusOK, response)
}

func (h *VideoHandler) GetVideo(c *gin.Context) {
	var input video.GetVideoDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("error to get video", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getVideo, err := h.service.GetVideo(input)
	if err != nil {
		response := helper.APIResponse("Failed to get video", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("video found", http.StatusOK, "success", video.FormatVideo(getVideo))
	c.JSON(http.StatusOK, response)
}

func (h *VideoHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.UserTable)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_,err = h.service.SaveImage(userID, file.Filename)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Image uploaded", http.StatusOK, "success", data)
	
	c.JSON(http.StatusOK, response)
}

func (h *VideoHandler) GetTematikList(c *gin.Context){	
	userID, _ := strconv.Atoi(c.Query("user_id"))

	tematik, err := h.service.GetVideoTematik(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get video tematik", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("video tematik found", http.StatusOK, "success", video.FormatVideoList(tematik))
	c.JSON(http.StatusOK, response)
}

func (h *VideoHandler) GetShortList(c *gin.Context){
	userID, _ := strconv.Atoi(c.Query("user_id"))

	short, err := h.service.GetVideoShort(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get video short", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("video short found", http.StatusOK, "success", video.FormatVideoList(short))
	c.JSON(http.StatusOK, response)
}
