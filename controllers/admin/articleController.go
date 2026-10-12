package admin

import (
	"go-study/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (con ArticleController) Index(c *gin.Context) {
	// 一对一
	// articleList := []models.Article{}
	// models.DB.Preload("ArticleCate").Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })
	// 一对多
	articleCateList := []models.ArticleCate{}
	models.DB.Preload("ArticleXX").Find(&articleCateList)
	c.JSON(200, gin.H{
		"result": articleCateList,
	})
	// con.success(c)
}
func (con ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "-add--文章-")
}
func (con ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "-Edit---文章---")
}
