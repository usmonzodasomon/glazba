package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/logger"
	"github.com/usmonzodasomon/glazba/models"
)

func (h *handler) createGenre(c *gin.Context) {
	logger.GetLogger().Info("Creating genre")
	var input models.Genre

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateGenre(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"id":      id,
	})
	logger.GetLogger().Infof("Genre with Id %v created succesfully", id)
}

func (h *handler) readGenreMusicsById(c *gin.Context) {
	genreId, err := GetIdFromParam(c, "id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.GetLogger().Infof("Reading Genre with id %v", genreId)

	musics, err := h.services.ReadGenreMusicsById(genreId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    musics,
	})
	logger.GetLogger().Infof("Genre with name %v readed succesfully", genreId)
}

func (h *handler) readGenre(c *gin.Context) {
	logger.GetLogger().Info("Reading Genre")

	var genres []models.Genre
	if err := h.services.ReadGenre(&genres); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    genres,
	})
	logger.GetLogger().Info("Genres read succesfully")

}

func (h *handler) updateGenre(c *gin.Context) {
	logger.GetLogger().Info("Updating Genre")
	genreId, err := GetIdFromParam(c, "id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.GenreUpdateRequest
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateGenre(genreId, input.Name); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	logger.GetLogger().Infof("Genre with id %v updated succesfully", genreId)
}

func (h *handler) deleteGenre(c *gin.Context) {
	logger.GetLogger().Info("Deletating Genre")

	genreId, err := GetIdFromParam(c, "id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DeleteGenre(genreId); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	logger.GetLogger().Infof("Genre with id %v deleted succesfully", genreId)
}

func (h *handler) test(c *gin.Context) {
	musics, err := h.services.Test()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, musics)
}
