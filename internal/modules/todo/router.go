package todo

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	mutation *Mutation
}

func NewRouter(m *Mutation) *Router {
	return &Router{mutation: m}
}

func (r *Router) Register(rg *gin.RouterGroup) {
	group := rg.Group("/todos")
	{
		group.POST("", r.mutation.Create)
	}
}
