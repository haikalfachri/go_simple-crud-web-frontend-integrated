package models

import (
	"time"

	"gorm.io/gorm"
)

type Biodata struct {
	gorm.Model
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	DOB         time.Time `json:"dob"`
	Gender      string    `json:"gender" gorm:"type:enum('male', 'female', 'not-selected');default:'not-selected';not_null"`
	Address     string    `json:"address"`
	URL         string    `json:"url"`
}
