package models

import (
	"gorm.io/gorm"
)

type Machine struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}
