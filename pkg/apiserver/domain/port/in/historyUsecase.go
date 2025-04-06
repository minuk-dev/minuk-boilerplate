// Package in provides the input ports for the application.
package in

import (
	"github.com/google/uuid"

	"github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/model"
)

// HistoryUsecase is the interface for the history use case.
type HistoryUsecase interface {
	Get(puid uuid.UUID) (string, error)
	List() ([]model.History, error)
	Save(puid uuid.UUID, headers map[string][]string) error
}
