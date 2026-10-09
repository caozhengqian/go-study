package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	//判断用户是否登录
	fmt.Println("api路由的中间件", c.Request.URL)
	c.Set("username", "张三")

	//定义一个goroutine统计日志  当在中间件或 handler 中启动新的 goroutine 时，不能使用原始的上下文（c *gin.Context）， 必须使用其只读副本（c.Copy()）
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("协程取值 " + cCp.Request.URL.Path)
	}()
}
