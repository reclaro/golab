package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/reclaro/golab/ghapi"
)

func main() {
	fmt.Println("Started")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		sq := ghapi.NewSearchQuery("fnproject/fn", "TODO")
		result := ghapi.Search(sq)
		c.JSON(200, result)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
