package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(200, "首页")
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "新闻")
		})
	}
	//http://localhost:8080/api/userlist
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "我是一个api接口")
		})
		apiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(200, "我是一个api接口-userlist")
		})
	}
	r.Run()
}
