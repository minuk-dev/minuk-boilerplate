// Package sqlite provides the SQLite implementation of the repository pattern.
package sqlite

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/adapter/out/sqlite/entity"
	"github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/model"
	outport "github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/port/out"
)

var _ outport.HistoryRepository = (*HistorySqliteAdapter)(nil)

// HistorySqliteAdapter is the SQLite implementation of the HistoryRepository interface.
type HistorySqliteAdapter struct {
	db *gorm.DB
}

// NewHistorySqliteAdapter creates a new HistorySqliteAdapter instance.
func NewHistorySqliteAdapter(
	database *gorm.DB,
) (*HistorySqliteAdapter, error) {
	err := database.AutoMigrate(
		//exhaustruct:ignore
		&entity.History{},
	)
	if err != nil {
		return nil, fmt.Errorf("(HistorySqliteAdapter) failed to migrate history table: %w", err)
	}

	return &HistorySqliteAdapter{
		db: database.Model(
			//exhaustruct:ignore
			&entity.History{},
		),
	}, nil
}

// Get implements out.HistoryRepository.
func (h *HistorySqliteAdapter) Get(puid uuid.UUID) (*model.History, error) {
	var historyEntity entity.History
	if err := h.db.Where("puid = ?", puid).First(&historyEntity).Error; err != nil {
		return nil, fmt.Errorf("(HistorySqliteAdapter) failed to get history: %w", err)
	}

	history := &model.History{
		PUID:      historyEntity.PUID,
		Headers:   historyEntity.Headers,
		CreatedAt: historyEntity.CreatedAt,
	}

	return history, nil
}

// List implements out.HistoryRepository.
func (h *HistorySqliteAdapter) List() ([]model.History, error) {
	var historyEntities []entity.History
	if err := h.db.Find(&historyEntities).Error; err != nil {
		return nil, fmt.Errorf("(HistorySqliteAdapter) failed to list histories: %w", err)
	}

	histories := make([]model.History, len(historyEntities))
	for i, historyEntity := range historyEntities {
		histories[i] = model.History{
			PUID:      historyEntity.PUID,
			Headers:   historyEntity.Headers,
			CreatedAt: historyEntity.CreatedAt,
		}
	}

	return histories, nil
}

// Save implements out.HistoryRepository.
func (h *HistorySqliteAdapter) Save(history *model.History) error {
	//exhaustruct:ignore
	historyEntity := &entity.History{
		PUID:    history.PUID,
		Headers: history.Headers,
	}

	if err := h.db.Save(historyEntity).Error; err != nil {
		return fmt.Errorf("(HistorySqliteAdapter) failed to save history: %w", err)
	}

	return nil
}

// Delete implements out.HistoryRepository.
func (h *HistorySqliteAdapter) Delete(puid uuid.UUID) error {
	var historyEntity entity.History
	if err := h.db.Where("puid = ?", puid).Delete(&historyEntity).Error; err != nil {
		return fmt.Errorf("(HistorySqliteAdapter) failed to delete history: %w", err)
	}

	return nil
}
