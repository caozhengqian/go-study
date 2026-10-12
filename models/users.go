package models

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	AddTime  int    `json:"addTime"`
}

// 表示配置操作数据库的表名称
func (Users) TableName() string {
	return "users"
}
