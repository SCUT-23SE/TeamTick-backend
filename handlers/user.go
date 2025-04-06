package handlers

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetUsersMe 获取当前用户信息
// (GET /users/me)
func (h *UserHandler) GetUsersMe(c *gin.Context) {
}
