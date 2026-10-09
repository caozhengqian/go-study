package routers

import (
	"go-study/controllers/api"
	"go-study/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	apiRouters.Use(middlewares.InitMiddleware) //单个路由，单个过滤器
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userlist", api.ApiController{}.Userlist)
		apiRouters.GET("/plist", api.ApiController{}.Plist)
	}

}
