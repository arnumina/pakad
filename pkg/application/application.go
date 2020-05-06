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

package application

import (
	"time"

	"github.com/arnumina/pakad/pkg/runner"
	"github.com/arnumina/pakad/pkg/value"
)

// Application AFAIRE
type Application interface {
	PluginName() string
	PluginVersion() string
	PluginBuiltAt() time.Time

	Name() string

	Initialize(runner *runner.Runner, cfg *value.Value) error
	Start(runner *runner.Runner) error
	Stop()
}

/*
######################################################################################################## @(°_°)@ #######
*/
