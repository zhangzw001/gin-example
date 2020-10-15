package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/zhangzw001/gin-example/ginBlog/models"
	"github.com/zhangzw001/gin-example/ginBlog/pkg/e"
	"github.com/zhangzw001/gin-example/ginBlog/pkg/setting"
	"github.com/zhangzw001/gin-example/ginBlog/pkg/util"
	"log"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		maps["state"] = state
	}

	code := e.SUCCESS
	var err error
	data["lists"],err  = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	if err != nil {
		log.Fatal(err)
	}
	//data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}