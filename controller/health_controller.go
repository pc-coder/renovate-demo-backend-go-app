package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthController some
type HealthController interface {
	GetStatus(c *gin.Context)
}

type healthController struct {
}

func (hc healthController) GetStatus(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}

// NewHealthController some
func NewHealthController() HealthController {
	return healthController{}
}