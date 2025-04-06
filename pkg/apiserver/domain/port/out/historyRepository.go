// Package out provides the output ports for the application.
package out

import (
	"github.com/google/uuid"

	"github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/model"
)

// HistoryRepository is the interface for the history repository.
type HistoryRepository interface {
	Get(puid uuid.UUID) (*model.History, error)
	List() ([]model.History, error)
	Save(history *model.History) error
	Delete(puid uuid.UUID) error
}
