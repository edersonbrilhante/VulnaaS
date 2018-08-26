package scripts

import (
	"errors"

	"github.com/VulnaaS/VulnaaS-API/types"
)

var (
	// VulnaasConfig represents all vulnaas scripts.
	VulnaasConfig = []types.ConfigScript{}
	// ServiceConfig represents all service scripts.
	ServiceConfig = []types.ConfigScript{}
)

func init() {
	vulnaas1001 := types.ConfigScript{
		ID:     1001,
		Tittle: "Test",
		CmdYum: "echo \"Wow, its a Yum!\"",
		CmdApt: "echo \"Wow, its a Apt!\"",
	}

	VulnaasConfig = append(VulnaasConfig, vulnaas1001)
}

// Get returns a configurationScrit struct based on it id.
func Get(id int) (types.ConfigScript, error) {
	for _, config := range ServiceConfig {
		if config.ID == id {
			return config, nil
		}
	}
	for _, config := range VulnaasConfig {
		if config.ID == id {
			return config, nil
		}
	}
	nilConfig := types.ConfigScript{}
	err := errors.New("id not found")
	return nilConfig, err
}
