// Package model provides the domain model for the application.
package model

import (
	"time"

	"github.com/google/uuid"
)

// History represents the history model.
type History struct {
	PUID      uuid.UUID
	Headers   string
	CreatedAt time.Time
}
