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

func (h *UserHandler) UploadAvatar(c *gin.Context) {
	file, header, err := c.Request.FormFile("avatar")
	// Validate file ...
	// ...code...
	// End validate file
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}
	defer file.Close()

	fileName := header.Filename

	fullPath, err := h.s.UploadAvatar(file, fileName)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"fullPath": fullPath,
	})
}
