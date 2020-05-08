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
		builders []Builder
		values   map[string]interface{}
	}
)

// NewManager AFAIRE
func NewManager(builders ...Builder) *Manager {
	return &Manager{
		builders: builders,
		values:   make(map[string]interface{}),
	}
}

// Get AFAIRE
func (man *Manager) Get(name string) interface{} {
	value, ok := man.values[name]
	if !ok {
		// TODO
		return nil
	}

	return value
}

// Run AFAIRE
func (man *Manager) Run() error {
	defer func() {
		for _, b := range man.builders {
			b.Close()
		}
	}()

	for _, b := range man.builders {
		if err := b.Build(man); err != nil {
			return err
		}

		man.values[b.Name()] = b.Value()
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
