package gin_metrics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var (
	taskCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Subsystem: "workpool",
		Name:      "complete_task_total",
		Help:      "Total number of task completed",
	})
)

type GinRoutesInfo struct {
	Method  string
	Path    string
	Handler string
}

var routeInfo []GinRoutesInfo

func init() {
	prometheus.MustRegister(taskCounter)
}

func Default(app *gin.Engine) *gin.Engine {
	app.Use(metricsMiddleware)

	metricsHandler(app)
	generateRouteInfo(app)
	return app
}

func metricsMiddleware(c *gin.Context) {
	//path := c.Request.URL.Path
	start := time.Now()
	c.Next()
	taskCounter.Add(1)
	//fmt.Println(c.ClientIP(), c.Request.URL.Path, start.Sub(time.Now()))
	fmt.Println(time.Now().Sub(start))
	return
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

	//go func() {
	//	for {
	//		taskCounter.Add(1)
	//		time.Sleep(1 * time.Second)
	//	}
	//}()

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
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
