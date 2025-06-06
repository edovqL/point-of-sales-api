package database

import (
	"point-of-sales-api/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	err := config.LoadConfig("../../../config.test.yml")
	if err != nil {
		panic(err)
	}
}

func TestConnectPostgreSQL(t *testing.T) {
	cfg := config.GetConfig()

	db, err := ConnectPostgreSQL(cfg.DB)
	require.Nil(t, err)
	require.NotNil(t, db)
}
