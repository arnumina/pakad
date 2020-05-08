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
		builders []Builder
		values   map[string]interface{}
	}
)

var (
	// ErrNoError AFAIRE
	ErrNoError = errors.New("no error")
)

// NewManager AFAIRE
func NewManager(builders ...Builder) *Manager {
	return &Manager{
		builders: builders,
		values:   make(map[string]interface{}),
	}
}

// Get AFAIRE
func (m *Manager) Get(name string) interface{} {
	value, ok := m.values[name]
	if !ok {
		// TODO
		return nil
	}

	return value
}

func (m *Manager) build() error {
	for _, b := range m.builders {
		if err := b.Build(m); err != nil {
			if errors.Is(err, ErrNoError) {
				return nil
			}

			return err
		}

		m.values[b.Name()] = b.Value()
	}

	return nil
}

func (m *Manager) close() {
	for _, b := range m.builders {
		b.Close()
	}
}

// Run AFAIRE
func (m *Manager) Run() error {
	defer m.close()
	return m.build()
}

/*
######################################################################################################## @(°_°)@ #######
*/
