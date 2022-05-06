package vpn

import (
	"net/http"
	"net/url"
)

// Create VPN config for VpnClient
// This function send http request to ovpn service to register new user
// It returns nill or error from vpn service if user already exists or connection error
func (c *VpnClient) CreateConfig() error {
	urlAddr := c.VpnService + "api/user/create"
	_, httpErr := http.PostForm(urlAddr, url.Values{
		"username": {c.Login},
		"password": {c.Password},
	})
	if httpErr != nil {
		return httpErr
	}
	return nil
}
