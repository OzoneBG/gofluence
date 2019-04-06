package main

import (
	"net/http"

	"github.com/ozonebg/gofluence/context"
	"github.com/ozonebg/gofluence/controllers"
	"github.com/ozonebg/gofluence/repository"
	"github.com/ozonebg/gofluence/routes"
	log "github.com/sirupsen/logrus"
)

var logger = log.WithField("component", "main")

func main() {
	context := context.NewContext()

	// instantiate repositories
	context.ArticlesRepository = repository.NewArticlesDao()

	// insitantiate deps
	context.ArticlesController = controllers.NewArticlesController(context.ArticlesRepository)

	router := routes.NewRouter(context)
	logger.Info("Starting HTTP server listening on port 80")
	logger.Fatal(http.ListenAndServe(":80", router))
}
