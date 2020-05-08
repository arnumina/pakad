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
	// Manager AFAIRE
	Manager struct {
		cpts map[string]Component
	}
)

// Get AFAIRE
func (man *Manager) Get(name string) interface{} {
	c, ok := man.cpts[name]
	if !ok {
		return nil
	}

	return c.Value()
}

/*
######################################################################################################## @(°_°)@ #######
*/
