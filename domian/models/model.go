package models

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	UUID      uuid.UUID  `gorm:"primary_key" json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
