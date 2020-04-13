package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pocockn/mono-repo/services/recs-api/config"
	"github.com/pocockn/mono-repo/services/recs-api/persistance"
	"github.com/pocockn/mono-repo/services/recs-api/recs/delivery"
	"github.com/pocockn/mono-repo/services/recs-api/recs/store"
	"log"
	"net/http"
)

func main() {
	config := config.NewConfig()

	db, err := persistance.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	connection, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:8081", "http://localhost:3000"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
		middleware.Logger(),
	)

	recRepo := store.NewRecsStore(connection)
	echoHandler := delivery.NewHandler(config, recRepo)
	echoHandler.Register(e)

	e.Logger.Fatal(e.Start(":1323"))
}
