package interfaces

import (
	"net/http"

	"github.com/ozonebg/gofluence/internal/models"
	"github.com/ozonebg/gofluence/pkg/api"
)

// ArticlesController will provide the articles interaction interface for the API.
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

// UsersRepository is a provider for users.
type UsersRepository interface {
	CreateUser(*api.User) error
	All() ([]*api.User, error)
	GetUser(int) (*api.User, error)
	UpdateUser(int, *api.User) error
	DeleteUser(int) error
	GetUserByUsername(string) (*api.User, error)
}
