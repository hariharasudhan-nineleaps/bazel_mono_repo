package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hariharasudhan-nineleaps/bazel_mono_repo/shared"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		shared.Log("ping request receieved....")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/encode", func(c *gin.Context) {
		shared.Log("encode request receieved....")
		c.JSON(http.StatusOK, gin.H{
			"encoded": shared.Encode(),
		})
	})

	r.GET("/decode", func(c *gin.Context) {
		shared.Log("decode request receieved....")
		c.JSON(http.StatusOK, gin.H{
			"decoded": shared.Decode(),
		})
	})

	r.Run()
}
