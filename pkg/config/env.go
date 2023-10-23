package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
)

// LoadEnvFile loads .env file from default path or from custom path if provided
func LoadEnvFile() {
	env := flag.String("e", ".env", "-e /full/path/.env")
	flag.Parse()

	// First attempt to load default .env
	err := godotenv.Load()
	if err != nil && *env == ".env" { // if default file load fails and no alternative specified
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load from custom path if provided
	if *env != ".env" {
		err := godotenv.Load(*env)
		if err != nil {
			log.Fatalf("Error loading .env file from %s: %v", *env, err)
		}
		log.Printf("Loaded .env file from %s", *env)
	} else {
		log.Println("Loaded default .env file")
	}

}
