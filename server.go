package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/VulnaaS/VulnaaS-API/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	fmt.Println("[*] Starting VulnaaS-API...")

	if err := checkAPIRequirements(); err != nil {
		fmt.Println("[x] Error starting VulnaaS-API:")
		fmt.Println("[x]", err)
		os.Exit(1)
	}

	echoInstance := echo.New()
	echoInstance.HideBanner = true

	echoInstance.Use(middleware.Logger())
	echoInstance.Use(middleware.Recover())
	echoInstance.Use(middleware.RequestID())

	echoInstance.GET("/healthcheck", api.HealthCheck)
	echoInstance.GET("/install/:id", api.ReceiveInstallRequest)
	echoInstance.GET("/scripts/:pm/:id", api.GetScript)

	vulnaasAPIPort := fmt.Sprintf(":%d", getAPIPort())
	echoInstance.Logger.Fatal(echoInstance.Start(vulnaasAPIPort))
}

func checkAPIRequirements() error {

	// check if all environment variables are properly set.
	if err := checkEnvVars(); err != nil {
		return err
	}

	fmt.Println("[*] Environment Variables: OK!")

	return nil
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
