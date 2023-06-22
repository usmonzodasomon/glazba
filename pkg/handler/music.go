package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/logger"
)

func (h *handler) CreateMusic(c *gin.Context) {
	// logger.GetLogger().Info("Creating music")
	// var music models.ProductRequest
	// if err := c.Bind(&music); err != nil {
	// 	NewErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	req := c.Request
	req.Body = http.MaxBytesReader(c.Writer, req.Body, 10<<20)
	err := req.ParseMultipartForm(10 << 20)
	// if err != nil {
	// 	NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	file, err := c.FormFile("file")
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logger.GetLogger().Debug("file got")
	filename := filepath.Join("./files", file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logger.GetLogger().Debug(filename)
	c.JSON(http.StatusOK, gin.H{
		"m": "hurah",
	})

	// id, err := h.services.CreateMusic(&music)
	// if err != nil {
	// 	NewErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// if err := c.SaveUploadedFile(music.File, music.Filepath); err != nil {
	// 	NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"message":  "success",
	// 	"id":       id,
	// 	"filepath": music.Filepath,
	// })
	// logger.GetLogger().Infof("Music with id %v created succesfully", id)
}

func (h *handler) GetMusicTest(c *gin.Context) {
	c.File("./files/music.wav")
	logger.GetLogger().Debug("I'm in getting music header")
}
