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

import "time"

// Job AFAIRE
type Job struct {
	ID          string
	Application string
	Type        string
	Name        string
	Origin      string
	Priority    Priority
	CreatedAt   time.Time
}

/*
######################################################################################################## @(°_°)@ #######
*/
