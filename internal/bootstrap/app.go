package bootstrap

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/vector-ops/go-starter/configs"
	"github.com/vector-ops/go-starter/internal/bootstrap/database"
	"github.com/vector-ops/go-starter/internal/bootstrap/web"
)

type App struct {
	server *echo.Echo
	db     *database.Database
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	a.setupEnv()
	a.setupHTTP()
	a.setupDatabase()

	port := fmt.Sprintf(":%s", configs.GetEnv("PORT", "3000"))
	err := a.server.Start(port)
	if err != nil {
		log.Println(err)
		a.closeServices()
		os.Exit(1)
	}
}

func (a *App) setupEnv() {
	configs.LoadEnv()
}

func (a *App) setupHTTP() {
	a.server = web.SetupServer()

	middlewares := web.NewMiddleware(a.server, a.db)
	middlewares.Init()
}

func (a *App) setupDatabase() {
	db := database.NewDatabase()
	db.SetupPostgres()
	db.SetupRedis()

	a.db = db
}

// closeServices function gracefully shuts
// down all the services in case of a server crash
func (a *App) closeServices() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		log.Println("Failed to shutdown server.")
	}

	if err := a.db.DB.Close(); err != nil {
		log.Println("Failed to shutdown Postgres database")
	}

	if err := a.db.RDB.Shutdown(ctx); err != nil {
		log.Println("Failed to shutdown Redis database")
	}

}
