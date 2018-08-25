package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/VulnaaS/VulnaaS-API/api"
	apiContext "github.com/VulnaaS/VulnaaS-API/context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	fmt.Println("[*] Starting VulnaaS-API...")

	apiConfig := apiContext.GetAPIConfig()

	if err := checkAPIRequirements(apiConfig); err != nil {
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

	vulnaasAPIPort := fmt.Sprintf(":%d", apiConfig.APIPort)
	echoInstance.Logger.Fatal(echoInstance.Start(vulnaasAPIPort))
}

func checkAPIRequirements(apiConfig *apiContext.APIConfig) error {

	// check if all environment variables are properly set.
	if err := checkEnvVars(apiConfig); err != nil {
		return err
	}

	fmt.Println("[*] Environment Variables: OK!")

	return nil
}

func checkEnvVars(apiConfig *apiContext.APIConfig) error {
	envVars := []string{
		"MONGO_HOST",
		"MONGO_DATABASE_NAME",
		"MONGO_DATABASE_USERNAME",
		"MONGO_DATABASE_PASSWORD",
		// "MONGO_PORT", optional -> default value (27017)
		// "MONGO_TIMEOUT", optional -> default value (60s)
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
