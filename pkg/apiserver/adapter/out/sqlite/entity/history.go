// Package entity provides the entity structure for the database.
package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// History represents the history entity in the database.
type History struct {
	gorm.Model

	PUID    uuid.UUID `gorm:"index:idx_puid"`
	Headers string    // json
}
