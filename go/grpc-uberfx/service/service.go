package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// Service interface untuk gRPC
type Service interface {
	SayHello(ctx context.Context, name string) (string, error)
}

// Implementasi Service
type serviceImpl struct {
	logger *zap.Logger
}

// NewService membuat instance baru service
func NewService(logger *zap.Logger) Service {
	return &serviceImpl{logger: logger}
}

// SayHello mengembalikan pesan sapaan
func (s *serviceImpl) SayHello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
