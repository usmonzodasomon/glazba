package handler

import "github.com/gin-gonic/gin"

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *handler) register(c *gin.Context) {

}

func (h *handler) login(c *gin.Context) {

}
