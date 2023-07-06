package models

import "time"

type Request struct {
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	DOB         time.Time `json:"dob"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	URL         string    `json:"url"`
}
