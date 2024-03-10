package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouting() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	return router
}
