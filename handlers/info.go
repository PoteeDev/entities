package handlers

import (
	"net/http"

	"github.com/PoteeDev/auth/auth"
	"github.com/PoteeDev/entities/database"
	"github.com/gin-gonic/gin"
)

func GetEntityInfo(c *gin.Context) {
	metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"detail": "unauthorized"})
		return
	}
	team, err := database.GetEntity(metadata.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"team": team})
}
