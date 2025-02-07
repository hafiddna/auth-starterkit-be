package database

import (
	"database/sql"
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ConnectToPostgres() (gormDB *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Config.App.PostgreSQL.Host,
		config.Config.App.PostgreSQL.Port,
		config.Config.App.PostgreSQL.Username,
		config.Config.App.PostgreSQL.Password,
		config.Config.App.PostgreSQL.Database,
	)
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		//DryRun: true,
		//SkipHooks: true, // should be using &gorm.Session{SkipHooks: true}
		//QueryFields: true,
		//SkipDefaultTransaction: true,
		//PrepareStmt:            true,
		//DisableAutomaticPing:   true,
	})

	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
