package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkflowUser struct {
	WorkflowID uuid.UUID `gorm:"primaryKey"`
	UserID     uuid.UUID `gorm:"primaryKey"`
	CreatedAt  time.Time
}
