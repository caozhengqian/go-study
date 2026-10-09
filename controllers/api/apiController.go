package api

import (
	"fmt"
	"go-study/models"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (con ApiController) Index(c *gin.Context) {
	fmt.Println(models.UnixToTime(1685664000))
	username, _ := c.Get("username")
	//类型断言
	v, ok := username.(string)
	if ok == true {
		c.String(200, "我是一个api接口-Index--"+v)
		return
	}
	c.String(200, "我是一个api接口")
}
func (con ApiController) Userlist(c *gin.Context) {
	c.String(200, "我是一个api接口-Userlist")
}
func (con ApiController) Plist(c *gin.Context) {
	c.String(200, "我是一个api接口-Plist")
}
