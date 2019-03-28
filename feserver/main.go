package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/reclaro/golab/ghapi"
)

func main() {
	// Parameters are passed in the querystring
	//http://localhost:8080/?name=fnproject/fn&search=TODO
	fmt.Println("Started")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		repoName := c.Query("name")
		search := c.Query("search")
		sq := ghapi.NewSearchQuery(repoName, search)
		result := ghapi.Search(sq)
		c.JSON(200, result)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
