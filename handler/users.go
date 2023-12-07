package handler

import (
	"net/http"
	"test/helper"
	"test/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//get input dari user
	//map intput dari user ke struct
	//struct diatas kita passing ke parameter

	var input user.RegisterUser

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Account Cant Be Registered", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	formatter := user.FormatUserAPI(newUser, "newtokenauth")
	response := helper.APIResponse("Account Has Been Registered", http.StatusOK, "success", formatter)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, response)

}
