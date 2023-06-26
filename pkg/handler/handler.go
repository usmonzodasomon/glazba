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

		genre := api.Group("/genre")
		{
			genre.GET("/", h.readGenre)
			genre.GET("/:id", h.readGenreMusicsById)
		}

		genre = api.Group("/genre", h.userIdentity, h.checkAdminRole)
		{
			genre.POST("/", h.createGenre)
			genre.PUT("/:id", h.updateGenre)
			genre.DELETE("/:id", h.deleteGenre)
		}

		playlist := api.Group("/playlist", h.userIdentity)
		{
			playlist.POST("/", h.createPlaylist)
			playlist.GET("/", h.readPlaylist)
			playlist.GET("/:id", h.readPlaylistById)
			playlist.PUT("/:id", h.updatePlaylist)
			playlist.DELETE("/:id", h.deletePlaylist)
		}

		artist := api.Group("/artist")
		{
			artist.GET("/", h.readArtist)
			artist.GET("/:id", h.readArtistById)
		}

		artist = api.Group("/artist", h.userIdentity, h.checkAdminRole)
		{
			artist.POST("/", h.createArtist)
			artist.PUT("/:id", h.updateArtist)
			artist.DELETE("/:id", h.deleteArtist)
		}

		user := api.Group("/user", h.userIdentity)
		{
			user.GET("/", h.getUser)
			user.PUT("/", h.updateUser)
			user.PUT("/change_password", h.changeUserPassword)
		}

		music := api.Group("/music", h.userIdentity)
		{
			music.GET("/", h.getMusic)
			music.GET("/:id", h.getMusicById)
			music.GET("/:id/play", h.playMusic)
		}
		music = music.Group("/", h.checkAdminRole)
		{
			music.POST("/", h.createMusic)
			music.PUT("/:id", h.updateMusic)
			music.DELETE("/:id", h.deleteMusic)
		}

		playlistMusic := playlist.Group("/music", h.userIdentity)
		{
			playlistMusic.GET("/:playlist_id", h.getPlaylistMusics)
			playlistMusic.POST("/:playlist_id/:music_id", h.addPlaylistMusic)
			playlistMusic.POST("/favorites/:music_id", h.addFavoriteMusic)
			playlistMusic.DELETE("/:playlist_id/:music_id", h.deletePlaylistMusic)
			playlistMusic.DELETE("/favorites/:music_id", h.deleteFavoriteMusic)
		}

		like := api.Group("/music", h.userIdentity)
		{
			like.POST("/like/:music_id", h.addMusicLike)
			like.DELETE("/like/:music_id", h.deleteMusicLike)
		}
		api.GET("/test", h.test)
	}
	return router
}
