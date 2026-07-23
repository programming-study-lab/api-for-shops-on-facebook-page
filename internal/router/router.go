package router

import (
	v2 "api-for-shops-on-facebook-page/internal/router/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine       *gin.Engine
	dependencies *Dependencies
}

func NewRouer(engine *gin.Engine, dependencies *Dependencies) *Router {
	return &Router{
		engine:       engine,
		dependencies: dependencies,
	}
}

func (r *Router) Setup() *gin.Engine {
	groupV1 := r.engine.Group("/api/v2")
	{
		groupV1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"status":  true,
					"message": "success",
					"time":    time.Now(),
				},
			)
		})
		v2.FeedFacebookRouter(groupV1, r.dependencies.Feed)
		v2.CommentFacebookRouter(groupV1, r.dependencies.Comment)
	}

	return r.engine
}
