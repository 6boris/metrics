# The Gin Framework Metrics Middleware

<p align="center">
    <a href="https://www.travis-ci.org/kylesliu/gin_metrics">
        <img src="https://www.travis-ci.org/kylesliu/gin_metrics.svg?branch=master" alt="Build Status">
    </a>
    <a href="https://codecov.io/gh/kylesliu/gin_metrics">
        <img src="https://codecov.io/gh/kylesliu/gin_metrics/branch/master/graph/badge.svg" />
    </a>
    <a href="https://img.shields.io/github/stars/kylesliu/gin_metrics.svg?label=Stars&style=social">
        <img src="https://golangci.com/badges/github.com/kylesliu/gin_metrics.svg" />
    </a>
    <a href="https://img.shields.io/github/stars/kylesliu/awesome-golang-leetcode.svg?label=Stars&style=social">
        <img src="https://img.shields.io/badge/All_Contributors-1-blue.svg" />
    </a>
    <a href="http://hits.dwyl.io/kylesliu/gin_metrics">
        <img src="http://hits.dwyl.io/kylesliu/gin_metrics.svg" />
    </a>
</p>



## Preface
Many small companies don't have such a large architecture for micro-services when they do websites. A simple solution for viewing application traffic is very important. This repository is a middleware that integrates seamlessly with Gin.

![gin_metrics_v1](https://oss.kyle.link/images/2019/gin_metrics_v1.png)

## How to use 

* install the metrics lib

```bash
go get github.com/kylesliu/gin_metrics
```

* run the server

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin_metrics"
)

func main() {
	app := gin.Default()
	gin.SetMode(gin.DebugMode)

	app.GET("demo1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "demo1",
		})
	})

	gin_metrics.Default(app)

	if err := app.Run("127.0.0.1:9000"); err != nil {
		panic(err.Error())
	}
}
```


* Config the Prometheus

```yaml
  - job_name: 'gin_metrics'
    static_configs:
    - targets: ['localhost:9000']
```

* Config the Grafana

[Grafana Dashboard](https://snapshot.raintank.io/dashboard/snapshot/YELhgZTaIuynoKd3UPudNJdNBgDy83CC)

## Last
If you have any good suggestions to mention issue or PR, I will check it out in detail.
