package main

import (
	"manigandand-golang-test/pkg/api"
	"manigandand-golang-test/pkg/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	MainSetup()
	defer MainTearDown()
	os.Exit(m.Run())
}

func MainSetup() {
	api.InitAPI()
	testPort := "8001"
	os.Setenv("ENV", config.EnvDevelopment)
	os.Setenv("PORT", testPort)
	os.Setenv("API_HOST", "http://localhost:"+testPort)
	os.Setenv("SEREVR_RECIPE_ENDPOINT", "https://s3-eu-west-1.amazonaws.com/test-golang-recipes/%d")

}

func MainTearDown() {
	// flush all the env
	os.Clearenv()
}
