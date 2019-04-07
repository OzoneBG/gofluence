package db

import (
	"fmt"

	"github.com/gocraft/dbr"
	// for postgres
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/ozonebg/gofluence/config"
)

var dbLogger = log.WithField("component", "db")

// CreateDbConnection returns a new instantiated connection to the database.
func CreateDbConnection() *dbr.Connection {
	dsn := fmt.Sprintf("user=%v dbname=%v sslmode=disable password=%v", config.GetDBUser(), config.GetDBName(), config.GetDBPassword())
	conn, err := dbr.Open("postgres", dsn, nil)
	if err != nil {
		dbLogger.WithError(err).Fatal("failed to connect to database")
	}
	conn.SetMaxOpenConns(30)

	return conn
}
