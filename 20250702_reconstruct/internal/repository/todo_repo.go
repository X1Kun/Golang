package repository

import (
	"gorm.io/gorm"
)

// User 模型
type User struct {
	gorm.Model
	Username string `gorm:"unique"` // 用户名必须是唯一的
	Password string // 密码（我们先用明文，后面再讲加密）
	Todos    []Todo `gorm:"foreignKey:UserID"` // 这是关键！定义 "Has Many" 关系
}

// Todo 结构体的定义现在属于Repository层，因为它是数据的模型
type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Status string `json:"status"`
	UserID uint   `json:"userID"` // 这是关键！定义 "Belongs To" 关系
}

// TodoRepository 是我们的数据仓库，它包含一个数据库连接
type TodoRepository struct {
	DB *gorm.DB
}

// NewTodoRepository 是一个构造函数，方便我们创建Repository实例
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

// GetAll 从数据库获取所有todos
func (r *TodoRepository) GetAll() ([]Todo, error) {
	var todos []Todo
	err := r.DB.Find(&todos).Error
	return todos, err
}

// Create 在数据库中创建一个新的todo
func (r *TodoRepository) Create(todo *Todo) error {
	return r.DB.Create(todo).Error
}

// GetByID 从数据库按ID获取一个todo
func (r *TodoRepository) GetByID(id uint) (*Todo, error) {
	var todo Todo
	err := r.DB.First(&todo, id).Error
	return &todo, err
}

// Update 在数据库中更新一个todo
func (r *TodoRepository) Update(todo *Todo) error {
	return r.DB.Save(todo).Error
}

// Delete 从数据库删除一个todo
func (r *TodoRepository) Delete(id uint) (int64, error) {
	result := r.DB.Delete(&Todo{}, id)
	return result.RowsAffected, result.Error
}
