package main

import (
	"fmt"
	"go-study/models"
	"go-study/routers"
	"html/template"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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

	// 密钥, 需要16位以上
	// store := cookie.NewStore([]byte("secret1111111111123"))
	// 注意新版签名多了 username 参数:
	// NewStore(size, network, address, username, password string, keyPairs ...[]byte)
	// 没有用户名/密码也要占位, 传 ""
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", "", []byte("secret1111111111123"))
	if err != nil {
		panic(err)
	}

	// 注意: gorilla/sessions 新版(>=1.3)默认 Secure=true、SameSite=None,
	// 这会导致 http 环境下浏览器拒绝保存该 cookie, session 看起来"不生效"。
	// 本地开发必须显式关掉 Secure, 否则 http://localhost:8080 下 session 读不到值。
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24, // 一天
		HttpOnly: true,
		Secure:   false,                // 本地 http 调试必须为 false
		SameSite: http.SameSiteLaxMode, // 默认 None 需配合 Secure, 这里改用 Lax
	})
	r.Use(sessions.Sessions("mysession", store))

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run()
}
