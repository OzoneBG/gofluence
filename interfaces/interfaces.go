package interfaces

import (
	"net/http"

	"github.com/ozonebg/gofluence/models"
)

// ArticlesController as.
type ArticlesController interface {
	CreateArticle(w http.ResponseWriter, r *http.Request)
	AllArticles(w http.ResponseWriter, r *http.Request)
	GetArticle(w http.ResponseWriter, r *http.Request)
	UpdateArticle(w http.ResponseWriter, r *http.Request)
	DeleteArticle(w http.ResponseWriter, r *http.Request)
}

// ArticlesRepository is a provider for articles
type ArticlesRepository interface {
	CreateArticle(*models.Article) error
	All() (models.Articles, error)
	GetArticle(int) (*models.Article, error)
	UpdateArticle(int, *models.Article) error
	DeleteArticle(int) error
}
