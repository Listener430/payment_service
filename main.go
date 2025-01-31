package main

import (
	"flag"
	"fmt"
	"net/http"

	_ "github.com/Listener430/payment_service/docs"
	"github.com/Listener430/payment_service/handlers"
	"github.com/Listener430/payment_service/internal/config"
	"github.com/Listener430/payment_service/internal/db"
	"github.com/Listener430/payment_service/internal/rest"
	"github.com/Listener430/payment_service/internal/rest/errors"
	"github.com/Listener430/payment_service/repository"
	"github.com/Listener430/payment_service/verify"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var version string
var buildTime string
var commitHash string

// @title Apple Purchase Verification API
// @version 1.0
// @description Verifies Apple purchases and manages user features
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @type apiKey
func main() {
	cf := flag.String("config", "config.yaml", "Config file")
	flag.Parse()
	if err := config.Load(*cf); err != nil {
		panic(err)
	}
	e := echo.New()
	e.HideBanner = true
	e.Use(rest.RecoveryMiddleware)
	conn, err := db.Init()
	if err != nil {
		e.Logger.Fatal(errors.ErrDBInit.Error() + ": " + err.Error())
	}
	defer db.Close(conn)
	repo := repository.NewRepo(conn)
	v := verify.NewVerifier("https://sandbox.itunes.apple.com/verifyReceipt")
	g := e.Group("/api/v1")
	g.POST("/purchase", handlers.ApplePurchase(repo, v))
	g.GET("/check", handlers.CheckFeature(repo))
	e.GET("/api/v1/swagger/*", echoSwagger.WrapHandler)
	g.GET("/version", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"version":    version,
			"build_time": buildTime,
			"commit":     commitHash,
		})
	})
	p := config.Get().Port
	addr := fmt.Sprintf(":%d", p)
	e.Logger.Infof("Service up on %s", addr)
	e.Logger.Fatal(e.Start(addr))
}
