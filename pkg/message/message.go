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

package message

import (
	"github.com/arnumina/pakad/pkg/util"
)

type (
	// Message AFAIRE
	Message struct {
		ID        string
		Topic     string
		Data      interface{}
		Publisher string
	}
)

// New AFAIRE
func New(topic string, data interface{}) *Message {
	return &Message{
		ID:    util.NewUUID(),
		Topic: topic,
		Data:  data,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
