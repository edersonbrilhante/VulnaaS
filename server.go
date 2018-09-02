package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/VulnaaS/VulnaaS/api"
	"github.com/VulnaaS/VulnaaS/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {

	fmt.Println("[*] Starting VulnaaS...")

	checkAPIRequirements()

	// loading viper
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		vulnaasError(err)
	}
	if err := viper.Unmarshal(&config.VulnaasConfig); err != nil {
		vulnaasError(err)
	}

	fmt.Println("[*] Viper loaded: OK!")

	echoInstance := echo.New()
	echoInstance.HideBanner = true

	echoInstance.Use(middleware.Logger())
	echoInstance.Use(middleware.Recover())
	echoInstance.Use(middleware.RequestID())

	echoInstance.GET("/healthcheck", api.HealthCheck)
	echoInstance.GET("/install/:input", api.ReceiveInstallRequest)
	echoInstance.GET("/scripts/:pm/:input", api.InstallScript)

	vulnaasAPIPort := fmt.Sprintf(":%d", getAPIPort())
	echoInstance.Logger.Fatal(echoInstance.Start(vulnaasAPIPort))
}

func checkAPIRequirements() {
	if err := checkEnvVars(); err != nil {
		vulnaasError(err)
	}
	fmt.Println("[*] Environment Variables: OK!")
}

func checkEnvVars() error {
	envVars := []string{
		"API_HOST",
		"API_PORT",
	}
	var envIsSet bool
	var allEnvIsSet bool
	var errorString string
	env := make(map[string]string)
	allEnvIsSet = true
	for i := 0; i < len(envVars); i++ {
		env[envVars[i]], envIsSet = os.LookupEnv(envVars[i])
		if !envIsSet {
			errorString = errorString + envVars[i] + " "
			allEnvIsSet = false
		}
	}
	if allEnvIsSet == false {
		finalError := fmt.Sprintf("check environment variables: %s", errorString)
		return errors.New(finalError)
	}
	return nil
}

func getAPIPort() int {
	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		apiPort = 9999
	}
	return apiPort
}

func vulnaasError(err error) {
	fmt.Println("[x] Error starting VulnaaS:")
	fmt.Println("[x]", err)
	os.Exit(1)
}
