package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *handler) userIdentity(c *gin.Context) {
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
	c.Set("userID", id)
	c.Set("userRole", role)
}

func (h *handler) checkAdminRole(c *gin.Context) {
	role, ok := c.Get("userRole")
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "User role not found")
	}
	if role != "admin" {
		NewErrorResponse(c, http.StatusForbidden, "you are not admin, sorry")
		return
	}
}
