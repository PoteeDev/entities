package info

import (
	"net/http"

	"github.com/explabs/ad-ctf-paas-api/database"
	"github.com/explabs/ad-ctf-paas-api/models"
	"github.com/gin-gonic/gin"
)

type Team struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	SshPubKey string `json:"ssh_pub_key"`
}

func GetTeamInfo(c *gin.Context) {
	user, _ := c.Get("id")
	team, err := database.GetTeam(user.(*models.JWTTeam).TeamName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"team": team})
}
