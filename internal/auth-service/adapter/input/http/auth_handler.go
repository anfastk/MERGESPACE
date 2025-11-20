package http

import (
	"net/http"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    auth *service.AuthService
}

func NewAuthHandler(a *service.AuthService) *AuthHandler {
    return &AuthHandler{auth: a}
}

func (h *AuthHandler) Signup(c *gin.Context) {
    var input dto.SignupInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    out, err := h.auth.Signup(input)
    if err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, out)
}

func (h *AuthHandler) Login(c *gin.Context) {
    var input dto.LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    out, err := h.auth.Login(input)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, out)
}
