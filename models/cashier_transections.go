package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type TransectionType string

const (
	RECEIVE  TransectionType = "RECEIVE"
	CHANGE   TransectionType = "CHANGE"
	DEPOSIT  TransectionType = "DEPOSIT"
	WITHDRAW TransectionType = "WITHDRAW"
)

func (self *TransectionType) Scan(value interface{}) error {
	*self = TransectionType(value.([]byte))
	return nil
}

func (self TransectionType) Value() (driver.Value, error) {
	return string(self), nil
}

type CashierTransections struct {
	gorm.Model
	SessionId  string          `gorm:"type:varchar(255);not null" json:"session_id"`
	Type       TransectionType `gorm:"type:ENUM('RECEIVE', 'CHANGE', 'DEPOSIT', 'WITHDRAW')" gorm:"column:car_type"`
	MoneyValue float32         `gorm:"type:decimal(10,2);not null" json:"money_value"`
	Amount     int32           `gorm:"not null" json:"amount"`
}
