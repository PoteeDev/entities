package vpn

// Vpn CLient struct for comunicate with vpn service
type VpnClient struct {
	VpnService string
	Login      string
	Password   string
}

// Generator function for initializate Vpn Client
// It returns Vpn client Object
func CreateVpnCLient(login string, args ...string) *VpnClient {
	var vpnUrl = "http://openvpn:9000/"
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
