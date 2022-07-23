package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PoteeDev/team/database"
	"github.com/PoteeDev/team/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Team struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (t *Team) WriteTeam(login string) error {

	// ipAddress := generateIp(len(teams) + 1)
	hash, hashErr := HashPassword(t.Password)
	if hashErr != nil {
		return hashErr
	}
	dbTeam := &models.Team{
		ID:        primitive.NewObjectID(),
		Name:      t.Name,
		Login:     login,
		Blocked:   false,
		Visible:   true,
		Hash:      hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return database.AddTeam(dbTeam)
}

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
	var team Team
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
