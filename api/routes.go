package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Vulnaas/Vulnaas-API/scripts"
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
		return c.String(http.StatusOK, "echo \"[x][VulnaaS] Error: ID is not an integer\"")
	}

	var (
		UnixScript = "" +
			"#!/bin/bash\n" +
			"#\n" +
			"which yum 1> /dev/null\n" +
			"if [ $? -eq 0 ]; then\n" +
			// curl to API to get yum based scripts.
			fmt.Sprintf("	curl -s http://%s:%s/scripts/yum/%d | sh\n", os.Getenv("API_HOST"), os.Getenv("API_PORT"), ID) +
			"else\n" +
			// curl to API to get apt-get based scripts.
			fmt.Sprintf("	curl -s http://%s:%s/scripts/apt/%d | sh\n", os.Getenv("API_HOST"), os.Getenv("API_PORT"), ID) +
			"fi\n"

		WinScript = "" +
			"PS1 script would go here\n"
	)

	// ID > 1000 represents Unix vulnerabilities scripts.
	if ID > 1000 {
		return c.String(http.StatusOK, UnixScript)
	}

	// return Windows script.
	return c.String(http.StatusOK, WinScript)

}

// GetScript returns a script based on its ID and its package manager.
func GetScript(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusOK, "echo \"[x][VulnaaS] Error: ID is not an integer.\"")
	}

	scriptConfig, err := scripts.Get(ID)
	if err != nil {
		return c.String(http.StatusOK, "echo \"[x][VulnaaS] Error: ID not found.\"")
	}

	packageManager := c.Param("pm")
	err = checkPackageManager(packageManager)
	if err != nil {
		return c.String(http.StatusOK, "echo \"[x][VulnaaS] Error: Package Manager invalid.\"")
	}

	switch packageManager {
	case "yum":
		return c.String(http.StatusOK, scriptConfig.CmdYum)
	case "apt":
		return c.String(http.StatusOK, scriptConfig.CmdApt)
	case "win":
		return c.String(http.StatusOK, scriptConfig.CmdWindows)
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
