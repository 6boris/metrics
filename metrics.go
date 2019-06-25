package ginMetrics

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	GinRequestTotalCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gin_request_total",
			Help: "Number of hello requests in total",
		},
		[]string{"method", "endpoint"},
	)

	GinRequestGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gin_request_total2",
			Help: "Number of hello requests in total",
		},
		[]string{"method", "endpoint"},
	)

	GinRequestHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gin_request_total3",
			Help:    "Number of hello requests in total",
			Buckets: []float64{},
		},
		[]string{"method", "endpoint"},
	)

	GinRequestSummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "gin_request_time_millisecond",
			Help:       "Number of hello requests in total",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"method", "endpoint"},
	)
)

type Context struct {
	Request *http.Request

	index int8

	// Keys is a key/value pair exclusively for the context of each request.
	Keys map[string]interface{}

	// Errors is a list of errors attached to all the handlers/middlewares who used this context.

	// Accepted defines a list of manually accepted formats for content negotiation.
	Accepted []string
}

type GinRoutesInfo struct {
	Method  string
	Path    string
	Handler string
}

var routeInfo []GinRoutesInfo

func init() {
	//prometheus.MustRegister(GinRequestTotalCount)
	//prometheus.MustRegister(GinRequestGauge)
	//prometheus.MustRegister(GinRequestHistogram)
	//prometheus.MustRegister(GinRequestSummary)
	prometheus.MustRegister(GinRequestTotalCount, GinRequestGauge, GinRequestHistogram, GinRequestSummary)
	//registerHandler()
}

func registerHandler() {
	if err := prometheus.Register(GinRequestTotalCount); err != nil {
		log.Panicln(err.Error())
	}

	if err := prometheus.Register(GinRequestGauge); err != nil {
		log.Panicln(err.Error())
	}

	if err := prometheus.Register(GinRequestHistogram); err != nil {
		log.Panicln(err.Error())
	}
	if err := prometheus.Register(GinRequestSummary); err != nil {
		log.Panicln(err.Error())
	}
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
	context.Background()
	for _, v := range app.Routes() {
		routeInfo = append(routeInfo, GinRoutesInfo{
			Method:  v.Method,
			Path:    v.Path,
			Handler: v.Handler,
		})
	}
}
