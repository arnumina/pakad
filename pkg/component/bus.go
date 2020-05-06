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

import "github.com/arnumina/pakad/pkg/message"

// Bus AFAIRE
type Bus interface {
	NewPublisher(name string, chCapacity int, goCount int) chan<- *message.Message
	Subscribe(cb func(Logger, *message.Message), reList ...string) error
	Close()
}

/*
######################################################################################################## @(°_°)@ #######
*/
