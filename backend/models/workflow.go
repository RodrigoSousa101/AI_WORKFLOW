package models

import (
	"time"

	"github.com/google/uuid"
)

type Workflow struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User        *User     `gorm:"constraint:OnDelete:CASCADE;" json:"omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:'In Progress'"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
