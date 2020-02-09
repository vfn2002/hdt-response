package api

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	version string = "/v1"
)

// Init sets up api for service
func Init(db *sql.DB) {
	e := echo.New()

	e.Use(middleware.Logger())

	g := e.Group(version)

	InitRoutes(g, db)

	e.Logger.Fatal(e.Start(":1323"))
}
