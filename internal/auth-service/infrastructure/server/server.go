package server

/* import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	port   string
}

func NewServer(authHandler *http.AuthHandler, port string) *Server {
	router := gin.Default()

	authHandler.SignupRoutes(router)

	return &Server{
		engine: router,
		port:   port,
	}
}

func (s *Server) Run() error {
	return s.engine.Run(":" + s.port)
}
 */