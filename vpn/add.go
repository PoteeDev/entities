package vpn

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Create VPN config for VpnClient
// This function send http request to ovpn service to register new user
// It returns nill or error from vpn service if user already exists or connection error
func (c *VpnClient) CreateConfig() error {
	urlAddr := c.VpnService + "api/user/create"
	response, httpErr := http.PostForm(urlAddr, url.Values{
		"username": {c.Login},
	})
	if httpErr != nil {
		return httpErr
	}
	if response.Status != "200 OK" {
		b, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf(strings.TrimSuffix(string(b), "\n\n"))

	}
	return nil
}
