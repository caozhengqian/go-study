package admin

import (
	"go-study/models"

	"github.com/gin-gonic/gin"
)

// 测试的结构体
type UsersJson struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func (UsersJson) TableName() string {
	return "users"
}

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {

	// 使用原生sql
	// models.DB.Exec("select * from 'users' where id<?", 9)

	// 使用原生获取数据
	usersList := []models.Users{}
	models.DB.Raw("select * from users").Scan(&usersList)

	// 查询所有
	// usersList := []models.Users{}
	// models.DB.Find(&usersList)

	// 排序Order(desc倒叙)
	// usersList := []models.Users{}
	// models.DB.Order("id desc").Order("age asc").Find(&usersList)

	// Limit限制返回2条，offset跳过2条，取3和4条
	// usersList := []models.Users{}
	// models.DB.Limit(2).Offset(2).Find(&usersList)

	// Count 统计数量
	// usersList := []models.Users{}
	// var num int64
	// models.DB.Find(&usersList).Count(&num)
	// c.JSON(200, gin.H{
	// 	"result": num,
	// })

	// Select（id,email)
	// usersList := []UsersJson{}
	// models.DB.Select("id,email").Find(&usersList)

	// 查询一条数据
	// usersList := models.Users{Id: 2} //不需要切片
	// models.DB.Find(&usersList)

	// Where条件 "<" "<=" ">" ">=" "!="
	// usersList := []models.Users{}
	// models.DB.Where("id<?", 6).Find(&usersList)

	// Where条件 "AND"
	// usersList := []models.Users{}
	// models.DB.Where("id<? AND id>?", 12, 6).Find(&usersList)

	// Where条件 "OR"
	// usersList := []models.Users{}
	// models.DB.Where("id<? ", 12).Or("or id>?", 6).Find(&usersList)

	// Where条件 "IN"
	// usersList := []models.Users{}
	// models.DB.Where("id IN (?)", []int{8, 9, 10}).Find(&usersList)

	// Where条件 "LIKE"
	// usersList := []models.Users{}
	// models.DB.Where("email LIKE ?", "%qq.com%").Find(&usersList)

	// Where条件 "BETWEEN AND""NOT BETWEEN AND"
	// usersList := []models.Users{}
	// models.DB.Where("id BETWEEN ? AND ?", 1, 3).Find(&usersList)

	// c.JSON(200, gin.H{
	// 	"result": usersList,
	// })

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
