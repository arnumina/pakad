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

package jw

import (
	"time"

	"github.com/arnumina/pakad/pkg/failure"
)

// Priority AFAIRE
type Priority int

const (
	// None AFAIRE
	None = 0
	// Low AFAIRE
	Low = 20
	// Medium AFAIRE
	Medium = 50
	// High AFAIRE
	High = 80
	// Critical AFAIRE
	Critical = 100
)

type (
	// Result AFAIRE
	Result struct {
		Completed   bool
		Failure     failure.Failure
		WaitingTime time.Duration
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
