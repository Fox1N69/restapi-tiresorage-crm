package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouting() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/clients", h.getAllClients)
		api.GET("/clientreq", h.getAllRequest)
		api.GET("/test", h.Test)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	return router
}
