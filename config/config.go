package config

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var configLogger = log.WithField("component", "configuration")

// GetPort returns the port env variable
func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		configLogger.Fatal("$PORT must be set")
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}

// GetDBUser returns the db username env variable
func GetDBUser() string {
	dbUser := os.Getenv("DB_USER")

	if dbUser == "" {
		configLogger.Fatal("$DB_USER must be set")
	}

	return dbUser
}

// GetDBName returns the db name env variable
func GetDBName() string {
	dbName := os.Getenv("DB_NAME")

	if dbName == "" {
		configLogger.Fatal("$DB_NAME must be set")
	}

	return dbName
}

// GetDBPassword returns the db password env variable
func GetDBPassword() string {
	dbPwd := os.Getenv("DB_PWD")

	if dbPwd == "" {
		configLogger.Fatal("$DB_PWD must be set")
	}

	return dbPwd
}
