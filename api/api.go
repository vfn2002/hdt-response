package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	version string = "/v1"
)

// Init sets up api for service
func Init(c *mongo.Client) {
	e := echo.New()

	e.Use(middleware.Logger())

	g := e.Group(version)

	InitRoutes(g, c)

	e.Logger.Fatal(e.Start(":1323"))
}
