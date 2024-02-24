package main

import (
	"backend/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	r.POST("/changePassword", api.ChangePassword)

	err := r.Run()

	if err != nil {
		return
	}
}
