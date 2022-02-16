package main

import (
	"fmt"
	"nicholas/newsfeeder/platform/newsfeed"
	"nicholas/newsfeeder/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
  // r := gin.Default()
	// r.GET("/ping", handler.PingGet())
	// r.Run()

  feed := newsfeed.New()

  fmt.Println(feed)

  feed.Add(newsfeed.Item{"Hello", "How ya' doing mate?"})

  fmt.Println(feed)
}
