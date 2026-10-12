package admin

import (
	"go-study/models"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	BaseController
}

func (con BankController) Index(c *gin.Context) {

	//开始事务
	tx := models.DB.Begin()

	//抛出异常
	defer func() {
		if r := recover(); r != nil {
			// 遇到错误时回滚事务
			tx.Rollback()
			//转账失败
			con.error(c)
		}
	}()

	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）

	//张三账户里面减去100元
	u1 := models.Bank{Id: 1}
	tx.Find(&u1)
	u1.Balance = u1.Balance - 100
	if err := tx.Save(&u1).Error; err != nil {
		tx.Rollback()
		//转账失败
		con.error(c)
		return
	}
	// panic("异常了")

	//在李四的账户里面增加100元
	u2 := models.Bank{Id: 2}
	tx.Find(&u2)
	u2.Balance = u2.Balance + 100
	if err := tx.Save(&u2).Error; err != nil {
		tx.Rollback()
		con.error(c)
		return
	}
	tx.Commit()
	//转账成功
	con.success(c)
}
