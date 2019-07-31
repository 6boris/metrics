package main

import (
	"github.com/gin-gonic/gin"
	ginMetrics "github.com/kylesliu/gin-metrics"
)

func main() {
	app := gin.New()
	ginMetrics.Default(app)
	app.GET("api/rand_str/:str", func(ctx *gin.Context) {
		str := ctx.Param("str")
		ctx.JSON(200, gin.H{
			"message": str,
		})
	})
	app.GET("api/rand_int/:number", func(ctx *gin.Context) {
		number := ctx.Param("number")
		ctx.JSON(200, gin.H{
			"message": number,
		})
	})
	if err := app.Run(); err != nil {
		panic(err)
	}


	
}
