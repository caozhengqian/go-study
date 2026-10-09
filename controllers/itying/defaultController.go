package itying

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) News(c *gin.Context) {
	username, _ := c.Cookie("username")
	fmt.Println("cookie中username的值为：", username)
	c.String(200, "News")
}
