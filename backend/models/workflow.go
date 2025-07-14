package models

import (
	"time"

	"github.com/google/uuid"
)

type Workflow struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;constraint:OnDelete:CASCADE" json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:'draft'"` // 'draft', 'active', 'archived'
	Workers     []User    `gorm:"many2many:workflow_users;constraint:OnDelete:CASCADE" json:"workers"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
