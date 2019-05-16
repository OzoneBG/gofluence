package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var configLogger = log.WithField("component", "configuration")

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

// GetTokenPassword returns the token password env variable
func GetTokenPassword() string {
	tokenPwd := os.Getenv("TOKEN_PWD")

	if tokenPwd == "" {
		configLogger.Fatal("$TOKEN_PWD must be set")
	}

	return tokenPwd
}

// GetHTTPPort returns the port in use for HTTP.
func GetHTTPPort() string {
	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		configLogger.Fatal("$HTTP_PORT must be set")
	}

	return httpPort
}

// GetGRPCPort returns the port in use for GRPC.
func GetGRPCPort() string {
	grpcPort := os.Getenv("GRPC_PORT")

	if grpcPort == "" {
		configLogger.Fatal("$GRPC_PORT must be set")
	}

	return grpcPort
}
