package web

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vector-ops/go-starter/internal/bootstrap/database"
	"golang.org/x/time/rate"
)

type Middleware struct {
	Server *echo.Echo
	DB     *database.Database
}

func NewMiddleware(server *echo.Echo, DB *database.Database) *Middleware {
	return &Middleware{
		Server: server,
		DB:     DB,
	}
}

func (m *Middleware) Init() {
	m.CORS()
	m.RateLimitter()
}

func (m *Middleware) CORS() {
	m.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.DELETE,
			echo.PUT,
			echo.PATCH,
			echo.HEAD,
		},
		MaxAge: 3600,
	}))
}

func (m *Middleware) RateLimitter() {
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(10), Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, "you have exceeded your limit, please try again after sometime")
		},
	}

	m.Server.Use(middleware.RateLimiterWithConfig(config))
}
