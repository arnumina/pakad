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

import (
	"time"

	"github.com/arnumina/pakad/pkg/failure"
)

type (
	// Exclusivity AFAIRE
	Exclusivity string

	// Priority AFAIRE
	Priority int

	// Status AFAIRE
	Status string

	// Job AFAIRE
	Job struct {
		ID          string
		Name        string
		Application string
		Type        string
		Origin      string
		Category    *string
		Exclusivity Exclusivity
		Group       *string
		Priority    Priority
		Emails      *string
		Data        map[string]interface{}
		WID         *string
		WFailed     *bool
		CreatedAt   time.Time
		RefTime     time.Time
		RunAfter    time.Time
		Status      Status
		Attempt     int
	}

	// Result AFAIRE
	Result struct {
		Failure     *failure.Failure
		WaitingTime time.Duration
	}

	// Runner AFAIRE
	Runner interface {
		Run() *Result
	}

	// Step AFAIRE
	Step struct {
		Application string
		Type        string
		Exclusivity Exclusivity
		Group       *string
		Config      map[string]interface{}
		Next        map[string]interface{}
	}

	// Workflow AFAIRE
	Workflow struct {
		ID        string
		Name      string
		Title     string
		Origin    string
		Priority  Priority
		Emails    *string
		FirstStep string
		Steps     map[string]*Step
		Data      map[string]interface{}
		CreatedAt time.Time
		Status    Status
		EndedAt   *time.Time
	}
)

const (
	// No AFAIRE
	No Exclusivity = "no"
	// Itself AFAIRE
	Itself = "itself"
	// Application AFAIRE
	Application = "application"

	// None AFAIRE
	None Priority = 0
	// Low AFAIRE
	Low = 20
	// Medium AFAIRE
	Medium = 50
	// High AFAIRE
	High = 80
	// Critical AFAIRE
	Critical = 100

	// Todo AFAIRE
	Todo Status = "todo"
	// Running AFAIRE
	Running = "running"
	// Pending AFAIRE
	Pending = "pending"
	// Succeeded AFAIRE
	Succeeded = "succeeded"
	// Failed AFAIRE
	Failed = "failed"
	// Aborted AFAIRE
	Aborted = "aborted"
)

// NewJob AFAIRE
func NewJob(id, n, a, t, o string, c *string, e Exclusivity, g *string, p Priority, em *string) *Job {
	return &Job{
		ID:          id,
		Name:        n,
		Application: a,
		Type:        t,
		Origin:      o,
		Category:    c,
		Exclusivity: e,
		Group:       g,
		Priority:    p,
		Emails:      em,
		CreatedAt:   time.Now(),
		RefTime:     time.Now(),
		RunAfter:    time.Now(),
		Status:      Todo,
	}
}

// NewResult AFAIRE
func NewResult(f *failure.Failure, wt time.Duration) *Result {
	return &Result{
		Failure:     f,
		WaitingTime: wt,
	}
}

// NewWorkflow AFAIRE
func NewWorkflow(id, n, t, o string, p Priority, em *string, fs string, s map[string]*Step) *Workflow {
	return &Workflow{
		ID:        id,
		Name:      n,
		Title:     t,
		Origin:    o,
		Priority:  p,
		Emails:    em,
		FirstStep: fs,
		Steps:     s,
		CreatedAt: time.Now(),
		Status:    Running,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
