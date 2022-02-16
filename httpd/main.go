package main

import (
	"nicholas/newsfeeder/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

	r.GET("/ping", handler.PingGet())
  r.GET("/newsfeed", handler.NewsfeedGet())
  r.GET("/newsfeed", handler.NewsfeedPost())

	r.Run()

}
