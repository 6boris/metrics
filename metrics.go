package gin_metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	GinRequestTotalCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gin_reqeust_total",
			Help: "Number of hello requests in total",
		},
		[]string{"method", "path"},
	)
)

type GinRoutesInfo struct {
	Method  string
	Path    string
	Handler string
}

var routeInfo []GinRoutesInfo

func init() {
	//prometheus.MustRegister(taskCounter)
	prometheus.MustRegister(GinRequestTotalCount)
	//prometheus.MustRegister(respTime)
	//prometheus.MustRegister(respSum)
}

func Default(app *gin.Engine) {
	app.Use(metricsMiddleware)
	app.GET("metrics", metricsHandler())
	//app.GET("metrics.json", metricsHandlerJson)
	generateRouteInfo(app)
}

//	Gin Metrics api handler
func metricsHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func metricsHandlerJson(c *gin.Context) {
	c.JSON(200, gin.H{
		"metrics": routeInfo,
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
