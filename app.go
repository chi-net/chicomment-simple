package main

import (
  "net/http"
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"strconv"

  "github.com/gin-gonic/gin"
)

type comments struct {
	IP string // 评论者IP地址
	Content string // 评论内容
	Email string // 邮箱
	To int // 评论对象 文章
	Site string // 网站名称
	Name string // 用户名
	Reply int // 回复 若无则0
	Time int64 // 评论时间
	Id int // 评论id
	UA string // 评论ua
}

func main() {
  //读取文件
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
		bytes, e := ioutil.ReadFile("commentdata.json")
		if e != nil {
			// ioutil.WriteFile("commentdata.json", []byte(""), 666)
			c.JSON(http.StatusOK, e)
		}
    c.JSON(http.StatusOK, string(bytes))
  })
	r.POST("/", func(c *gin.Context) {
		u := comments{}
		u.Email = c.PostForm("email")
		u.Content = c.PostForm("content")
		u.To,_ = strconv.Atoi(c.PostForm("to"))
		u.Site = c.PostForm("site")
		u.Name = c.PostForm("name")
		u.Reply,_ = strconv.Atoi(c.PostForm("reply"))
		u.IP = c.ClientIP()
		u.UA = c.Request.Header["User-Agent"][0]
		u.Time = time.Now().Unix()
		// 获取id
		data, e := ioutil.ReadFile("countdata.txt")
		if e != nil {
			ioutil.WriteFile("countdata.txt", []byte("0"), 666)
		}
		id,_ := strconv.Atoi(string(data))
		id += 1
		ioutil.WriteFile("countdata.txt", []byte(strconv.Itoa(id)), 666)
		u.Id = id
		// 获取文件
		fmt.Println(u) // 打印记录
		var commentdata = []comments{}
		// 写入commentsdata
		bytes, e := ioutil.ReadFile("commentdata.json")
		if e != nil {
			commentdata = append(commentdata, u)
			jsondata,_ := json.Marshal(commentdata)
			fmt.Println(jsondata)
			ioutil.WriteFile("commentdata.json", jsondata, 666)
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
			c.JSON(http.StatusOK, e)
		}
		json.Unmarshal([]byte(bytes), &commentdata)
		commentdata = append(commentdata, u)
		jsondata,_ := json.Marshal(commentdata)
		fmt.Println(jsondata)
		ioutil.WriteFile("commentdata.json", jsondata, 666)
		// fmt.Println(commentdata)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
    })
	})
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}