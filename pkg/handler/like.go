package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) addMusicLike(c *gin.Context) {
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

	if err := h.services.AddMusicLike(userID, musicID); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *handler) deleteMusicLike(c *gin.Context) {
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

	if err := h.services.DeleteMusicLike(userID, musicID); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
