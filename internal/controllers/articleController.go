package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ozonebg/gofluence/internal/interfaces"
	"github.com/ozonebg/gofluence/internal/models"
	"github.com/ozonebg/gofluence/internal/repository"
	"github.com/ozonebg/gofluence/internal/utils"

	log "github.com/sirupsen/logrus"
)

var articleLogger = log.WithField("component", "article controller")

type articlesController struct {
	articlesRepository interfaces.ArticlesRepository
}

// NewArticlesController returns a new controller for the articles.
func NewArticlesController(articlesDao interfaces.ArticlesRepository) interfaces.ArticlesController {
	return &articlesController{
		articlesRepository: articlesDao,
	}
}

func (ac *articlesController) AllArticles(w http.ResponseWriter, r *http.Request) {
	articleLogger.Info("endpoint hit: all articles")

	articles, err := ac.articlesRepository.All()

	if err != nil {
		if err.Error() == repository.NotFoundArticlesError {
			articleLogger.Info("no articles found")
			json.NewEncoder(w).Encode(articles)
			return
		}
		articleLogger.WithError(err).Info("failed to get articles")
	}

	json.NewEncoder(w).Encode(articles)
}

func (ac *articlesController) GetArticle(w http.ResponseWriter, r *http.Request) {
	articleLogger.Info("endpoint hit: get article")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		articleLogger.WithError(err).Error("failed to get id from url")
		fmt.Fprint(w, "invalid id")
		return
	}

	article, err := ac.articlesRepository.GetArticle(id)
	if err != nil {
		articleLogger.WithError(err).WithField("article_id", id).Error("failed to get article")
		fmt.Fprint(w, "no article found")
		return
	}
	json.NewEncoder(w).Encode(article)
}

func (ac *articlesController) CreateArticle(w http.ResponseWriter, r *http.Request) {
	articleLogger.Info("endpoint hit: create article")

	body, err := utils.ReadRequestBody(r)
	if err != nil {
		articleLogger.WithError(err).Info("failed to read body contents")
	}

	var article models.Article
	json.Unmarshal(body, &article)
	articleLogger.Info(article.AuthorID)

	err = ac.articlesRepository.CreateArticle(&article)
	if err != nil {
		articleLogger.WithError(err).Info("failed to create article")
	}

	articleLogger.WithField("id", article.ID).Info("successfully saved id")
}

func (ac *articlesController) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	articleLogger.Info("endpoint hit: update article")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		articleLogger.WithError(err).Error("failed to get id from url")
		fmt.Fprint(w, "invalid id")
		return
	}

	var updatedArticle models.Article
	body, err := utils.ReadRequestBody(r)
	if err != nil {
		articleLogger.WithError(err).Info("failed to read body contents")
	}

	json.Unmarshal(body, &updatedArticle)

	err = ac.articlesRepository.UpdateArticle(id, &updatedArticle)
	if err != nil {
		articleLogger.WithError(err).Info("failed to update article")
	}
}

func (ac *articlesController) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	articleLogger.Info("endpoint hit: delete article")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		articleLogger.WithError(err).Error("failed to get id from url")
		fmt.Fprint(w, "invalid id")
		return
	}

	err = ac.articlesRepository.DeleteArticle(id)
	if err != nil {
		articleLogger.WithError(err).Info("failed to delete article")
	}
}
