package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name      string `json:"name"`
	Backstory string `json:"backstory"`
}
