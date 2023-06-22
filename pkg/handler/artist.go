package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/logger"
	"github.com/usmonzodasomon/glazba/models"
)

func (h *handler) CreateArtist(c *gin.Context) {
	logger.GetLogger().Info("Creating artist")
	var input models.Artist

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateArtist(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"id":      id,
	})
	logger.GetLogger().Infof("Artist with Id %v created succesfully", id)
}

func (h *handler) ReadArtistById(c *gin.Context) {
	artistId, err := GetIdFromParam(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.GetLogger().Infof("Reading Artist with id %v", artistId)

	artist, err := h.services.ReadArtistById(artistId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    artist,
	})
	logger.GetLogger().Infof("Artist with name %v readed succesfully", artistId)
}

func (h *handler) ReadArtist(c *gin.Context) {
	logger.GetLogger().Info("Reading Artist")

	var artists []models.Artist
	if err := h.services.ReadArtist(&artists); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    artists,
	})
	logger.GetLogger().Info("Artists read succesfully")

}

func (h *handler) UpdateArtist(c *gin.Context) {
	logger.GetLogger().Info("Updating Artist")
	artistId, err := GetIdFromParam(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.ArtistUpdateRequest
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateGenre(artistId, input.Name); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	logger.GetLogger().Infof("Artist with id %v updated succesfully", artistId)
}

func (h *handler) DeleteArtist(c *gin.Context) {
	logger.GetLogger().Info("Deletating Artist")

	artistId, err := GetIdFromParam(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DeleteArtist(artistId); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	logger.GetLogger().Infof("Artist with id %v deleted succesfully", artistId)
}
