package admin

import (
	"go-study/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	usersList := []models.Users{}
	models.DB.Where("age>20").Find(&usersList)
	c.JSON(200, gin.H{
		"result": usersList,
	})
}
func (con UserController) Add(c *gin.Context) {
	user := models.Users{
		Username: "itying GORM",
		Age:      22,
		Email:    "222@qq.con",
		AddTime:  int(models.GetUnix()),
	}

	models.DB.Create(&user)
	c.String(200, "增加数据成功")
}
func (con UserController) Edit(c *gin.Context) {
	user := models.Users{}
	models.DB.Where("id = ?", 13).Find(&user)
	user.Username = "哈哈"
	user.Email = "aaa@qqq.com"
	user.AddTime = int(models.GetUnix())
	models.DB.Save(&user)
	c.String(200, "用户列表-Edit------")
}
func (con UserController) Del(c *gin.Context) {
	//删除数据
	user := models.Users{}
	models.DB.Where("username = ?", "zhangsan").Delete(&user)

	c.String(200, "用户列表-Edit------")
}
