package gin_metrics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	app := gin.New()
	Default(app)

	gin.SetMode(gin.DebugMode)

	app.GET("rand_str", func(c *gin.Context) {
		c.JSON(200, getRandomString(32))
	})
	app.GET("rand_int", func(c *gin.Context) {
		c.JSON(200, rand.Intn(100000000000))
	})
	app.GET("time_now", func(c *gin.Context) {
		c.JSON(200, time.Now())
	})
	app.GET("rand_sleep", func(c *gin.Context) {
		time.Sleep(time.Duration(time.Duration(rand.Intn(1000)) * time.Millisecond))
		c.JSON(200, time.Now())
	})
	for i := 0; i < 100; i++ {
		app.GET(fmt.Sprint("rand_route", strconv.Itoa(i)), func(c *gin.Context) {
			c.JSON(200, time.Now())
		})
	}

	//if err := app.Run("127.0.0.1:9000"); err != nil {
	//	panic(err.Error())
	//}
}

func BenchmarkDefault(b *testing.B) {
	b.N = 10000
	var wg sync.WaitGroup
	wg.Add(1)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				Req("http://127.0.0.1:9000/rand_str")
				Req("http://127.0.0.1:9000/rand_int")
				Req("http://127.0.0.1:9000/time_now")
			}
		}()
	}

	go func() {
		for i := 0; i < b.N; i++ {
			Req("http://127.0.0.1:9000/rand_sleep")
		}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			Req("http://127.0.0.1:9000/rand_route" + strconv.Itoa(rand.Intn(100)))
		}
	}()
	time.Sleep(10 * time.Second)
	//wg.Wait()

}

//生成随机字符串
func getRandomString(str_len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < str_len; i++ {
		result = append(result, bytes[r.Intn(str_len)])
	}
	return string(result)
}

func Req(url string) {
	time.Sleep(time.Duration(time.Duration(rand.Intn(100)) * time.Millisecond))

	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//fmt.Println(string(body))
}
