package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type RouteProvider interface {
	RegisterRoutes(engine *gin.Engine)
}

type ValidationProvider interface {
	RegisterValidation(injector *do.Injector)
}

type Router struct {
	g *gin.Engine
}

func NewRouter() *Router {
	gin.SetMode(gin.ReleaseMode)

	r := &Router{
		g: gin.New(),
	}

	r.middlewares()

	return r
}

func (r *Router) Handler() http.Handler {
	return r.g
}

func (r *Router) middlewares() {
	r.g.Use(
		gin.Recovery(),
		gin.ErrorLogger(),
	)
}

func (r *Router) RegisterApplicationRoutes(providers ...RouteProvider) {
	for _, provider := range providers {
		provider.RegisterRoutes(r.g)
	}
}
