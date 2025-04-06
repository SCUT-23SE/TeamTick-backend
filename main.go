package main

import (
	"TeamTickBackend/gen"
	"TeamTickBackend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化各个 handler
	authHandler := handlers.NewAuthHandler()
	groupHandler := handlers.NewGroupHandler()
	userHandler := handlers.NewUserHandler()

	// 注册路由
	gen.RegisterAuthHandlers(r, authHandler)
	gen.RegisterGroupsHandlers(r, groupHandler)
	gen.RegisterUsersHandlers(r, userHandler)

	// 启动服务器
	r.Run(":8080")
}
