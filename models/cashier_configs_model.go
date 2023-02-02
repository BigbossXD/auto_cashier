package models

import (
	"gorm.io/gorm"
)

type CashierConfigs struct {
	gorm.Model
	MachineId     uint    `gorm:"not null" json:"machine_id"`
	MoneyValue    float32 `gorm:"type:decimal(10,2);not null" json:"money_value"`
	MaximumAmount int32   `gorm:"not null" json:"maximum_amount"`
	CurrentAmount int32   `gorm:"not null" json:"current_amount"`
}
