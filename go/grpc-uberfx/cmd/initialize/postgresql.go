package initialize

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"go.uber.org/zap"

	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/internal/config"
)

// NewPostgresPrimaryDB initializes PostgreSQL Primary
func NewPostgresPrimaryDB(cfg *config.DBConfig, logger *zap.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.PostgreSQL.Primary.GenerateDSN())
	if err != nil {
		logger.Error("Failed to connect to PostgreSQL Primary", zap.Error(err))
		return nil, err
	}

	db.SetMaxIdleConns(cfg.PostgreSQL.Primary.MaxIdleConn)
	db.SetMaxOpenConns(cfg.PostgreSQL.Primary.MaxOpenConn)

	logger.Info("Connected to PostgreSQL Primary successfully")
	return db, nil
}

// NewPostgresStandbyDB initializes PostgreSQL Standby
func NewPostgresStandbyDB(cfg *config.DBConfig, logger *zap.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.PostgreSQL.Standby.GenerateDSN())
	if err != nil {
		logger.Error("Failed to connect to PostgreSQL Standby", zap.Error(err))
		return nil, err
	}

	db.SetMaxIdleConns(cfg.PostgreSQL.Standby.MaxIdleConn)
	db.SetMaxOpenConns(cfg.PostgreSQL.Standby.MaxOpenConn)

	logger.Info("Connected to PostgreSQL Standby successfully")
	return db, nil
}
