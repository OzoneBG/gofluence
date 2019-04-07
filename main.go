package main

import (
	"net/http"

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

	logger.WithFields(log.Fields{
		"db_name": config.GetDBName(),
		"db_pass": config.GetDBPassword(),
		"db_user": config.GetDBUser(),
	}).Info("database env variables")

	context := context.NewContext()

	// instantiate repositories
	context.ArticlesRepository = repository.NewArticlesDao()

	// insitantiate deps
	context.ArticlesController = controllers.NewArticlesController(context.ArticlesRepository)

	router := routes.NewRouter(context)
	logger.WithField("port", port).Infof("Starting HTTP server listening")
	logger.Fatal(http.ListenAndServe(port, router))
}
