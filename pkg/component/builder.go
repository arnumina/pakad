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

type (
	// Builder AFAIRE
	Builder interface {
		Name() string
		Build(cm *Manager) error
		Value() interface{}
		Close()
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/
