package config

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"

	"looker.com/neutral-farming/internal/types"
	"looker.com/neutral-farming/pkg"
)

var EnvConfig types.Config

var AnalyticsAccessCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "analytics_access_total",
		Help: "Total number of analytics access requests",
	}, []string{"farm_id"},
)

func init() {
	// Load .env by default
	err := godotenv.Load()

	if testing.Testing() {
		log.Println("We are testing - not loading envs")
		return
	}

	if err != nil {
		log.Printf("Error loading envs: %s", err.Error())
	}

	errs := pkg.ParseEnvSchema(&EnvConfig)

	if len(errs) > 0 {
		log.Println("Error loading environment schema:")

		for _, err := range errs {
			log.Println(err)
		}

		log.Fatal("Environment variable validation failed")
	}

	prometheus.MustRegister(AnalyticsAccessCount)

	log.Println("Environment variables loaded successfully")
}
