package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) addPlaylistMusic(c *gin.Context) {
	playlistID, err := GetIdFromParam(c, "playlist_id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	musicID, err := GetIdFromParam(c, "music_id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.AddPlaylistMusic(userID, playlistID, musicID); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *handler) addFavoriteMusic(c *gin.Context) {
	userID, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	musicID, err := GetIdFromParam(c, "music_id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.AddFavoriteMusic(userID, musicID); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *handler) deletePlaylistMusic(c *gin.Context) {
	playlistID, err := GetIdFromParam(c, "playlist_id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	musicID, err := GetIdFromParam(c, "music_id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.DeletePlaylistMusic(userID, playlistID, musicID); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *handler) deleteFavoriteMusic(c *gin.Context) {
	userID, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	musicID, err := GetIdFromParam(c, "music_id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DeleteFavoriteMusic(userID, musicID); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *handler) getPlaylistMusics(c *gin.Context) {
	playlistID, err := GetIdFromParam(c, "playlist_id")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	musics, err := h.services.GetPlaylistMusics(playlistID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    musics,
	})

}
