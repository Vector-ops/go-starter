package tests

import (
	"testing"

	"github.com/vector-ops/go-starter/configs"
)

func TestPostgresConnection(t *testing.T) {
	db := configs.NewPostgresDB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		t.Fatalf("Failed to connect to db, %v", err)
	}
}
