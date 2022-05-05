package vpn

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/PoteeDev/auth/auth"
	"github.com/gin-gonic/gin"
)

func DownloadVpnConfig(username string) (string, error) {
	urlAddr := vpnUrl + "api/user/config/show"
	response, httpErr := http.PostForm(urlAddr, url.Values{
		"username": {username},
	})
	if httpErr != nil {
		log.Println(httpErr)
		return "", httpErr
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseData), nil
}

func GetVpnConfigHandler(c *gin.Context) {
	metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	username := metadata.UserId
	vpnConfig, err := DownloadVpnConfig(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	c.Data(200, "plain/text; charset=utf-8", []byte(vpnConfig))

}
