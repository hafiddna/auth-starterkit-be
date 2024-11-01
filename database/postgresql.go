package database

import (
	"database/sql"
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgreSQL interface {
	Connect(db string) (*sql.DB, *gorm.DB, error)
	Disconnect(*sql.DB) error
}

type postgreSQL struct {
	config config.CfgStruct
}

func NewPostgreSQL(config config.CfgStruct) PostgreSQL {
	return &postgreSQL{
		config: config,
	}
}

func (p *postgreSQL) Connect(db string) (*sql.DB, *gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.config.App.PostgreSQL.Host,
		p.config.App.PostgreSQL.Port,
		p.config.App.PostgreSQL.Username,
		p.config.App.PostgreSQL.Password,
		db,
	)
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
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
		return nil, nil, err
	}

	return sqlDB, gormDB, nil
}

func (p *postgreSQL) Disconnect(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}

	return nil
}
