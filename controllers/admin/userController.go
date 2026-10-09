package admin

import (
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
	username := session.Get("usernames")

	c.String(200, "用户列表-add---", username)
}
func (con UserController) Edit(c *gin.Context) {
	c.String(200, "用户列表-Edit------")
}
