package config

import (
	"os"
)

// VulnaasConfiguration holds all information regarding the project.
type VulnaasConfiguration struct {
	VulnaasScripts            []InstallScript
	ServiceScripts            []InstallScript
	CheckPackageManagerScript CheckPackageManagerScript
}

// InstallScript represents all information regarding a vulnerability or a
// required service to be installed.
type InstallScript struct {
	ID         int
	Title      string
	Author     string
	Date       string
	Platform   string
	ExploitDB  int
	CmdYum     string
	CmdApt     string
	CmdWindows string
}

// CheckPackageManagerScript represents all information regarding the scripts
// required to check what package manager is installed in the box.
type CheckPackageManagerScript struct {
	UnixCmd    string
	WindowsCmd string
}

// Project configuration.
var (
	VulnaasConfig = new(VulnaasConfiguration)
	APIhost       = os.Getenv("API_HOST")
	APIport       = os.Getenv("API_PORT")
)
