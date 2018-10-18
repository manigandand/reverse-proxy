package main

import (
	"manigandand-golang-test/pkg/api"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Init APIs.
	api.InitAPI()

	log.Infoln("Starting server on port :8080")
	http.ListenAndServe(":8080", api.LogHandler(api.BaseRoutes.Root))
}
