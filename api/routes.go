package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// HealthCheck is the heath check function.
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING!\n")
}

// ReceiveRequest returns a script to do further enumeration of the box being used
// to install the vulnerability.
func ReceiveRequest(c echo.Context) error {

	var scriptCheck string

	// check if vulnaasID is an integer.
	vulnaasID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// vulnaasID > 1000 are Unix based.
	if vulnaasID > 1000 {
		scriptCheck = "" +
			"#!/bin/bash\n" +
			"#\n" +
			"which yum 1> /dev/null\n" +
			"if [ $? -eq 0 ]; then\n" +
			"echo \"Found yum\"\n" + // curl to vulnaas-api based on yum need to be done
			"else\n" +
			"echo \"Not found\"\n" + // curl to vulnaas-api based on apt-get need to be done
			"fi\n"
	} else {
		//scriptCheck = "PS1 script would go here"
	}

	return c.String(http.StatusOK, scriptCheck)
}
