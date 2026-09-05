package database

import (
	"context"
	"fmt"

	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// RunMigrations executes database migrations using goose against the DB
// configured via the given prefix (e.g. "DB_ORDER"). Falls back to the base
// "DB_*" keys when no prefix-specific keys are set.
// path: directory containing migration files.
func RunMigrations(log logger.LoggerInterface, prefix, path string) error {
	if prefix == "" {
		prefix = "DB"
	}

	hostKey := fmt.Sprintf("%s_HOST", prefix)
	portKey := fmt.Sprintf("%s_PORT", prefix)
	userKey := fmt.Sprintf("%s_USERNAME", prefix)
	nameKey := fmt.Sprintf("%s_NAME", prefix)
	passKey := fmt.Sprintf("%s_PASSWORD", prefix)

	host := viper.GetString(hostKey)
	if host == "" {
		host = viper.GetString("DB_HOST")
	}
	port := viper.GetString(portKey)
	if port == "" {
		port = viper.GetString("DB_PORT")
	}
	user := viper.GetString(userKey)
	if user == "" {
		user = viper.GetString("DB_USERNAME")
	}
	dbname := viper.GetString(nameKey)
	if dbname == "" {
		dbname = viper.GetString("DB_NAME")
	}
	password := viper.GetString(passKey)
	if password == "" {
		password = viper.GetString("DB_PASSWORD")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password,
	)

	db, err := goose.OpenDBWithDriver("pgx", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database for migrations: %w", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Error("Failed to close database after migrations", zap.Error(err))
		}
	}()

	log.Info("Running database migrations",
		zap.String("path", path),
		zap.String("dbname", dbname),
		zap.String("prefix", prefix),
	)

	if err := goose.RunContext(context.Background(), "up", db, path); err != nil {
		return fmt.Errorf("migration 'up' failed: %w", err)
	}

	log.Info("Database migrations completed successfully")
	return nil
}
