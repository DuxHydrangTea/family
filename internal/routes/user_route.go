package routes

import (
	"family/internal/handlers"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	h *handlers.UserHandler
}

func NewUserRoute(h *handlers.UserHandler) *UserRoute {
	return &UserRoute{
		h: h,
	}
}

func (r *UserRoute) Register(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", r.h.GetAll)
	}
}
