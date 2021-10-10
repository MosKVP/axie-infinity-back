package api

import (
	"breeding/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Calculate(c *gin.Context) {
	var req model.CalculateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"req": req})
	
}
