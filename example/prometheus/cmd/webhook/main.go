package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()

	r.GET("/ping", ping)
	r.POST("/webhook", webhook)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func webhook(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	c.JSON(http.StatusOK, reqBody)
}

