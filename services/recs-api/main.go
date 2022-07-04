package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/pocockn/mono-repo/pkg/logs"
	internalMiddleware "github.com/pocockn/mono-repo/pkg/middleware"
	"github.com/pocockn/mono-repo/services/recs-api/config"
	"github.com/pocockn/mono-repo/services/recs-api/db"
	"github.com/pocockn/mono-repo/services/recs-api/recs/delivery"
	"github.com/pocockn/mono-repo/services/recs-api/recs/store"
)

func main() {
	logs.New(logs.WithDebug(), logs.WithService("recs-api"), logs.WithVersion("0.1.0"))

	cfg := config.NewConfig()

	db, err := db.NewConnection(cfg)
	if err != nil {
		logs.Logger.Fatal().Err(err).Send()
	}

	connection, err := db.Connect()
	if err != nil {
		logs.Logger.Fatal().Err(err).Send()
	}

	e := echo.New()
	e.Use(
		middleware.RequestID(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:8081", "http://localhost:3000"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
		internalMiddleware.LoggerZerolog(logs.Logger),
	)

	recRepo := store.NewRecsStore(connection)
	echoHandler := delivery.NewHandler(cfg, recRepo)
	echoHandler.Register(e)

	e.Logger.Fatal(e.Start(":1323"))
}
