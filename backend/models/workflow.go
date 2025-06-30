package models

import (
	"time"

	"github.com/google/uuid"
)

type workflow struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:'draft'"` // 'draft', 'active', 'archived'
	Tasks       []Task    `gorm:"foreignKey:WorkflowID" json:"tasks"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
