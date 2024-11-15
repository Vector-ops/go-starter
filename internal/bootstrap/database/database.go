package database

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/vector-ops/go-starter/configs"
)

type Database struct {
	DB  *sqlx.DB
	RDB *redis.Client
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) SetupPostgres() {
	db.DB = configs.NewPostgresDB()

}

func (db *Database) SetupRedis() {
	db.RDB = configs.NewRedisClient()
}
