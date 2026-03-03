package routes

import "github.com/gin-gonic/gin"

type Route interface {
	Register(routes *gin.Engine)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	for _, route := range routes {
		route.Register(r)
	}
}
