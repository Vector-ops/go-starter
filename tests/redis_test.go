package tests

import (
	"context"
	"testing"
	"time"

	"github.com/vector-ops/go-starter/configs"
)

func TestRedisConnection(t *testing.T) {
	client := configs.NewRedisClient()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	cmd := client.Ping(ctx)
	pong, err := cmd.Result()
	if err != nil {
		t.Fatalf("Failed ping: %v", err)
	}

	expected := "PONG"
	if pong != expected {
		t.Fatalf("expected: %s\ngot: %s", expected, pong)
	}
}
