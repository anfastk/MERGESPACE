package http

import "github.com/gin-gonic/gin"

func (h *AuthHandler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1/auth")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
	}
}
