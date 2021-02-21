package models

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	UUID      uuid.UUID `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
