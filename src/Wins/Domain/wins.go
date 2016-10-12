package domain

import "time"

type win struct{
	Success bool `json:"success"`
	Time time.Time `json:"time"`
}

type wins []win
