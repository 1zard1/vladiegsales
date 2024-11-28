package models

import "time"

// User - Data Model To Store user Data in
type BodyParams struct {
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Departure string    `json:"departure"`
	Arrival   string    `json:"arrival"`
}
