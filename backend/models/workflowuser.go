package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkflowUser struct {
	WorkflowID uuid.UUID `gorm:"primaryKey"`
	UserID     uuid.UUID `gorm:"primaryKey"`

	Workflow *Workflow `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	User     *User     `gorm:"constraint:OnDelete:CASCADE;" json:"-"`

	CreatedAt time.Time
}
