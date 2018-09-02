package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/VulnaaS/VulnaaS-API/config"
	"github.com/VulnaaS/VulnaaS-API/script"
	"github.com/labstack/echo"
)

// HealthCheck is the heath check function.
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING!\n")
}

// ReceiveInstallRequest returns a script to do further enumeration of the box being used
// to install the vulnerability.
func ReceiveInstallRequest(c echo.Context) error {

	inputScript := c.Param("input")
	installScript, err := validateInput(inputScript)
	if err != nil {
		errMsg := fmt.Sprintf("echo \"[x][VulnaaS] %s not found\"", inputScript)
		return c.String(http.StatusOK, errMsg)
	}

	// Unix scripts = between 2000 and 2999.
	if (installScript.ID >= 2000) && (installScript.ID <= 2999) {
		UnixScript := replaceEnvVars(config.VulnaasConfig.CheckPackageManagerScript.UnixCmd, installScript.Alias)
		return c.String(http.StatusOK, UnixScript)
	}

	// Windows scripts = between 3000 and 3999.
	if (installScript.ID >= 3000) && (installScript.ID <= 3999) {
		WinScript := replaceEnvVars(config.VulnaasConfig.CheckPackageManagerScript.WindowsCmd, installScript.Alias)
		return c.String(http.StatusOK, WinScript)
	}

	// ServiceScript ID is hit but Vulnaas do not know the package manager to do the correct redirect.
	errMsg := fmt.Sprintf("echo \"[x][VulnaaS] %d is a ServiceScript. Use http://%s:%s/scripts/:pm/%d instead.\"", installScript.ID, config.APIhost, config.APIport, installScript.ID)
	return c.String(http.StatusOK, errMsg)
}

// InstallScript returns a script based on its ID and its package manager.
func InstallScript(c echo.Context) error {

	inputScript := c.Param("input")
	installScript, err := validateInput(inputScript)
	if err != nil {
		errMsg := fmt.Sprintf("echo \"[x][VulnaaS] %s not found\"", inputScript)
		return c.String(http.StatusOK, errMsg)
	}

	// check if pm is yum, apt or win.
	inputPacketManager := c.Param("pm")
	if err = checkPackageManager(inputPacketManager); err != nil {
		errMsg := fmt.Sprintf("echo \"[x][VulnaaS] Package Manager %s invalid.\"", inputPacketManager)
		return c.String(http.StatusOK, errMsg)
	}

	switch inputPacketManager {
	case "yum":
		return c.String(http.StatusOK, replaceEnvVars(installScript.CmdYum, "0"))
	case "apt":
		return c.String(http.StatusOK, replaceEnvVars(installScript.CmdApt, "0"))
	case "win":
		return c.String(http.StatusOK, replaceEnvVars(installScript.CmdWindows, "0"))
	default:
		return c.String(http.StatusOK, "echo \"[x][VulnaaS] Internal Error: inputPacketManager\"")
	}

}

// checkPackageManager returns an error if a string is equal yum or apt.
func checkPackageManager(pm string) error {
	if pm == "yum" || pm == "apt" || pm == "win" {
		return nil
	}
	err := errors.New("package manager invalid")
	return err
}

// replaceEnvVars will extract %API_HOST% and %API_PORT% from a string and replace it
// with the proper values from config package. Also replaces %SCRIPT% given an input.
func replaceEnvVars(receivedString string, input string) string {
	step1 := strings.Replace(receivedString, "%API_HOST%", config.APIhost, -1)
	step2 := strings.Replace(step1, "%API_PORT%", config.APIport, -1)
	if input != "0" {
		step3 := strings.Replace(step2, "%SCRIPT%", input, -1)
		return step3
	}
	return step2
}

// validateInput validates the input and returns corresponding InstallScript and
// error.
func validateInput(input string) (config.InstallScript, error) {

	var installScript config.InstallScript
	isID := false
	scriptFound := false

	ID, err := strconv.Atoi(input)
	if err == nil {
		isID = true
	}

	if isID {
		searchScript, err := script.GetByID(ID)
		if err == nil {
			installScript = searchScript
			scriptFound = true
		}
	} else {
		searchScript, err := script.GetByAlias(input)
		if err == nil {
			installScript = searchScript
			scriptFound = true
		}
	}

	if !scriptFound {
		errorMsg := fmt.Sprintf("input not valid")
		return installScript, errors.New(errorMsg)
	}

	return installScript, nil
}
