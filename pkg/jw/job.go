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

package jw

import "time"

// Job AFAIRE
type Job struct {
	ID          string `bson:"_id"`
	Name        string
	Application string
	Type        string
	Origin      string
	Priority    Priority
	CreatedAt   time.Time
}

// NewJob AFAIRE
func NewJob(id, name, application, jt, origin string, priority Priority) *Job {
	return &Job{
		ID:          id,
		Name:        name,
		Application: application,
		Type:        jt,
		Origin:      origin,
		Priority:    priority,
		CreatedAt:   time.Now(),
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
