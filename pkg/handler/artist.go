package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/logger"
	"github.com/usmonzodasomon/glazba/models"
)

func (h *handler) createArtist(c *gin.Context) {
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

func (h *handler) readArtistById(c *gin.Context) {
	artistId, err := GetIdFromParam(c, "id")
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

func (h *handler) readArtist(c *gin.Context) {
	logger.GetLogger().Info("Reading Artist")

	find := GetQueryParam(c, "find")

	var artists []models.Artist
	if err := h.services.ReadArtist(&artists, find); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    artists,
	})
	logger.GetLogger().Info("Artists read succesfully")

}

func (h *handler) updateArtist(c *gin.Context) {
	logger.GetLogger().Info("Updating Artist")
	artistId, err := GetIdFromParam(c, "id")
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

func (h *handler) deleteArtist(c *gin.Context) {
	logger.GetLogger().Info("Deletating Artist")

	artistId, err := GetIdFromParam(c, "id")
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
