package api

import (
	"backend/config"
	"backend/model"
	"backend/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register @Summary 注册新用户
// @Description 添加一个新用户到系统中
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param   username     body    string     true        "用户名"
// @Param   password     body    string     true        "密码"
// @Success 200 {object} map[string]interface{} "成功注册用户"
// @Router /register [post]
func Register(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 对密码进行哈希处理
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashedPassword)

	db := config.ConnectDatabase()
	result := db.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {
	var loginInfo struct {
		Username string
		Password string
	}
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 在数据库中查找用户
	db := config.ConnectDatabase()
	var user model.User
	if err := db.Where("username = ?", loginInfo.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// 生成JWT令牌
	token, _ := util.GenerateJWT(user.Username)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func ChangePassword(c *gin.Context) {
	var changePasswordInfo struct {
		Username    string
		OldPassword string
		NewPassword string
	}
	if err := c.ShouldBindJSON(&changePasswordInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 数据库操作省略，逻辑与登录类似，先验证旧密码，然后更新为新密码
}
