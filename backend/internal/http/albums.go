package http

import (
	"net/http"

	"github.com/elginas/onyx-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, domain.GetAlbums())
}
