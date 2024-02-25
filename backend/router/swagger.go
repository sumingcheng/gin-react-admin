package router

import (
	_ "backend/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// SwaggerRouter 设置路由
func SwaggerRouter(r *gin.Engine) {
	// 注册Swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
