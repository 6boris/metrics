package ginMetrics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func metricsMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()

	//	Request time
	GinRequestTotalCount.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.FullPath()}).Inc()

	//GinRequestGauge.With(prometheus.Labels{"method": c.Request.Method, "path": c.Request.URL.Path}).
	//fmt.Println(time.Now().UnixNano() - start)
	//fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().Sub(start))
	//fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)

	//start := time.Now()
	c.Next()
	//fmt.Println(c.ClientIP(), c.Request.URL.Path, start.Sub(time.Now()))
	//costTime := time.Since(start)
	//respTime.WithLabelValues(fmt.Sprintf("%d", costTime)).Observe(costTime.Seconds())

	//
	end := time.Now().UnixNano()
	use_time := end - start
	GinRequestSummary.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.FullPath()}).Observe(float64(use_time))
	fmt.Println(c.FullPath())

	//use_time := float64(end / 1e6)
	//fmt.Println((end - start) / 1e6)
	//fmt.Println()

	//fmt.Println(path)
}
