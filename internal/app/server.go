package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"school/internal/config"
	"school/internal/database"
	"school/internal/routes"
)

type Server struct {
	engine *gin.Engine
	cfg    *config.Config
}

func NewServer() *Server {
	cfg := config.Load()
	db := database.Connect(cfg)
	r := gin.New()
	r.Use(gin.Recovery())

	// health
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })

	routes.Register(r, db, cfg)
	return &Server{engine: r, cfg: cfg}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	if addr == ":" {
		addr = ":8080"
	}
	return s.engine.Run(addr)
}
