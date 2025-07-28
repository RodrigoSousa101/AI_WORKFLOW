package models

import (
	"time"

	"github.com/google/uuid"
)

type TaskUser struct {
	TaskID uuid.UUID `gorm:"primaryKey"`
	UserID uuid.UUID `gorm:"primaryKey"`

	Task *Task `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	User *User `gorm:"constraint:OnDelete:CASCADE;" json:"-"`

	CreatedAt time.Time
}
