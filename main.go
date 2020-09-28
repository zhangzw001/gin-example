package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)



func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/someJSON",func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag": "<br>",
		}
		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK,data)

	})

	r.MaxMultipartMemory = 8 << 20 //8 Mib
	r.POST("/upload",func(c *gin.Context) {
		//单文件
		file,_ := c.FormFile("file")
		log.Println(file.Filename)
		// 上传至指定目录
		err := c.SaveUploadedFile(file, "/tmp/b.txt")
		if err != nil {
			fmt.Println(err)
		}
		c.String(http.StatusOK,fmt.Sprintf("'%s' uploaded!",file.Filename))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}