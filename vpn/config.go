package vpn

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// Function for download vpn config from vpn service.
// In future release this function will be changed to request to S3 storage
// Now it returns vpn config in plain text and error
func (c *VpnClient) DownloadVpnConfig() (string, error) {
	urlAddr := c.VpnService + "api/user/config/show"
	response, httpErr := http.PostForm(urlAddr, url.Values{
		"username": {c.Login},
	})
	if httpErr != nil {
		return "", httpErr
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(responseData), nil
}
