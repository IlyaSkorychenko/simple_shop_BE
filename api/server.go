package api

import (
	"github.com/IlyaSkorychenko/simple_shop_BE/pkg"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() Server {
	srv := Server{
		router: gin.Default(),
	}
	srv.registerMiddlewares()
	srv.registerRouts()

	return srv
}

func (s Server) Start(port string) {
	err := s.router.Run(port)
	pkg.Check(err)
}

func (s Server) registerRouts() {
	rg := s.router.Group("/v1")

	s.addProductRouts(rg)
}

func (s Server) registerMiddlewares() {
	s.router.Use(errHandler)
}
