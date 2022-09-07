package handler

import (
	"net/http"
	newsfeed "newsfeeder/cmd/platform"

	"github.com/gin-gonic/gin"
)

func NewsFeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(g *gin.Context) {
		results := feed.GetAll()
		g.JSON(http.StatusOK, results)
	}
}
