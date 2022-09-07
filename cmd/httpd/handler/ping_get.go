package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGet(g *gin.Context) {
	g.JSON(http.StatusOK, map[string]string{
		"hello": "found me",
	})
}

func HelloGet() gin.HandlerFunc {
	return func(g *gin.Context) {
		g.JSON(http.StatusOK, map[string]string{
			"hello": "found me",
		})
	}
}
