package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []Todo{
	{ID: 1, Title: "学习 Go 语言", Status: "completed"},
	{ID: 2, Title: "学习 Gin 框架", Status: "in-progress"},
	{ID: 3, Title: "做一个超棒的项目", Status: "todo"},
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/todos", func(c *gin.Context) {
		c.JSON(200, todos)
	})

	r.POST("/todos", func(c *gin.Context) {
		var newTodo Todo
		if err := c.BindJSON(&newTodo); err != nil {
			return
		}
		todos = append(todos, newTodo)
		c.JSON(201, newTodo)
	})

	// 新增的路由
	r.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		// 更好的做法是使用 strconv
		for _, todo := range todos {
			todoID, _ := strconv.Atoi(id)
			if todo.ID == todoID {
				c.JSON(200, todo)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Todo not found"})
	})

	// 7. 新增一个路由，用来更新一个已存在的待办事项
	r.PUT("/todos/:id", func(c *gin.Context) {
		// a. 获取路径参数ID并转换为整数
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}
		// b. 从请求体中绑定更新后的todo数据
		var updatedTodo Todo
		if err := c.BindJSON(&updatedTodo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON format"})
			return
		}
		// c. 在todos切片中查找并替换
		found := false
		for i, todo := range todos {
			if todo.ID == id {
				// 找到了！用新的数据替换掉旧的
				todos[i] = updatedTodo
				found = true
				break // 找到就退出循环
			}
		}
		// d. 如果没找到，返回 404
		if !found {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}
		// e. 返回 200 OK 和更新后的对象
		c.JSON(200, updatedTodo)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID format"})
			return
		}
		index := -1
		for i, todo := range todos {
			if todo.ID == id {
				index = i
				break
			}
		}
		if index == -1 {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}
		todos = append(todos[:index], todos[index+1:]...)
		c.Status(204)
	})

	r.Run()
}
