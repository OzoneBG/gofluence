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

	context := context.NewContext()

	// instantiate repositories
	context.ArticlesRepository = repository.NewArticlesDao()
	context.UsersRepository = repository.NewUsersDao()

	// insitantiate deps
	context.ArticlesController = controllers.NewArticlesController(context.ArticlesRepository)
	context.UsersController = controllers.NewUsersController(context.UsersRepository)

	router := routes.NewRouter(context)
	logger.WithField("port", port).Infof("Starting HTTP server listening")
	logger.Fatal(http.ListenAndServe(port, router))
}
