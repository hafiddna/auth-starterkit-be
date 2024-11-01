package database

import (
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	postgreConfig = config.NewConfig().GetConfig()

	postgre = NewPostgreSQL(postgreConfig)
)

func TestPostgreSQL_ConnectDisconnect(t *testing.T) {
	sqlDB, gormDB, err := postgre.Connect("school_intelligence_suite")
	assert.NotNil(t, sqlDB)
	assert.NotNil(t, gormDB)
	assert.Nil(t, err)
	err = postgre.Disconnect(sqlDB)
	assert.Nil(t, err)
}
