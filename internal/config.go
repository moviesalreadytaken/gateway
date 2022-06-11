package internal

import (
	"log"
	"os"
)

type AppConfig struct {
	MoviesServiceUrl         string
	UsersServiceUrl          string
	RecommendationServiceUrl string
}

func LoadCnfFromEnv() *AppConfig {
	return &AppConfig{
		MoviesServiceUrl:         loadEnvByKey("MOVIES_URL"),
		UsersServiceUrl:          loadEnvByKey("USERS_URL"),
		RecommendationServiceUrl: loadEnvByKey("RECOM_URL"),
	}
}

func loadEnvByKey(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("can't load env var by key = %s", key)
	}
	return val
}
