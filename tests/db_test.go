package tests

import (
	"testing"
	db "upay/databases"
)

func TestDBConection(t *testing.T) {
	if db.Mysql() == nil {
		t.Errorf("db is null")
	}
}
