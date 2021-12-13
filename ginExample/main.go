package main

import (
	"time"

	_ "encoding/json"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pingTime", func(c *gin.Context) {
		//JSON Serializer available in gin context
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})
	r.Run(":8080")
}

//curl -X GET "http://localhost:8080/pingTime"

//{"serverTime":"2020-02-27T19:08:05.470955Z"}
