package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PoteeDev/entities/database"
	"github.com/PoteeDev/entities/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
)

type Entity struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (t *Entity) WriteEntity(login string) error {

	// ipAddress := generateIp(len(teams) + 1)
	hash, hashErr := HashPassword(t.Password)
	if hashErr != nil {
		return hashErr
	}
	dbEntity := &models.Entity{
		Name:  t.Name,
		Login: login,
		Hash:  hash,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return database.AddEntity(dbEntity)
}

func CreateEntity(c *gin.Context) {
	// status, err := database.RegistrationStatus()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"detail": err})
	// 	return
	// }
	// if status == "close" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"detail": "registration closed"})
	// 	return
	// }
	var entity Entity
	jsonErr := c.BindJSON(&entity)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": jsonErr.Error()})
		return
	}
	// create slug name for entity
	login := slug.Make(entity.Name)
	// todo: change name policy in vpn service and remove this line
	login = strings.Replace(login, "-", "_", -1)

	// check if user already exists
	// todo: create function in mongo to check exists usern, not find
	dbEntity, err := database.GetEntity(login)
	fmt.Println(dbEntity, err)
	if dbEntity != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "entity already exists"})
		return
	}
	// write user to database
	if writeErr := entity.WriteEntity(login); writeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": writeErr.Error()})
	}

	// generate vpn config
	// vpnErr := vpn.CreateVpnCLient(login, entity.Password).CreateConfig()
	// if vpnErr != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"detail": vpnErr.Error()})
	// 	return
	// }
	// if all ok - return message and entity login
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("The entity %s created", entity.Name),
		"login":   login,
	})
}
