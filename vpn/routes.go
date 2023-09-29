package vpn

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
)

// Function for download vpn config from vpn service.
// In future release this function will be changed to request to S3 storage
// Now it returns vpn config in plain text and error
func (c *VpnClient) AddRoute(subnetAddress string) error {
	urlAddr := c.VpnService + "api/user/ccd/apply"

	ip, ipnet, err := net.ParseCIDR(subnetAddress)
	if err != nil {
		return err
	}

	var jsonData = []byte(fmt.Sprintf(`{
		"User":"%s",
		"ClientAddress":"dynamic",
		"CustomRoutes":[
			{"Address":"%s","Mask":"%s"}
			]
		}`, c.Login, ip, net.IP(ipnet.Mask).String()))

	request, _ := http.NewRequest("POST", urlAddr, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.Status != "200 OK" {
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf(string(responseData))
	}
	return nil
}
