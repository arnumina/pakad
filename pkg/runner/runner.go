/*
#######
##                    __           __
##         ___  ___ _/ /_____ ____/ /
##        / _ \/ _ `/  '_/ _ `/ _  /
##       / .__/\_,_/_/\_\\_,_/\_,_/
##      /_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package runner

import (
	"time"

	"github.com/arnumina/pakad/pkg/component"
	"github.com/arnumina/pakad/pkg/util"
	"github.com/arnumina/pakad/pkg/value"
)

// Runner AFAIRE
type Runner struct {
	ID           string
	Name         string
	Version      string
	BuiltAt      time.Time
	StartedAt    time.Time
	Config       *value.Value
	Applications component.Applications
	Backend      component.Backend
	Bus          component.Bus
	Logger       component.Logger
	Scheduler    component.Scheduler
	Server       component.Server
	Runtime      bool
}

// New AFAIRE
func New(name, version, builtAt string) *Runner {
	return &Runner{
		ID:        util.NewUUID(),
		Name:      name,
		Version:   version,
		BuiltAt:   util.UnixToTime(builtAt),
		StartedAt: time.Now(),
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
