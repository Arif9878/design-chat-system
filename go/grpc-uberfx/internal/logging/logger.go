package logging

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// NewLogger menyediakan logger Uber Zap untuk aplikasi
func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment() // Bisa diganti dengan zap.NewDevelopment() untuk log lebih readable
	if err != nil {
		return nil, err
	}
	return logger, nil
}

// Modul logging untuk Uber Fx
var Module = fx.Provide(NewLogger)
