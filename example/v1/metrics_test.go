package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	app.GET("api/rand_str/:str", func(ctx *gin.Context) {
		str := ctx.Param("str")
		ctx.JSON(200, gin.H{
			"message": str,
		})
	})
	app.GET("api/rand_int/:str", func(ctx *gin.Context) {
		str := ctx.Param("str")
		ctx.JSON(200, gin.H{
			"message": str,
		})
	})

}
