package handlers

import (
	"net/http"

	"github.com/PoteeDev/auth/auth"
	"github.com/PoteeDev/team/database"
	"github.com/gin-gonic/gin"
)

func GetTeamInfo(c *gin.Context) {
	metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"detail": "unauthorized"})
		return
	}
	team, err := database.GetTeam(metadata.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"team": team})
}
