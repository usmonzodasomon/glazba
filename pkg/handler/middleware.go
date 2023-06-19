package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	id, role, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("userId", id)
	c.Set("userRole", role)
}
