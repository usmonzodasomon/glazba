package handler

import (
	"errors"
	"fmt"
	"mime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/logger"
)

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewErrorResponse(c *gin.Context, status int, err string) {
	logger.GetLogger().Error(err)
	c.AbortWithStatusJSON(status, errorResponse{
		Message: "error",
		Error:   err,
	})
}

func GetUserId(c *gin.Context) (uint, error) {
	userId, ok := c.Get("userID")
	if !ok {
		return 0, errors.New("error while getting userId from header")
	}
	id, ok := userId.(uint)
	if !ok {
		return 0, errors.New("error while convertation userId to uint")
	}
	return id, nil
}

func GetIdFromParam(c *gin.Context, paramName string) (uint, error) {
	idInt, err := strconv.Atoi(c.Param(paramName))
	if err != nil {
		return 0, err
	}

	idUint := uint(idInt)
	return idUint, nil

}

func CheckAudio(ext string) error {
	fileType := mime.TypeByExtension(ext)

	fileTypeParts := strings.Split(fileType, "/")
	logger.GetLogger().Debug(ext, fileTypeParts)
	if len(fileTypeParts) == 0 || string(fileTypeParts[0]) != "audio" {
		return fmt.Errorf("error audio format")
	}
	return nil
}

func GetQueryParam(c *gin.Context, queryParam string) string {
	return c.Query(queryParam)
}
