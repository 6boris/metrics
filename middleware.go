package gin_metrics

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func metricsMiddleware(c *gin.Context) {
	//path := c.Request.URL.Path

	GinRequestTotalCount.With(prometheus.Labels{"method": c.Request.Method, "path": c.Request.URL.Path}).Inc()

	//start := time.Now()
	c.Next()
	//fmt.Println(c.ClientIP(), c.Request.URL.Path, start.Sub(time.Now()))
	//costTime := time.Since(start)
	//respTime.WithLabelValues(fmt.Sprintf("%d", costTime)).Observe(costTime.Seconds())

	//

	//fmt.Println(path)
}
