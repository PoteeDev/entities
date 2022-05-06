package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PoteeDev/team/registration"
	"github.com/explabs/ad-ctf-paas-api/database"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func CreateTeam(c *gin.Context) {
	// status, err := database.RegistrationStatus()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"detail": err})
	// 	return
	// }
	// if status == "close" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"detail": "registration closed"})
	// 	return
	// }
	var team registration.Team
	jsonErr := c.BindJSON(&team)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": jsonErr.Error()})
		return
	}
	// create slug name for team
	login := slug.Make(team.Name)
	// todo: change name policy in vpn service and remove this line
	login = strings.Replace(login, "-", "_", -1)

	// check if user already exists
	// todo: create function in mongo to check exists usern, not find
	dbTeam, err := database.GetTeam(login)
	fmt.Println(dbTeam, err)
	if dbTeam != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "team already exists"})
		return
	}
	// write user to database
	if writeErr := team.WriteTeam(login); writeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": writeErr.Error()})
	}

	// generate vpn config
	// vpnErr := vpn.CreateVpnCLient(login, team.Password).CreateConfig()
	// if vpnErr != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"detail": vpnErr.Error()})
	// 	return
	// }
	// if all ok - return message and team login
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("The team %s created", team.Name),
		"login":   login,
	})
}
