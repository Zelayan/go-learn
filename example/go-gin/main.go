package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.GET("panic", func(context *gin.Context) {
		// 在协程中的panic 是无法被gin的中间件捕获的
		/*		go func() {
				panic("xx")
			}()*/
	})
	r.Run()
}
