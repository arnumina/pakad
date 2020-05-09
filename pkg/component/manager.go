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

import "errors"

type (
	// Manager AFAIRE
	Manager struct {
		cpts      []Component
		instances map[string]interface{}
	}
)

var (
	// ErrNoError AFAIRE
	ErrNoError = errors.New("no error")
)

// NewManager AFAIRE
func NewManager(cpts ...Component) *Manager {
	return &Manager{
		cpts:      cpts,
		instances: make(map[string]interface{}),
	}
}

// Get AFAIRE
func (m *Manager) Get(name string) interface{} {
	instance, ok := m.instances[name]
	if !ok {
		// TODO
		return nil
	}

	return instance
}

func (m *Manager) start() error {
	for _, c := range m.cpts {
		instance, err := c.Start(m)
		if err != nil {
			if errors.Is(err, ErrNoError) {
				return nil
			}

			return err
		}

		m.instances[c.Name()] = instance
	}

	return nil
}

func (m *Manager) stop(err error) {
	l := len(m.cpts)

	for i := range m.cpts {
		c := m.cpts[l-1-i]

		_, ok := m.instances[c.Name()]
		if ok {
			c.Stop(m, err)
		}
	}
}

// Run AFAIRE
func (m *Manager) Run() error {
	err := m.start()
	m.stop(err)

	return err
}

/*
######################################################################################################## @(°_°)@ #######
*/
