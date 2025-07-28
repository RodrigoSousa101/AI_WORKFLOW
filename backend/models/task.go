package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	WorkflowID  uuid.UUID `gorm:"type:uuid;not null" json:"workflow_id"`
	Workflow    *Workflow `gorm:"constraint:OnDelete:CASCADE;" json:"omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:'pending'"`
	Created_at  time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
