package handler

import (
	"20250702/internal/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TodoHandler 是我们的处理器，它依赖于Repository
type TodoHandler struct {
	Repo *repository.TodoRepository
}

// 为创建Todo请求创建一个专门的结构体
type CreateTodoRequest struct {
	Title  string `json:"title" binding:"required,max=100"`
	Status string `json:"status" binding:"omitempty"` // status是可选的
}

// NewTodoHandler 是一个构造函数
// 传入repo，获得handler
func NewTodoHandler(repo *repository.TodoRepository) *TodoHandler {
	return &TodoHandler{Repo: repo}
}

// GetAllTodos 处理获取所有todos的请求
// 逻辑：数据库没问题->执行
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve todos"})
		return
	}
	c.JSON(200, todos)
}

// CreateTodo 处理创建新todo的请求
// 逻辑：检查JSON格式->检查数据库（id重复/其他问题）->执行
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	// 同样，绑定并验证
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 注意：这里的 UserID 需要从认证中间件中获取，你以后会加上
	// userID, _ := c.Get("userID")

	// 将请求数据映射到数据库模型
	newTodo := repository.Todo{
		Title:  req.Title,
		Status: req.Status, // 如果请求中没提供Status，这里会是空字符串
		// UserID: userID.(uint),
	}

	if err := h.Repo.Create(&newTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}
	c.JSON(http.StatusCreated, newTodo)
}

// GetTodoByID 处理按ID获取todo的请求
// 逻辑：
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32) // 使用ParseUint更安全

	todo, err := h.Repo.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(500, gin.H{"error": "Database error"})
		}
		return
	}
	c.JSON(200, todo)
}

// UpdateTodo 处理更新todo的请求
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	todo, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.BindJSON(todo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := h.Repo.Update(todo); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update todo"})
		return
	}
	c.JSON(200, todo)
}

// DeleteTodo 处理删除todo的请求
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	rowsAffected, err := h.Repo.Delete(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete todo"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}
	c.Status(204)
}
