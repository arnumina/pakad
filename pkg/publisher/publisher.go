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
		name    string
		channel chan<- *message.Message
	}
)

// New AFAIRE
func New(name string, channel chan<- *message.Message) *Publisher {
	return &Publisher{
		name:    name,
		channel: channel,
	}
}

// Publish AFAIRE
func (p *Publisher) Publish(logger component.Logger, topic string, data interface{}) {
	msg := message.New(topic, p.name, data)

	logger.Debug( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"Publish message",
		"id", msg.ID,
		"topic", msg.Topic,
		"publisher", p.name,
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
