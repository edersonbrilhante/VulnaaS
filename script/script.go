package script

import (
	"errors"

	"github.com/VulnaaS/VulnaaS/config"
)

// GetByID returns an InstallScript given its ID.
func GetByID(ID int) (config.InstallScript, error) {
	var installScript config.InstallScript
	foundScript := false
	for _, searchScript := range config.VulnaasConfig.ServiceScripts {
		if searchScript.ID == ID {
			foundScript = true
			installScript = searchScript
			break
		}
	}
	if !foundScript {
		for _, searchScript := range config.VulnaasConfig.VulnaasScripts {
			if searchScript.ID == ID {
				foundScript = true
				installScript = searchScript
				break
			}
		}
	}
	if !foundScript {
		errMsg := errors.New("id not found")
		return installScript, errMsg
	}
	return installScript, nil
}

// GetByAlias returns an InstallScript given its alias.
func GetByAlias(alias string) (config.InstallScript, error) {
	var installScript config.InstallScript
	foundScript := false
	for _, searchScript := range config.VulnaasConfig.ServiceScripts {
		if searchScript.Alias == alias {
			foundScript = true
			installScript = searchScript
			break
		}
	}
	if !foundScript {
		for _, searchScript := range config.VulnaasConfig.VulnaasScripts {
			if searchScript.Alias == alias {
				foundScript = true
				installScript = searchScript
				break
			}
		}
	}
	if !foundScript {
		errMsg := errors.New("alias not found")
		return installScript, errMsg
	}
	return installScript, nil
}
