package main

import (
	"backend/logger"
	"backend/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	l := logger.New()

	// 创建路由
	r := gin.Default()
	// 注册路由
	router.UserRoutes(r)
	router.SwaggerRouter(r)

	// 设置信任的代理服务器列表
	l.Info("Starting the application...")
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// 启动服务
	if err := r.Run(":33333"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	l.Info("Application started successfully")
}
