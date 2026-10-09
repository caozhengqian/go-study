package main

import (
	"fmt"
	"go-study/models"
	"go-study/routers"
	"html/template"

	"github.com/gin-gonic/gin"
)

func initMiddlewareOne(c *gin.Context) {

	fmt.Println("开始-第一个中间件initMiddlewareOne")
	//调用该请求的剩余处理程序
	c.Next()

	fmt.Println("结束-第一个中间件initMiddlewareOne")

}
func initMiddlewareTwo(c *gin.Context) {

	fmt.Println("开始-第二个中间件initMiddlewareTwo")
	//调用该请求的剩余处理程序
	c.Next()

	fmt.Println("结束-第二个中间件initMiddlewareTwo")

}
func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前,
	r.SetFuncMap(template.FuncMap{
		"Totime": models.UnixToTime,
	})
	//全局中间件
	r.Use(initMiddlewareOne, initMiddlewareTwo)
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run()
}
