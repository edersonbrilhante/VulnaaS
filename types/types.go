package types

import (
	"time"
)

// ConfigScript represents all information regarding a vulnerability or a
// required service to be installed.
type ConfigScript struct {
	ID         int
	Tittle     string
	Author     string
	Date       time.Time
	Platform   string
	ExploitDB  int
	CmdYum     string
	CmdApt     string
	CmdWindows string
}
