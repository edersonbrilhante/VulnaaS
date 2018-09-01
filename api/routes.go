package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/VulnaaS/VulnaaS-API/config"
	"github.com/labstack/echo"
)

// HealthCheck is the heath check function.
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING!\n")
}

// ReceiveInstallRequest returns a script to do further enumeration of the box being used
// to install the vulnerability.
func ReceiveInstallRequest(c echo.Context) error {

	// check if ID is an integer.
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// StatusOK so vagrant up do not crash and send this echo message
		errMsg := fmt.Sprintf("echo \"[x][VulnaaS] Error: %s is not valid script.\"", c.Param("id"))
		return c.String(http.StatusOK, errMsg)
	}

	UnixScript := replaceEnvVars(config.VulnaasConfig.CheckPackageManagerScript.UnixCmd, c.Param("id"))
	WinScript := replaceEnvVars(config.VulnaasConfig.CheckPackageManagerScript.WindowsCmd, c.Param("id"))

	// ID > 1000 represents Unix vulnerabilities scripts.
	if ID > 1000 {
		return c.String(http.StatusOK, UnixScript)
	}

	// return Windows script.
	return c.String(http.StatusOK, WinScript)

}

// InstallScript returns a script based on its ID and its package manager.
func InstallScript(c echo.Context) error {

	// double check if ID is an integer.
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMsg := fmt.Sprintf("echo \"[x][VulnaaS] Error: %s is not valid script.\"", c.Param("id"))
		return c.String(http.StatusOK, errMsg)
	}

	// check if pm is yum, apt or win.
	packageManager := c.Param("pm")
	if err = checkPackageManager(packageManager); err != nil {
		errMsg := fmt.Sprintf("echo \"[x][VulnaaS] Error: Package Manager %s invalid.\"", packageManager)
		return c.String(http.StatusOK, errMsg)
	}

	foundScript := false
	var installScript config.InstallScript

	for _, searchScript := range config.VulnaasConfig.VulnaasScripts {
		if searchScript.ID == ID {
			foundScript = true
			installScript = searchScript
			break
		}
	}

	if !foundScript {
		errMsg := fmt.Sprintf("echo \"[x][VulnaaS] Error: VulnaaS ID %d not found.\"", ID)
		return c.String(http.StatusOK, errMsg)
	}

	switch packageManager {
	case "yum":
		return c.String(http.StatusOK, installScript.CmdYum)
	case "apt":
		return c.String(http.StatusOK, installScript.CmdApt)
	case "win":
		return c.String(http.StatusOK, installScript.CmdWindows)
	default:
		return c.String(http.StatusOK, "echo \"[x][VulnaaS] Internal Error: Package Manager.\"")
	}

}

// checkPackageManager returns an error if a string is equal yum or apt.
func checkPackageManager(s string) error {
	if s == "yum" || s == "apt" || s == "win" {
		return nil
	}
	err := errors.New("package manager invalid")
	return err
}

// replaceEnvVars will extract %API_HOST% and %API_PORT% from a string and replace it
// with the proper values from config package. Also replaces %SCRIPT_ID% given an id.
func replaceEnvVars(receivedString string, id string) string {
	step1 := strings.Replace(receivedString, "%API_HOST%", config.APIhost, -1)
	step2 := strings.Replace(step1, "%API_PORT%", config.APIport, -1)
	step3 := strings.Replace(step2, "%SCRIPT_ID%", id, -1)
	return step3
}
