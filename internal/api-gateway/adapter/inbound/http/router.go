package http

import "github.com/gin-gonic/gin"

func NewRouter(auth *AuthHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/signup", auth.Signup)
		}
	}

	return r
}
