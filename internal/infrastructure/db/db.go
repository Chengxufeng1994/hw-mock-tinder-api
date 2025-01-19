package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/migrations"
)

func NewDB(config *config.Database) (*gorm.DB, func(), error) {
	var loggerLvl logger.LogLevel
	if config.LogLevel == "debug" {
		loggerLvl = logger.Info
	} else {
		loggerLvl = logger.Silent
	}

	_default := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      loggerLvl,   // Log level
			Colorful:      true,        // Disable color
			// Ignore ErrRecordNotFound error for logger
			IgnoreRecordNotFoundError: false,
			// Don't include params in the SQL log
			ParameterizedQueries: false,
		},
	)
	cfg := &gorm.Config{
		Logger: _default,
		// PrepareStmt:            true,
		SkipDefaultTransaction: false,
		// TranslateError: true,
	}
	dsn := fmt.Sprintf("host=%s port=%d  user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC", config.Host, config.Port, config.User, config.Password, config.DBName)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), cfg)
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, nil, err
	}

	if config.MaxLifetime > 0 {
		t, err := time.ParseDuration(fmt.Sprintf("%ds", config.MaxLifetime))
		if err != nil {
			return nil, nil, err
		}
		sqlDB.SetConnMaxLifetime(t)
	}
	if config.MaxIdleTime > 0 {
		t, err := time.ParseDuration(fmt.Sprintf("%ds", config.MaxIdleTime))
		if err != nil {
			return nil, nil, err
		}
		sqlDB.SetConnMaxIdleTime(t)
	}
	if config.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	}

	if err := MigrationDB(sqlDB); err != nil {
		return nil, nil, err
	}

	return db, func() {
		_ = sqlDB.Close()
	}, nil
}

func MigrationDB(db *sql.DB) error {
	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "."); err != nil {
		return err
	}

	return nil
}
