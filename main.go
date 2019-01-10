package main

import (
	"fmt"
	"manigandand-golang-test/pkg/api"
	"manigandand-golang-test/pkg/config"
	"net/http"

	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Init APIs.
	api.InitAPI()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"*", "Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Header"},
		ExposedHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin", "Origin"},
		AllowCredentials: true,
	})

	log.Infoln("Starting server on port :", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port),
		c.Handler(api.LogHandler(api.BaseRoutes.Root)),
	)
}
