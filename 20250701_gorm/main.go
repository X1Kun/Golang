package main

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// gorm的结构体不需要加ID和时间，会自动创建
type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Status string `json:"status"`
}

// 全局的数据库连接对象
var db *gorm.DB

func main() {
	// 连接数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动迁移
	db.AutoMigrate(&Todo{})
	// 创建Gin引擎
	r := gin.Default()

	// 测试页面：localhost:8080/ping
	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })
	// 清单列表 localhost:8080/todos
	r.GET("/todos", func(c *gin.Context) {

		var todos []Todo
		// 去数据库里查找所有`todos`表的记录，把结果填充到todos切片内
		// result返回是否错误
		result := db.Find(&todos)
		// 如果有错，返回一个服务器内部错误
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve todos"})
			return
		}
		// OK，返回gorm在数据库的Find结果
		c.JSON(200, todos)
	})

	// 创建一项
	r.POST("/todos", func(c *gin.Context) {
		var newTodo Todo
		// POST内容 -> 结构体
		if err := c.BindJSON(&newTodo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON format"})
			return
		}

		// gorm将db.Create()转化为SQL语句，将&newTodo传入数据库
		result := db.Create(&newTodo)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Failed to create todo"})
			return
		}

		// 若创建成功，GORM会自动把数据库生成的ID等信息回填到newTodo对象中，可直接返回它
		c.JSON(201, newTodo)
	})

	// 展示一项
	r.GET("/todos/:id", func(c *gin.Context) {
		// 获取并转换ID为int
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}

		// 接收查询结果
		var todo Todo
		// db.First()会根据主键(id)去查找数据库，若找到，把结果填充到todo变量里。若找不到，它会返回gorm.ErrRecordNotFound错误
		result := db.First(&todo, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(404, gin.H{"error": "Todo not found"})
			} else {
				// 数据库本身错误
				c.JSON(500, gin.H{"error": "Database error"})
			}
			return
		}

		// 返回找到的todo
		c.JSON(200, todo)
	})

	// 更新一项
	r.PUT("/todos/:id", func(c *gin.Context) {
		// 获取并转换ID
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}

		// 先按ID查找记录，并返回找到的记录
		var todo Todo
		if err := db.First(&todo, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}

		// 将客户端发来的JSON数据覆盖到todo上
		if err := c.BindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON format"})
			return
		}

		// 使用db.Save()根据主键ID更新数据库
		result := db.Save(&todo)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Failed to update todo"})
			return
		}

		// 返回更新后的对象
		c.JSON(200, todo)
	})

	// 删除一项
	r.DELETE("/todos/:id", func(c *gin.Context) {
		// 获取并转换ID
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}

		// 使用db.Delete()根据主键删除记录
		result := db.Delete(&Todo{}, id)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Failed to delete todo"})
			return
		}
		// 检查是否真的有记录被删除了
		if result.RowsAffected == 0 {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}
		// 返回204 No Content
		c.Status(204)
	})

	r.Run()
}
