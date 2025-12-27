package http

import (
	"net/http"

	authpb "github.com/anfastk/MERGESPACE/api/proto/v1"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authClient authpb.AuthServiceClient
}

func NewAuthHandler(client authpb.AuthServiceClient) *AuthHandler {
	return &AuthHandler{authClient: client}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var req struct {
		Email     string `json:"email"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Password  string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authClient.InitiateSignup(
		c.Request.Context(),
		&authpb.SignUpRequest{
			Email:     req.Email,
			Firstname: req.FirstName,
			Lastname:  req.LastName,
			Password:  req.Password,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
