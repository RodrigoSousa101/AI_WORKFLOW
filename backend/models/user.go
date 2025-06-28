package models

import (
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
)

type struct User {
	ID 	uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Role string `json:"role" gorm:"default:'user'"` // 'user', 'admin', etc.
	Speciality []string `json:"speciality" gorm:"type:text[]"` // Array of specialities
	isavailable bool `json:"is_available" gorm:"default:true"` // Indicates if the user is available for tasks
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}