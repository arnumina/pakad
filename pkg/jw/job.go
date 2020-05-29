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
	ID          string
	Name        string
	Application string
	Type        string
	Origin      string
	Category    string
	Exclusivity Exclusivity
	Group       string
	Priority    Priority
	WID         string
	WFailed     bool
	CreatedAt   time.Time
	RefTime     time.Time
	RunAfter    time.Time
	Status      Status
	Attempt     int
}

// NewJob AFAIRE
func NewJob(
	id, name, application, jt, origin, category string, exclusivity Exclusivity, group string, priority Priority) *Job {
	return &Job{
		ID:          id,
		Name:        name,
		Application: application,
		Type:        jt,
		Origin:      origin,
		Category:    category,
		Exclusivity: exclusivity,
		Group:       group,
		Priority:    priority,
		WID:         "",
		WFailed:     false,
		CreatedAt:   time.Now(),
		RefTime:     time.Now(),
		RunAfter:    time.Now(),
		Status:      Todo,
		Attempt:     0,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
