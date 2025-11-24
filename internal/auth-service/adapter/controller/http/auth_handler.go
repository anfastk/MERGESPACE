package http

import (
	"net/http"

	"github.com/anfastk/MERGESPACE/internal/auth-service/application/dto"
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/input"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/apperrors"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	useCase input.AuthUseCase
}

func NewAuthHandler(useCase input.AuthUseCase) *AuthHandler {
	return &AuthHandler{useCase: useCase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.useCase.Register(c.Request.Context(), req)
	if err != nil {
		if err == apperrors.ErrUserAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract Metadata
	req.UserAgent = c.Request.UserAgent()
	req.IPAddress = c.ClientIP()

	res, err := h.useCase.Login(c.Request.Context(), req)
	if err != nil {
		if err == apperrors.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
