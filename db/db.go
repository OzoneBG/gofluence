package db

import (
	"github.com/gocraft/dbr"
	// for postgres
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	maxConnections = 30
)

var dbLogger = log.WithField("component", "db")

// CreateDbConnection returns a new instantiated connection to the database.
func CreateDbConnection(dsn string) *dbr.Session {
	conn, err := dbr.Open("postgres", dsn, nil)
	if err != nil {
		dbLogger.WithError(err).Fatal("failed to connect to database")
	}
	conn.SetMaxOpenConns(maxConnections)

	return conn.NewSession(nil)
}
