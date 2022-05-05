package vpn

import (
	"log"
	"net/http"
	"net/url"

	"github.com/explabs/ad-ctf-paas-api/models"
)

var vpnUrl = "http://openvpn:9000/"

func AddVpnTeam(team *models.Team, rawPassword string) error {
	urlAddr := vpnUrl + "api/user/create"
	_, httpErr := http.PostForm(urlAddr, url.Values{
		"username": {team.Login},
		"password": {rawPassword},
	})
	if httpErr != nil {
		log.Println("here")
		log.Println(httpErr)
		return httpErr
	}
	return nil
}
