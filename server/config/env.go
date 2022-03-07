package config

import (
    "os"
    "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

)

func EnvMongoURI() string {
	var l = log.WithFields(log.Fields{
		"method":"ENVMongoURI",
	})
    err := godotenv.Load()
    if err != nil {
        l.Fatal("Error loading .env file")
    }
  
    return os.Getenv("MONGO_URI")
}
