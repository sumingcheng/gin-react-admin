package main

import (
	"backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 用户相关路由
	router.UserRoutes(r)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
