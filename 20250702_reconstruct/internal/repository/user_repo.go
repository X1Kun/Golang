package repository

import (
	"gorm.io/gorm"
)

// UserRepository 负责用户数据的存储操作
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository 是UserRepository的构造函数
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser 在数据库中创建一个新用户
func (r *UserRepository) CreateUser(user *User) error {
	// GORM的Create方法会处理插入操作
	// 如果用户名重复，由于我们在模型中定义了unique，数据库会返回错误
	err := r.DB.Create(user).Error
	return err
}

// GetUserByUsername 根据用户名查找用户
func (r *UserRepository) GetUserByUsername(username string) (*User, error) {
	var user User
	// 使用 Where 查询条件来查找第一条匹配的记录
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
