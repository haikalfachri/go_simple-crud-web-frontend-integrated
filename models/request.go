package models

import "time"

type Request struct {
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	BOD         time.Time `json:"bod"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	URL         string    `json:"url"`
}
