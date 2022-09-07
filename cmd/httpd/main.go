package main

import (
	"fmt"
	"newsfeeder/cmd/httpd/handler"
	newsfeed "newsfeeder/cmd/platform"

	"github.com/gin-gonic/gin"
)

// go mod init newsfeeder
// go test .\cmd\platform\
func main() {

	//singelton dependecy
	feed := newsfeed.New()

	fmt.Println("we are good to go")
	r := gin.Default()

	r.GET("/ping", handler.PingGet)
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))

	r.Run(":8090")

}
