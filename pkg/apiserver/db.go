package apiserver

import (
	"fmt"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InvalidDBFormatError is an error type that indicates an invalid database format.
type InvalidDBFormatError struct {
	Msg string
}

// NewDB creates a new SQLite database connection using the provided settings.
func NewDB(
	setting *Settings,
) (*gorm.DB, error) {
	filename, err := parseSqliteFileName(setting.DB)
	if err != nil {
		return nil, err
	}

	//exhaustruct:ignore
	config := &gorm.Config{}

	db, err := gorm.Open(sqlite.Open(filename), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create sqlite db: %w", err)
	}

	return db, nil
}

func parseSqliteFileName(db string) (string, error) {
	parsed := strings.Split(db, ":")
	//nolint:mnd
	if len(parsed) != 2 {
		return "", &InvalidDBFormatError{
			Msg: "could not parse DB format",
		}
	}

	if parsed[0] != "sqlite" {
		return "", &InvalidDBFormatError{
			Msg: "DB format is not sqlite",
		}
	}

	if parsed[1] == "" {
		return "", &InvalidDBFormatError{
			Msg: "DB filename is empty",
		}
	}

	return parsed[1], nil
}

func (e *InvalidDBFormatError) Error() string {
	return "invalid DB format: " + e.Msg
}
