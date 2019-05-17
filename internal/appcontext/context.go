package appcontext

import (
	"github.com/gocraft/dbr"
	"github.com/ozonebg/gofluence/internal/interfaces"
	"github.com/ozonebg/gofluence/internal/repository"
)

// Context keeps track of all services
type Context struct {
	ArticlesRepository interfaces.ArticlesRepository
	UsersRepository    interfaces.UsersRepository
}

// NewContext creates a new context
func NewContext(s *dbr.Session) *Context {
	usersRepository := repository.NewUsersDao(s)
	articlesRepository := repository.NewArticlesDao(s)

	return &Context{
		UsersRepository:    usersRepository,
		ArticlesRepository: articlesRepository,
	}
}
