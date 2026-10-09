package admin

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("usernames", "lisi")
	session.Save()
	con.success(c)
}
func (con UserController) Add(c *gin.Context) {
	session := sessions.Default(c)
	usernames := session.Get("usernames")
	fmt.Println("session中usernames的值为：", usernames)
	c.String(200, "用户列表-add---")
}
func (con UserController) Edit(c *gin.Context) {
	c.String(200, "用户列表-Edit------")
}
