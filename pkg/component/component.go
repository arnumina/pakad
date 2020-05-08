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

	"github.com/arnumina/pakad/pkg/value"
)

type (
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
)

/*
######################################################################################################## @(°_°)@ #######
*/
