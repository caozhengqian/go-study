package admin

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	//key，value，过期时间，路径，域名，是否https，是否允许js访问
	c.SetCookie("username", "小曹一", 3600, "/", "localhost", false, true)
	c.String(200, "用户列表--")
}
