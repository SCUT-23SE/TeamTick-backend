package handlers

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// PostAuthLogin 用户登录
// (POST /auth/login)
func (h *AuthHandler) PostAuthLogin(c *gin.Context) {
}

// PostAuthRegister 用户注册
// (POST /auth/register)
func (h *AuthHandler) PostAuthRegister(c *gin.Context) {
}
