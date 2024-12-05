// Package services provides application business logic.
package services

import (
	"context"
	"log/slog"
)

// TestulatorService provides api for calculating aggregates.
type TestulatorService struct {
	log *slog.Logger
}

// NewTestulatorService is a constructor for TestulatorService.
func NewTestulatorService(log *slog.Logger) *TestulatorService {
	return &TestulatorService{
		log: log,
	}
}

// Testulate calculates aggregates based on params and program.
func (s *TestulatorService) Testulate(
	_ context.Context,
) (string, error) {
	return "hello", nil
}
