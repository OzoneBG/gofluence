package context

import (
	"github.com/ozonebg/gofluence/interfaces"
)

type Context struct {
	ArticlesController interfaces.ArticlesController
	ArticlesRepository interfaces.ArticlesRepository
}

func NewContext() *Context {
	return &Context{}
}
