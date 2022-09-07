package handler

import (
	"net/http"
	newsfeed "newsfeeder/cmd/platform"

	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsFeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(g *gin.Context) {
		reqbody := newsfeedPostRequest{}
		if err := g.Bind(&reqbody); err == nil {
			item := newsfeed.Item{
				Title: reqbody.Title,
				Post:  reqbody.Post,
			}
			feed.Add(item)
		}

		results := feed.GetAll()
		g.JSON(http.StatusOK, results)
	}
}
