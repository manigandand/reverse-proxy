package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

var (
	Env                  string
	Port                 string
	ServerRecipeEndpoint string
)

func init() {
	GetAllEnv()
}

// GetAllEnv should get all the env configs required for the app.
func GetAllEnv() {
	// API Configs
	mustEnv("ENV", &Env, EnvDevelopment)
	mustEnv("PORT", &Port, "8080")
	mustEnv("SEREVR_RECIPE_ENDPOINT", &ServerRecipeEndpoint,
		"https://s3-eu-west-1.amazonaws.com/test-golang-recipes/%d")
}

// mustEnv get the env variable with the name 'key' and store it in 'value'
func mustEnv(key string, value *string, defaultVal string) {
	if *value = os.Getenv(key); *value == "" {
		*value = defaultVal
		log.Infof("%s env variable not set, using default value.\n", key)
	}
}
