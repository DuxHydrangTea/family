package handlers

import (
	"family/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	s *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{
		s: s,
	}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users := h.s.GetAll()
	c.JSON(200, users)
}
