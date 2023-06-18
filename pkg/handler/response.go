package handler

import (
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
