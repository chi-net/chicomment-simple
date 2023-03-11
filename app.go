package main

import (
  "net/http"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"

  "github.com/gin-gonic/gin"
)

type comments struct {
	ip string // 评论者IP地址
	content string // 评论内容
	email string // 邮箱
	to int // 评论对象 文章
	site string // 网站名称
	name string // 用户名
	reply int // 回复 若无则0
	time int // 评论时间
	id int // 评论id
	ua string // 评论ua
}

func main() {
  //读取文件
	// bytes, e := ioutil.ReadFile("commentdata.json")
	// if e != nil {
	// 	fmt.Println(e)
	// }
	// fmt.Println(string(bytes))
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
			"status": 200,
    })
  })
	r.POST("/", func(c *gin.Context) {
		u := comments{}
		u.email = c.PostForm("email")
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
    })
	})
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}