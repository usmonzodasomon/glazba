package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
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

	ext := filepath.Ext(music.File.Filename)
	if err := CheckAudio(ext); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	filePath := fmt.Sprintf("./files/genre_%v/%s_%d%s", music.GenreID, music.Title, time.Now().Unix(), ext)
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

func (h *handler) playMusic(c *gin.Context) {
	logger.GetLogger().Info("playing music")

	id, err := GetIdFromParam(c, "id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logger.GetLogger().Debugf("playing music with id %v", id)

	Filepath, err := h.services.GetFilepathMusic(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.File(Filepath)
	logger.GetLogger().Infof("Music with id %v played succesfully", id)
}

func (h *handler) getMusicById(c *gin.Context) {
	logger.GetLogger().Info("Getting music")

	id, err := GetIdFromParam(c, "id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logger.GetLogger().Debugf("getting music with id %v", id)

	music, err := h.services.GetMusicById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    music,
	})
	logger.GetLogger().Infof("Music with id %v read succesfully", id)
}

func (h *handler) getMusic(c *gin.Context) {
	logger.GetLogger().Info("Getting musics")

	find := GetQueryParam(c, "find")
	artistID := GetQueryParam(c, "artist_id")
	genreID := GetQueryParam(c, "genre_id")
	releaseDataMin := GetQueryParam(c, "release_data_min")
	releaseDataMax := GetQueryParam(c, "release_data_max")
	duration := GetQueryParam(c, "duration")

	musics, err := h.services.GetMusic(find, artistID, genreID, releaseDataMin, releaseDataMax, duration)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(musics) == 0 {
		NewErrorResponse(c, http.StatusNotFound, "music not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    musics,
	})
	logger.GetLogger().Info("Musics read succesfully")
}

func (h *handler) updateMusic(c *gin.Context) {
	logger.GetLogger().Info("Updating music")

	id, err := GetIdFromParam(c, "id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logger.GetLogger().Debugf("updating music with id %v", id)

	var musicUpdate models.MusicUpdate
	if err := c.BindJSON(&musicUpdate); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateMusic(id, musicUpdate); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	logger.GetLogger().Infof("Music with id %v updated succesfully", id)
}

func (h *handler) deleteMusic(c *gin.Context) {
	id, err := GetIdFromParam(c, "id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	logger.GetLogger().Debugf("deleting music with id %v", id)

	if err := h.services.DeleteMusic(id); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	logger.GetLogger().Infof("Music with id %v deleted succesfully", id)
}
