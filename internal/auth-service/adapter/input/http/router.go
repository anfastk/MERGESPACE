package http

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine, h *AuthHandler) {
    group := r.Group("/auth")
    {
        group.POST("/signup", h.Signup)
        group.POST("/login", h.Login)
    }
}
