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

// Exclusivity AFAIRE
type Exclusivity string

const (
	// No AFAIRE
	No Exclusivity = "no"
	// Itself AFAIRE
	Itself = "itself"
	// Application AFAIRE
	Application = "application"
)

// Priority AFAIRE
type Priority int

const (
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
)

// Status AFAIRE
type Status string

const (
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

type (
	// Result AFAIRE
	Result struct {
		Failure     *failure.Failure
		WaitingTime time.Duration
	}

	// Runner AFAIRE
	Runner interface {
		Run() *Result
	}
)

// NewResult AFAIRE
func NewResult(f *failure.Failure, wt time.Duration) *Result {
	return &Result{
		Failure:     f,
		WaitingTime: wt,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
