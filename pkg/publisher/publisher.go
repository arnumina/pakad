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

package publisher

import (
	"github.com/arnumina/pakad/pkg/component"
	"github.com/arnumina/pakad/pkg/message"
)

type (
	// Publisher AFAIRE
	Publisher struct {
		channel chan<- *message.Message
	}
)

// New AFAIRE
func New(channel chan<- *message.Message) *Publisher {
	return &Publisher{
		channel: channel,
	}
}

// Publish AFAIRE
func (p *Publisher) Publish(logger component.Logger, topic, publisher string, data interface{}) {
	msg := message.New(topic, publisher, data)

	logger.Debug( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"Publish message",
		"id", msg.ID,
		"topic", msg.Topic,
		"publisher", msg.Publisher,
	)

	p.channel <- msg
}

// Close AFAIRE
func (p *Publisher) Close() {
	close(p.channel)
}

/*
######################################################################################################## @(°_°)@ #######
*/
