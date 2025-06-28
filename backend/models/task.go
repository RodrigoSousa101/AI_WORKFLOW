package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Task struct {
	ID		  uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	WorkflowID uuid.UUID      `gorm:"type:uuid;not null" json:"workflow_id"`
	Name 	  string         `json:"name"`
	description string        `json:"description"`
	Users 	[]User         `gorm:"many2many:task_users;"`
	Status string         `json:"status" gorm:"default:'pending'"` // 'pending', 'in_progress', 'completed', etc.
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	
}