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

package component

import (
	"time"

	"github.com/arnumina/pakad/pkg/jw"
	"github.com/arnumina/pakad/pkg/message"
	"github.com/arnumina/pakad/pkg/value"
)

type (
	// Component AFAIRE
	Component interface {
		Name() string
		Start(cm *Manager) (interface{}, error)
		Stop(cm *Manager, err error)
	}

	// Application AFAIRE
	Application interface {
		PluginName() string
		PluginVersion() string
		PluginBuiltAt() time.Time

		Name() string
		Start(cm *Manager) error
		RunJob(job *jw.Job) error
		Stop()
	}

	// Applications AFAIRE
	Applications interface {
		Find(appName string) Application
		List() []Application
		Stop(cm *Manager, err error)
	}

	// Backend AFAIRE
	Backend interface {
		InsertJob(job *jw.Job) error
		Close()
	}

	// Bus AFAIRE
	Bus interface {
		NewPublisher(name string, chCapacity, consumer int) chan<- *message.Message
		Subscribe(cb func(*message.Message), reList ...string) error
		Close()
	}

	// Config AFAIRE
	Config interface {
		MaybeGet(keys ...string) (*value.Value, error)
	}

	// Logger AFAIRE
	Logger interface {
		Clone(prefix string) Logger
		Trace(msg string, ctx ...interface{})
		Debug(msg string, ctx ...interface{})
		Info(msg string, ctx ...interface{})
		Notice(msg string, ctx ...interface{})
		Warning(msg string, ctx ...interface{})
		Error(msg string, ctx ...interface{})
		Critical(msg string, ctx ...interface{})
		Close()
	}

	// Runner AFAIRE
	Runner interface {
		ID() string
		Name() string
		Version() string
		BuiltAt() time.Time
		StartedAt() time.Time
	}

	// Scheduler AFAIRE
	Scheduler interface {
		Start()
		Stop()
	}

	// Server AFAIRE
	Server interface {
		Start() error
		Stop()
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
