package api

import (
	"backend/initialize"
	"backend/model"
	"backend/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register
// @Summary 注册新用户
// @Description 用户注册。
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param   username     body    string     true        "用户名"
// @Param   password     body    string     true        "密码"
// @Param   email        body    string     true        "邮箱"
// @Success 200 {object} map[string]interface{} "message:注册成功"
// @Failure 400 {object} map[string]interface{} "error:错误信息"
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

	db, _ := initialize.GetDB()
	result := db.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Login
// @Summary 用户登录
// @Description 用户登录，成功返回 JWT 令牌。
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param   username     body    string     true        "用户名"
// @Param   password     body    string     true        "密码"
// @Success 200 {object} map[string]interface{} "token:JWT令牌"
// @Failure 401 {object} map[string]interface{} "error:认证失败的错误信息"
// @Router /login [post]
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
	db, _ := initialize.GetDB()
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

// ChangePassword
// @Summary 修改密码
// @Description 修改用户的密码。
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param   username     body    string     true        "用户名"
// @Param   oldPassword  body    string     true        "旧密码"
// @Param   newPassword  body    string     true        "新密码"
// @Success 200 {object} map[string]interface{} "message:密码修改成功"
// @Failure 400 {object} map[string]interface{} "error:错误信息"
// @Router /changePassword [post]
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
