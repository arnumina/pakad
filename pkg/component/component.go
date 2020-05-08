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
	"github.com/arnumina/pakad/pkg/jw"
	"github.com/arnumina/pakad/pkg/message"
)

// Applications AFAIRE
type Applications interface {
	Close()
}

// Backend AFAIRE
type Backend interface {
	InsertJob(job *jw.Job) error
	Close()
}

// Bus AFAIRE
type Bus interface {
	NewPublisher(name string, chCapacity int, goCount int) chan<- *message.Message
	Subscribe(cb func(*message.Message), reList ...string) error
	Close()
}

// Logger AFAIRE
type Logger interface {
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

// Scheduler AFAIRE
type Scheduler interface {
	Start()
	Stop()
	Close()
}

// Server AFAIRE
type Server interface {
	Start() error
	Stop()
}

/*
######################################################################################################## @(°_°)@ #######
*/
