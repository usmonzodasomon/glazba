package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/glazba/pkg/service"
)

type handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *handler {
	return &handler{service}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	api := router.Group("api")
	{

		auth := api.Group("auth")
		{
			auth.POST("/register", h.register)
			auth.POST("/login", h.login)
		}

		category := api.Group("category", h.UserIdentity, h.CheckAdminRole)
		{
			category.POST("/", h.CreateCategory)
			category.GET("/", h.ReadCategory)
			category.GET("/:category", h.ReadCategoryByName)
			category.PUT("/:category", h.UpdateCategory)
			category.DELETE("/:category", h.DeleteCategory)
		}

		playlist := api.Group("playlist", h.UserIdentity)
		{
			playlist.POST("/", h.CreatePlaylist)
			playlist.GET("/", h.ReadPlaylist)
			playlist.GET("/:id", h.ReadPlaylistById)
			playlist.PUT("/:id", h.UpdatePlaylist)
			playlist.DELETE("/:id", h.DeletePlaylist)
		}

	}
	return router
}
