package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	FirebaseCredentials string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system env variables.")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	firebaseCredentials := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if firebaseCredentials == "" {
		log.Fatal("FIREBASE_CREDENTIALS_PATH is not set")
	}
	return &Config{
		Port:                port,
		FirebaseCredentials: firebaseCredentials,
	}
}
