package main

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/gilaz/step-by-step/version"
	info "github.com/k8s-community/handlers/info"
	"github.com/k8s-community/utils/shutdown"
	"github.com/takama/router"
)

var log = logrus.New()

// Run server: go build && step-by-step
// Try requests: curl http://127.0.0.1:8000/test
func main() {
	port := os.Getenv("SERVICE_PORT")

	if len(port) == 0 {
		log.Fatal("SERVICE_PORT is not set.")
	}

	r := router.New()
	r.Logger = logger

	r.GET("/", home)

	r.GET("/info", info.Handler(version.RELEASE, version.REPO, version.COMMIT))

	r.GET("/hearthz", func(c *router.Control) {
		c.Code(http.StatusOK).Body(http.StatusText(http.StatusOK))
	})

	go r.Listen("0.0.0.0:" + port)

	logger := log.WithField("event", "shutdown")
	sdHandler := shutdown.NewHandler(logger)
	sdHandler.RegisterShutdown(sd)
}

// sd does graceful dhutdown of the service
func sd() (string, error) {
	// if service has to finish some tasks before shutting down, these tasks must be finished her
	return "Ok", nil
}
