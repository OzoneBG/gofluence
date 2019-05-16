package context

import (
	"github.com/ozonebg/gofluence/internal/interfaces"
)

// Context keeps track of all services
type Context struct {
	ArticlesController interfaces.ArticlesController
	ArticlesRepository interfaces.ArticlesRepository

	UsersController interfaces.UsersController
	UsersRepository interfaces.UsersRepository
}

// NewContext creates a new context
func NewContext() *Context {
	return &Context{}
}
