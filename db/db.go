package db

import (
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var dbLogger = log.WithField("component", "db")

// CreateDbConnection returns a new instantiated connection to the database.
// Don't forget to defer using defer tx.RollbackUnlessCommited() and then tx.Commit()
func CreateDbConnection() *dbr.Connection {
	conn, err := dbr.Open("postgres", "user=gofluence dbname=gofluence sslmode=disable password=gofluencer", nil)
	if err != nil {
		dbLogger.WithError(err).Fatal("failed to connect to database")
	}
	conn.SetMaxOpenConns(30)

	return conn
}
