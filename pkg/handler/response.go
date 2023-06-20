package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/logger"
)

type errorResponse struct {
	Message string `json:"message" binding:"required"`
}

func NewErrorResponse(c *gin.Context, status int, message string) {
	logger.GetLogger().Error(message)
	c.AbortWithStatusJSON(status, errorResponse{Message: message})
}

func GetUserId(c *gin.Context) (uint, error) {
	userId, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("error while getting userId from header")
	}
	id, ok := userId.(uint)
	if !ok {
		return 0, errors.New("error while convertation userId to uint")
	}
	return id, nil
}

func GetIdFromParam(c *gin.Context) (uint, error) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}

	idUint := uint(idInt)
	return idUint, nil

}
