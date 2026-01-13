package http

import (
	"github.com/elginas/onyx-backend/internal/config"
	"github.com/gin-gonic/gin"
)

func Run(config config.Config) {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run(config.Addr)
}
