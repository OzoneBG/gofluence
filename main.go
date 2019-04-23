package main

import (
	"fmt"
	"net/http"

	"github.com/gocraft/dbr"

	"github.com/DavidHuie/gomigrate"

	"github.com/ozonebg/gofluence/config"
	"github.com/ozonebg/gofluence/context"
	"github.com/ozonebg/gofluence/controllers"
	"github.com/ozonebg/gofluence/db"
	"github.com/ozonebg/gofluence/repository"
	"github.com/ozonebg/gofluence/routes"
	log "github.com/sirupsen/logrus"
)

var logger = log.WithField("component", "main")

var (
	adapter gomigrate.Migratable
)

func main() {
	port := config.GetPort()

	context := context.NewContext()

	dsn := fmt.Sprintf("postgres://%v:%v@db/%v?sslmode=disable", config.GetDBUser(), config.GetDBPassword(), config.GetDBName())
	log.WithField("dsn", dsn).Info("connection string")
	session := db.CreateDbConnection(dsn)
	defer session.Close()
	defer session.Connection.Close()

	err := performMigration(session)
	if err != nil {
		log.WithError(err).Fatal("failed to run migrations")
	}

	// instantiate repositories
	context.ArticlesRepository = repository.NewArticlesDao(session)
	context.UsersRepository = repository.NewUsersDao(session)

	// insitantiate deps
	context.ArticlesController = controllers.NewArticlesController(context.ArticlesRepository)
	context.UsersController = controllers.NewUsersController(context.UsersRepository)

	router := routes.NewRouter(context)
	logger.WithField("port", port).Infof("Starting HTTP server listening")
	logger.Fatal(http.ListenAndServe(port, router))
}

func performMigration(sess *dbr.Session) error {
	adapter = gomigrate.Postgres{}
	migrator, err := gomigrate.NewMigratorWithLogger(sess.DB, adapter, "./migrations", log.New())
	if err != nil {
		return err
	}

	err = migrator.Migrate()
	if err != nil {
		return err
	}

	return nil
}
