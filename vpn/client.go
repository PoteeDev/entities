package vpn

import (
	"fmt"
	"os"
)

// Vpn CLient struct for comunicate with vpn service
type VpnClient struct {
	VpnService string
	Login      string
	Password   string
}

// Generator function for initializate Vpn Client
// It returns Vpn client Object
func CreateVpnCLient(login string, args ...string) *VpnClient {
	var vpnUrl = fmt.Sprintf("http://%s:%s/", os.Getenv("OVPN_HOST"), os.Getenv("OVPN_PORT"))
	password := ""

	if len(args) > 0 {
		password = args[0]
	}
	return &VpnClient{
		VpnService: vpnUrl,
		Login:      login,
		Password:   password,
	}
}
