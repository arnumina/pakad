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
	id           string
	name         string
	version      string
	builtAt      time.Time
	startedAt    time.Time
	config       *value.Value
	applications component.Applications
	bus          component.Bus
	logger       component.Logger
	scheduler    component.Scheduler
	server       component.Server
	runtime      bool
}

// New AFAIRE
func New(name, version, builtAt string) *Runner {
	return &Runner{
		id:        util.NewUUID(),
		name:      name,
		version:   version,
		builtAt:   util.UnixToTime(builtAt),
		startedAt: time.Now(),
	}
}

// ID AFAIRE
func (runner *Runner) ID() string {
	return runner.id
}

// Name AFAIRE
func (runner *Runner) Name() string {
	return runner.name
}

// Version AFAIRE
func (runner *Runner) Version() string {
	return runner.version
}

// BuiltAt AFAIRE
func (runner *Runner) BuiltAt() time.Time {
	return runner.builtAt
}

// StartedAt AFAIRE
func (runner *Runner) StartedAt() time.Time {
	return runner.startedAt
}

// SetConfig AFAIRE
func (runner *Runner) SetConfig(config *value.Value) {
	runner.config = config
}

// Config AFAIRE
func (runner *Runner) Config() *value.Value {
	return runner.config
}

// SetApplications AFAIRE
func (runner *Runner) SetApplications(applications component.Applications) {
	runner.applications = applications
}

// Applications AFAIRE
func (runner *Runner) Applications() component.Applications {
	return runner.applications
}

// SetBus AFAIRE
func (runner *Runner) SetBus(bus component.Bus) {
	runner.bus = bus
}

// Bus AFAIRE
func (runner *Runner) Bus() component.Bus {
	return runner.bus
}

// SetLogger AFAIRE
func (runner *Runner) SetLogger(logger component.Logger) {
	runner.logger = logger
}

// Logger AFAIRE
func (runner *Runner) Logger() component.Logger {
	return runner.logger
}

// SetScheduler AFAIRE
func (runner *Runner) SetScheduler(scheduler component.Scheduler) {
	runner.scheduler = scheduler
}

// Scheduler AFAIRE
func (runner *Runner) Scheduler() component.Scheduler {
	return runner.scheduler
}

// SetServer AFAIRE
func (runner *Runner) SetServer(server component.Server) {
	runner.server = server
}

// Server AFAIRE
func (runner *Runner) Server() component.Server {
	return runner.server
}

// SetRuntime AFAIRE
func (runner *Runner) SetRuntime() {
	runner.runtime = true
}

// Runtime AFAIRE
func (runner *Runner) Runtime() bool {
	return runner.runtime
}

/*
######################################################################################################## @(°_°)@ #######
*/
