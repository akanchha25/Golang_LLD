package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Payment represents the payment entity
type Payment struct {
	gorm.Model

	Id     uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserId uint      `json:"userId"`
	Amount float64   `json:"amount"`
	Method string    `json:"method"`
	Status string    `json:"status"`
}
