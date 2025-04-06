// Package service provides the domain service layer for the application.
package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/model"
	"github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/port/in"
	"github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/port/out"
)

var _ in.HistoryUsecase = (*HistoryService)(nil)

// HistoryService is the service for managing history.
type HistoryService struct {
	HistoryRepository out.HistoryRepository
}

// NewHistoryService creates a new HistoryService instance.
func NewHistoryService(
	historyRepository out.HistoryRepository,
) *HistoryService {
	return &HistoryService{
		HistoryRepository: historyRepository,
	}
}

// Get implements in.HistoryUsecase.
func (h *HistoryService) Get(puid uuid.UUID) (string, error) {
	history, err := h.HistoryRepository.Get(puid)
	if err != nil {
		return "", fmt.Errorf("(HistoryService) failed to get history: %w", err)
	}

	if history == nil {
		return "", nil
	}

	headers := make(map[string]string)

	err = json.Unmarshal([]byte(history.Headers), &headers)
	if err != nil {
		return "", fmt.Errorf("(HistoryService) failed to unmarshal headers: %w", err)
	}

	return headers["Authorization"], nil
}

// List implements in.HistoryUsecase.
func (h *HistoryService) List() ([]model.History, error) {
	histories, err := h.HistoryRepository.List()
	if err != nil {
		return nil, fmt.Errorf("(HistoryService) failed to list histories: %w", err)
	}

	return histories, nil
}

// Save implements in.HistoryUsecase.
func (h *HistoryService) Save(puid uuid.UUID, headers map[string][]string) error {
	newHeader := make(map[string]string, len(headers))
	for k, v := range headers {
		newHeader[k] = v[0]
	}

	headersStr, err := json.Marshal(headers)
	if err != nil {
		return fmt.Errorf("(HistoryService) failed to marshal headers: %w", err)
	}

	err = h.HistoryRepository.Save(&model.History{
		PUID:      puid,
		Headers:   string(headersStr),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("(HistoryService) failed to save history: %w", err)
	}

	return nil
}
