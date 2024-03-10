package handler

import (
	"github.com/Fox1N69/rest-tsc/db"
	"github.com/Fox1N69/rest-tsc/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllClients(c *gin.Context) {
	var data []models.Clients
	db.DB.Find(&data)

	c.JSON(200, data)
}

func (h *Handler) getAllRequest(c *gin.Context) {

}

func (h *Handler) Test(c *gin.Context) {
}
