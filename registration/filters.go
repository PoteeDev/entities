package registration

import "fmt"

type CredPolicy struct {
	NameMaxLen     int
	NameMinLen     int
	PasswordMinLen int
	PasswordMaxLen int
}

func (c *CredPolicy) SetPolicy() {
	c.PasswordMinLen, c.PasswordMaxLen = 14, 40
	c.NameMinLen, c.NameMaxLen = 3, 40
}

func (c *CredPolicy) CheckCredPolicy(team *Team) error {
	if c.NameMinLen > len(team.Name) {
		return fmt.Errorf("short name: %s", team.Name)
	}

	return nil
}
