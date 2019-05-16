package interfaces

import (
	"net/http"

	"github.com/ozonebg/gofluence/internal/models"
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

// UsersController will provide the user interaction interface for the API.
type UsersController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	AllUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	Authenticate(w http.ResponseWriter, r *http.Request)
}

// UsersRepository is a provider for users.
type UsersRepository interface {
	CreateUser(*models.User) error
	All() (models.Users, error)
	GetUser(int) (*models.User, error)
	UpdateUser(int, *models.User) error
	DeleteUser(int) error
	GetUserByUsername(string) (*models.User, error)
}
