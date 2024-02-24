package router

import (
	"backend/api"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	r.POST("/changePassword", api.ChangePassword)
}
