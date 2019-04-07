package context

import (
	"github.com/ozonebg/gofluence/interfaces"
)

// Context keeps track of all services
type Context struct {
	ArticlesController interfaces.ArticlesController
	ArticlesRepository interfaces.ArticlesRepository
}

// NewContext creates a new context
func NewContext() *Context {
	return &Context{}
}
