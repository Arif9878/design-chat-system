package initialize

import (
	"github.com/gocql/gocql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/Arif9878/design-chat-system/go/grpc-uberfx/internal/config"
)

// NewScyllaDB initializes ScyllaDB
func NewScyllaDB(cfg *config.DBConfig, logger *zap.Logger) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cfg.ScyllaDB.Hosts...)
	cluster.Keyspace = cfg.ScyllaDB.Keyspace
	cluster.Consistency = gocql.ParseConsistency(cfg.ScyllaDB.Consistency)

	session, err := cluster.CreateSession()
	if err != nil {
		logger.Error("Failed to connect to ScyllaDB", zap.Error(err))
		return nil, err
	}

	logger.Info("Connected to ScyllaDB successfully")
	return session, nil
}
