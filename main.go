package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Post 定义帖子的结构体
type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	// 连接SQLite数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移数据库
	db.AutoMigrate(&Post{})

	// 创建一个gin引擎
	r := gin.Default()

	// 发帖接口
	r.POST("/posts", func(c *gin.Context) {
		// 从请求中读取标题和正文
		title := c.PostForm("title")
		body := c.PostForm("body")

		// 创建一个帖子对象
		post := Post{Title: title, Body: body}

		// 将帖子对象保存到数据库中
		db.Create(&post)

		// 返回帖子对象的JSON表示
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
	})

	// 查看帖子列表接口
	r.GET("/posts", func(c *gin.Context) {
		// 从数据库中获取所有帖子
		var posts []Post
		db.Find(&posts)

		// 返回所有帖子的JSON表示
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": posts})
	})

	// 查看单个帖子接口
	r.GET("/posts/:id", func(c *gin.Context) {
		// 从URL参数中读取帖子的ID
		id := c.Param("id")

		// 从数据库中获取指定ID的帖子
		var post Post
		db.First(&post, id)

		// 返回指定ID的帖子的JSON表示
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
	})

	// 启动HTTP服务器
	r.Run(":8080")
}
