package api

import (
	"flag"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	port       *string
	debug      *bool
	echoServer *echo.Echo
)

func init() {
	port = flag.String("port", "7000", "port for the service HTTP")
	debug = flag.Bool("debug", false, "mod of the debug")
}

func New() *echo.Echo {
	flag.Parse()

	echoServer = echo.New()

	// Esconde o cabeçalho do Echo
	echoServer.HideBanner = true

	echoServer.Use(middleware.CORS())      // Enable CORS on API
	echoServer.Use(middleware.Recover())   // Recover the API on fatal situations
	echoServer.Use(middleware.RequestID()) // Add a unique ID in each request
	echoServer.Use(middleware.Gzip())      // Compress the request with Gzip

	if *debug {
		echoServer.Debug = true
		echoServer.Use(middleware.Logger())
	}

	return echoServer
}

// Provides the instance of Echo
func ProvideEchoInstance(task func(e *echo.Echo)) {
	task(echoServer)
}

func Run() {
	// For Heroku compatibilite
	porta := os.Getenv("PORT")

	if porta == "" {
		porta = *port
	}

	echoServer.Logger.Fatal(echoServer.Start(":" + porta))
}

func Use(middleware ...echo.MiddlewareFunc) {
	echoServer.Use(middleware...)
}
