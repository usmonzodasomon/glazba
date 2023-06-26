package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/logger"
	"github.com/usmonzodasomon/glazba/models"
)

func (h *handler) createPlaylist(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.GetLogger().Infof("Creating playlist for user with id %v", userId)

	var input models.Playlist

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreatePlaylist(&input, userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"id":      id,
	})
	logger.GetLogger().Infof("Playlist created succesfully with id %v", id)
}

func (h *handler) readPlaylist(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	find := GetQueryParam(c, "find")
	logger.GetLogger().Infof("Reading Playlist from user with id %v", userId)

	var playlists []models.Playlist
	if err := h.services.ReadPlaylists(&playlists, userId, find); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    playlists,
	})
	logger.GetLogger().Infof("Playlists from user %v read success", userId)
}

func (h *handler) readPlaylistById(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := GetIdFromParam(c, "id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.GetLogger().Infof("Reading Playlist with id %v from user with id %v", id, userId)

	playlist, err := h.services.ReadPlaylistById(id, userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    playlist,
	})
	logger.GetLogger().Infof("Playlist with id %v read succesfully", id)
}

func (h *handler) updatePlaylist(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := GetIdFromParam(c, "id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.GetLogger().Infof("Updating Playlist with id %v from user with id %v", id, userId)

	var PlaylistUpdateRequest models.PlaylistUpdateRequest
	if err := c.BindJSON(&PlaylistUpdateRequest); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.UpdatePlaylist(id, userId, &PlaylistUpdateRequest); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	logger.GetLogger().Infof("Playlist with id %v updated succesfully", id)
}

func (h *handler) deletePlaylist(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := GetIdFromParam(c, "id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.GetLogger().Infof("Deleting Playlist with id %v from user with id %v", id, userId)

	if err := h.services.DeletePlaylist(id, userId); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
