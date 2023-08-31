package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

// server serves http requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP Server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
