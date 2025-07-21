package handler

import (
	"20250702/internal/repository" // <-- ！！！替换成你的项目模块名
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// UserHandler 负责处理用户相关的HTTP请求
type UserHandler struct {
	Repo *repository.UserRepository
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}

// NewUserHandler 是UserHandler的构造函数
func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

// Register 处理用户注册的请求
func (h *UserHandler) Register(c *gin.Context) {
	// 1. 使用我们新的请求结构体来接收和验证数据
	var req RegisterRequest

	// 2. 将请求体中的JSON绑定到req上。
	// Gin 会自动使用 validator 对带有 'binding' 标签的字段进行验证！
	if err := c.ShouldBindJSON(&req); err != nil {
		// 如果验证失败，err 会包含详细的错误信息
		// gin.H{"error": err.Error()} 会返回一个更清晰的错误给前端
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. 验证通过后，将请求数据映射到数据库模型
	newUser := repository.User{
		Username: req.Username,
		Password: req.Password,
	}

	// 4. 调用Repository层来创建用户 (这部分逻辑不变)
	err := h.Repo.CreateUser(&newUser)
	if err != nil {
		// 这里可能因为用户名已存在或其他数据库问题而出错
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	// 5. 为了安全，不应该在响应中返回密码。我们创建一个新的结构体来响应
	// 或者直接修改newUser对象的Password字段为空
	newUser.Password = ""

	// 6. 返回201 Created，以及新创建的用户信息（不含密码）
	c.JSON(201, newUser)
}

// Login 处理用户登录请求
func (h *UserHandler) Login(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 1. 绑定JSON到loginReq
	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 2. 根据用户名从数据库中查找用户
	user, err := h.Repo.GetUserByUsername(loginReq.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 为了安全，不明确提示是“用户不存在”还是“密码错误”，统一返回“无效”
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// 3. 比较密码 (!!!重要!!!: 这是非常不安全的做法，我们稍后会用bcrypt替换)
	if user.Password != loginReq.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// --- 登录成功，开始生成JWT ---

	// 4. 创建JWT的载荷 (Claims)，包含自定义信息
	claims := jwt.MapClaims{
		"userID":   user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 过期时间，这里设置为72小时
		"iat":      time.Now().Unix(),                     // 签发时间
	}

	// 5. 使用指定的签名方法创建 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 6. 使用一个密钥 (secret key) 来签名 token，得到最终的token字符串
	// !!!重要!!!: 这个密钥在真实项目中应该非常复杂，并从环境变量中读取
	mySecretKey := []byte("your_very_secret_key")
	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 7. 登录成功，返回token
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
