package registration

import (
	"time"

	"github.com/explabs/ad-ctf-paas-api/database"
	"github.com/explabs/ad-ctf-paas-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	SshPubKey string `json:"ssh_pub_key"`
}

// func generateIp(number int) string {
// 	ip, _, err := net.ParseCIDR(config.Conf.Network)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	ip = ip.To4()
// 	ip[2] = byte(number)
// 	ip[3] = 10
// 	// TODO: find better solution for generate cidr
// 	return ip.String() + "/24"
// }

func (t *Team) WriteTeam(login string) error {

	// ipAddress := generateIp(len(teams) + 1)
	hash, hashErr := HashPassword(t.Password)
	if hashErr != nil {
		return hashErr
	}
	dbTeam := &models.Team{
		ID:    primitive.NewObjectID(),
		Name:  t.Name,
		Login: login,
		// Address:   ipAddress,
		Hash:      hash,
		SshPubKey: t.SshPubKey,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return database.CreateTeam(dbTeam)
}
