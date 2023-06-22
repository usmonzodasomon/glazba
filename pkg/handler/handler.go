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
	router.MaxMultipartMemory = 10 << 20

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

		genre := api.Group("genre", h.UserIdentity, h.CheckAdminRole)
		{
			genre.POST("/", h.CreateGenre)
			genre.GET("/", h.ReadGenre)
			genre.GET("/:id", h.ReadGenreById)
			genre.PUT("/:id", h.UpdateGenre)
			genre.DELETE("/:id", h.DeleteGenre)
		}

		playlist := api.Group("playlist", h.UserIdentity)
		{
			playlist.POST("/", h.CreatePlaylist)
			playlist.GET("/", h.ReadPlaylist)
			playlist.GET("/:id", h.ReadPlaylistById)
			playlist.PUT("/:id", h.UpdatePlaylist)
			playlist.DELETE("/:id", h.DeletePlaylist)
		}

		artist := api.Group("artist", h.UserIdentity, h.CheckAdminRole)
		{
			artist.POST("/", h.CreateArtist)
			artist.GET("/", h.ReadArtist)
			artist.GET("/:id", h.ReadArtistById)
			artist.PUT("/:id", h.UpdateArtist)
			artist.DELETE("/:id", h.DeleteArtist)
		}

		music := api.Group("music", h.UserIdentity, h.CheckAdminRole)
		{
			music.POST("/", h.CreateMusic)
		}
		api.GET("/", h.GetMusicTest)

	}
	return router
}
