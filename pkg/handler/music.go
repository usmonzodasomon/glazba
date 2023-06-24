package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/logger"
	"github.com/usmonzodasomon/glazba/models"
)

func (h *handler) createMusic(c *gin.Context) {
	logger.GetLogger().Info("Creating music")
	var music models.MusicRequest
	if err := c.Bind(&music); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	filePath := fmt.Sprintf("./files/genre_%v/%s_%d.mp3", music.GenreID, music.Title, time.Now().Unix())
	filePath = strings.ReplaceAll(filePath, " ", "_")
	if err := c.SaveUploadedFile(music.File, filePath); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.CreateMusic(&music, filePath)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"id":      id,
	})
}

func (h *handler) getMusic(c *gin.Context) {
	logger.GetLogger().Info("Getting musics")

	id, err := GetIdFromParam(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logger.GetLogger().Debugf("getting music with id %v", id)

	filePath, err := h.services.GetMusicById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logger.GetLogger().Debugf("getting music with filePath %s", filePath)
	c.File(filePath)
	logger.GetLogger().Infof("Music with id %v read succesfully", id)
}
