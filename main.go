package main

import (
	"fmt"
	"manigandand-golang-test/pkg/api"
	"manigandand-golang-test/pkg/config"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Init APIs.
	api.InitAPI()

	log.Infoln("Starting server on port :", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port),
		api.LogHandler(api.BaseRoutes.Root))
}
