package tests

import (
	"testing"
	"upay/configs"
	"upay/databases"
)

func TestDBConection(t *testing.T) {
	db := database.SetupDBConnection(configs.DbConfig())
	if db == nil {
		t.Error("db is nil")
	}
}
