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
	router := gin.Default()
	router.MaxMultipartMemory = 32 << 20

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		auth := api.Group("/auth")
		{
			auth.POST("/register", h.register)
			auth.POST("/login", h.login)
		}

		genre := api.Group("/genre", h.userIdentity, h.checkAdminRole)
		{
			genre.POST("/", h.createGenre)
			genre.GET("/", h.readGenre)
			genre.GET("/:id", h.readGenreById)
			genre.PUT("/:id", h.updateGenre)
			genre.DELETE("/:id", h.deleteGenre)
		}

		playlist := api.Group("/playlist", h.userIdentity)
		{
			playlist.POST("/", h.createPlaylist)
			playlist.POST("/:playlist_id/:music_id", h.addPlaylistMusic)
			playlist.GET("/", h.readPlaylist)
			playlist.GET("/:id", h.readPlaylistById)
			playlist.PUT("/:id", h.updatePlaylist)
			playlist.DELETE("/:id", h.deletePlaylist)
		}

		artist := api.Group("/artist", h.userIdentity, h.checkAdminRole)
		{
			artist.POST("/", h.createArtist)
			artist.GET("/", h.readArtist)
			artist.GET("/:id", h.readArtistById)
			artist.PUT("/:id", h.updateArtist)
			artist.DELETE("/:id", h.deleteArtist)
		}

		user := api.Group("/user", h.userIdentity)
		{
			user.GET("/", h.getUser)
			user.PUT("/", h.updateUser)
			user.PUT("/change_password", h.changeUserPassword)
		}

		music := api.Group("/music", h.userIdentity, h.checkAdminRole)
		{
			music.POST("/", h.createMusic)
			music.GET("/:id", h.getMusic)
		}
		// api.GET("/", h.GetMusicTest)
	}
	return router
}
