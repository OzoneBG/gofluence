package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/ozonebg/gofluence/context"
	"github.com/ozonebg/gofluence/controllers"
	"github.com/ozonebg/gofluence/repository"
	"github.com/ozonebg/gofluence/routes"
	log "github.com/sirupsen/logrus"
)

var logger = log.WithField("component", "main")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		logger.Fatal("$PORT must be set")
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	logger.WithField("port", port).Info("current env port value")

	context := context.NewContext()

	// instantiate repositories
	context.ArticlesRepository = repository.NewArticlesDao()

	// insitantiate deps
	context.ArticlesController = controllers.NewArticlesController(context.ArticlesRepository)

	router := routes.NewRouter(context)
	logger.Infof("Starting HTTP server listening on port %v", port)
	logger.Fatal(http.ListenAndServe(port, router))
}
