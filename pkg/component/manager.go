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
		cpts   []Component
		values map[string]interface{}
	}
)

// NewManager AFAIRE
func NewManager(cpts ...Component) *Manager {
	return &Manager{
		cpts:   cpts,
		values: make(map[string]interface{}),
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
		for _, c := range man.cpts {
			c.Close()
		}
	}()

	for _, c := range man.cpts {
		if err := c.Build(man); err != nil {
			return err
		}

		man.values[c.Name()] = c.Value()
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
