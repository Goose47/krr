package services

import (
	"context"
	"log/slog"
)

type KRRService struct {
	log *slog.Logger
}

func NewKRRService(log *slog.Logger) *KRRService {
	return &KRRService{
		log: log,
	}
}

func (con *KRRService) Recs(ctx context.Context) (string, error) {
	return "Hello", nil
}
