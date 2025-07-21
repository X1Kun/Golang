package main

import (
	"20250702/internal/handler"
	"20250702/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// go + gin + gorm + sqlite

func main() {
	// 利用gorm打开sqlite
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// gorm按照结构体结构，自动识别sqlite中的表格
	db.AutoMigrate(&repository.User{}, &repository.Todo{})

	// 为db包装一层仓库：这个todoRepo只用来对todo数据表进行操作，逻辑业务包装在里面
	todoRepo := repository.NewTodoRepository(db)
	// 为仓库包装一层handler：这个todoHandler只用于接收http的指令，并转化为todo_repo的操作
	todoHandler := handler.NewTodoHandler(todoRepo)
	// 为db封装一层user_repo的仓库方法，使其只对user相关操作
	userRepo := repository.NewUserRepository(db)
	// 为user_repo封装一层user_handler，使user_handler只接收有关user的http请求，并只关心http，不关心具体业务逻辑
	userHandler := handler.NewUserHandler(userRepo)

	// 初始化gin
	r := gin.Default()
	// 静态内容，用于测试
	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	// 将具体业务分发给todoHandler
	// 展示全部清单
	r.GET("/todos", todoHandler.GetAllTodos)
	// 增加一项清单
	r.POST("/todos", todoHandler.CreateTodo)
	// 根据id展示清单
	r.GET("/todos/:id", todoHandler.GetTodoByID)
	// 根据id更新清单
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	// 根据id删除清单
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	// 将用户相关业务分发给userHandler
	// 用户注册
	r.POST("/register", userHandler.Register)
	// 用户登录
	r.POST("/login", userHandler.Login)

	// 启动服务器
	r.Run()
}
