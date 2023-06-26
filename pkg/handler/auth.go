package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/logger"
	"github.com/usmonzodasomon/glazba/models"
)

func (h *handler) register(c *gin.Context) {
	logger.GetLogger().Info("Registrating user")

	var input models.RegisterData

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"id":      id,
	})
	logger.GetLogger().Infof("User with Id %v created succesfully", id)
}

func (h *handler) login(c *gin.Context) {
	logger.GetLogger().Info("Loggining")

	var input models.LoginData
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})
	logger.GetLogger().Infof("User %v logged in succesfully", input.Login)
}
