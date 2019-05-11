package gin_exporter

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinRoutesInfo struct {
	Method  string
	Path    string
	Handler string
}

var routeInfo []GinRoutesInfo

func Default(app *gin.Engine) *gin.Engine {
	metricsHandler(app)
	generateRouteInfo(app)

	app.Use(metricsMiddleware())
	return app
}

func metricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.ClientIP())
		c.Next()
	}
}

//	metrics 接口
func metricsHandler(app *gin.Engine) {

	app.GET("metrics", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":   200,
			"msg":    "metrics",
			"routes": routeInfo,
		})
	})
}

// 初始化 Router信息
func generateRouteInfo(app *gin.Engine) {
	for _, v := range app.Routes() {
		routeInfo = append(routeInfo, GinRoutesInfo{
			Method:  v.Method,
			Path:    v.Path,
			Handler: v.Handler,
		})
	}
}
