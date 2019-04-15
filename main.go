package main

import (
	"fmt"
	"net/http"

	"github.com/ozonebg/gofluence/db"

	"github.com/ozonebg/gofluence/config"
	"github.com/ozonebg/gofluence/context"
	"github.com/ozonebg/gofluence/controllers"
	"github.com/ozonebg/gofluence/repository"
	"github.com/ozonebg/gofluence/routes"
	log "github.com/sirupsen/logrus"
)

var logger = log.WithField("component", "main")

func main() {
	port := config.GetPort()

	context := context.NewContext()

	dsn := fmt.Sprintf("user=%v dbname=%v sslmode=disable password=%v", config.GetDBUser(), config.GetDBName(), config.GetDBPassword())
	session := db.CreateDbConnection(dsn)
	defer session.Close()
	defer session.Connection.Close()

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
