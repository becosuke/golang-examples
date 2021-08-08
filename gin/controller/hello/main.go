package main

import (
	"example/hello-gin/service/hello"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", Index)
	r.Run()
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, hello.GetHello())
}
