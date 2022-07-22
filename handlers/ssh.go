package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SSHStruct struct {
	Key string `json:"key"`
}

func UploadSShKey(c *gin.Context) {
	var s SSHStruct
	err := c.BindJSON(&s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Key)
	// TODO: upload ssh key to s3 storage
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
