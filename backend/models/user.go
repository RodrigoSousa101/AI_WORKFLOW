package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name             string         `json:"name"`
	Email            string         `json:"email" gorm:"unique"`
	Password         string         `json:"password"`
	Phone            string         `json:"phone"`
	Role             string         `json:"role" gorm:"default:'user'"`
	Speciality       pq.StringArray `json:"speciality" gorm:"type:text[]"`
	IsAvailable      bool           `json:"is_available" gorm:"default:true"`
	CreatedWorkflows []Workflow     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CreatedAt        time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}
